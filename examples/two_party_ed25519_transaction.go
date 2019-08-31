// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package main

import (
	"bytes"
	"log"

	"github.com/keyfuse/libracore/xcore"
	"github.com/keyfuse/libracore/xcrypto"
	"github.com/keyfuse/libracore/xproto"
)

func main() {
	var sequence uint64

	// party1.
	seed1 := bytes.Repeat([]byte{'l'}, 31)
	keypair1 := xcore.GenerateKeyPair(seed1)
	prv1 := keypair1.Private.Value
	pub1 := keypair1.Public.Value
	party1 := xcrypto.NewEd25519Party(prv1)

	// party2.
	seed2 := bytes.Repeat([]byte{'i'}, 32)
	keypair2 := xcore.GenerateKeyPair(seed2)
	prv2 := keypair2.Private.Value
	pub2 := keypair2.Public.Value
	party2 := xcrypto.NewEd25519Party(prv2)

	// to.
	seed3 := bytes.Repeat([]byte{'b'}, 32)
	keypair3 := xcore.GenerateKeyPair(seed3)
	pub3 := keypair3.Public.Value
	toaddr := xcore.NewAddress([]byte(pub3.Serialize())).ToString()

	// phase1.
	sharepub1 := party1.Phase1(pub2)
	sharepub2 := party2.Phase1(pub1)
	from := xcore.NewAddress([]byte(sharepub1.Serialize()))
	fromaddr := from.ToString()
	log.Printf("twoparty.address:%v, to.address:%v", fromaddr, toaddr)

	client, err := xcore.NewClient(xcore.TestNet)
	if err != nil {
		log.Fatal(err)
	}

	// Mint.
	{
		amount := 1000000000
		if err := client.MintWithFaucetService(fromaddr, uint64(amount)); err != nil {
			log.Fatal(err)
		}
		log.Printf("twoparty.address.facuet[%v].coins", amount)
	}

	// Transfer.
	{
		states, err := client.QueryAccountStates([]string{fromaddr})
		if err != nil {
			log.Fatal(err)
		}
		accountState := states[0]
		sequence = accountState.SequenceNumber
		log.Printf("twoparty.accountstate:%+v", accountState)

		tx, err := xcore.NewTransaction().
			From(from).
			To(toaddr, 123456).
			SetSequence(sequence).
			SetGasUintPrice(0).
			SetMaxGasAmount(1000000).
			Build()

		if err != nil {
			log.Fatal(err)
		}
		sighash := tx.GetSigHash()
		log.Printf("twoparty.unsigned.hash:%x", sighash)

		// phase2.
		localnonce1, err := party1.Phase2(sighash[:])
		if err != nil {
			log.Fatal(err)
		}
		localnonce2, err := party2.Phase2(sighash[:])
		if err != nil {
			log.Fatal(err)
		}

		// phase3.
		sharenonce1 := party1.Phase3(localnonce2)
		sharenonce2 := party2.Phase3(localnonce1)

		// phase4.
		localsig1, err := party1.Phase4(sighash[:], sharepub1, sharenonce1)
		if err != nil {
			log.Fatal(err)
		}
		localsig2, err := party2.Phase4(sighash[:], sharepub2, sharenonce2)
		if err != nil {
			log.Fatal(err)
		}

		// phase5.
		final1, err := party1.Phase5(localsig1, localsig2)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("twoparty.signed.hash:%x", final1)

		ok := xcrypto.Ed25519Verify(&xcrypto.PublicKey{Value: sharepub1}, sighash, final1)
		if !ok {
			log.Fatal("twoparty.signature.failed")
		}
		log.Printf("twoparty.signed.hash.verify:%v", ok)

		signedTxn := &xproto.SignedTransaction{
			SenderPublicKey: sharepub1.Serialize(),
			RawTxnBytes:     tx.GetBytes(),
			SenderSignature: final1,
		}
		req := &xproto.SubmitTransactionRequest{
			SignedTxn: signedTxn,
		}

		log.Printf("transaction.submit....")
		if err := client.SubmitTransaction(req); err != nil {
			log.Fatal(err)
		}

		log.Printf("transaction.wait.to.valid....")
		if err := client.WaitForTransaction(fromaddr, sequence+1); err != nil {
			log.Fatal(err)
		}
		log.Printf("transaction.done https://libexplorer.com/address/%v", toaddr)
	}
}
