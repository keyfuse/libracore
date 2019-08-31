// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/keyfuse/libracore/xcrypto"
	"github.com/keyfuse/libracore/xproto"
	"google.golang.org/grpc"
)

func TestClientMintCoins(t *testing.T) {
	seed1 := bytes.Repeat([]byte("x"), 32)
	keypair1 := GenerateKeyPair(seed1)
	from := NewAddress([]byte(keypair1.Public.Value.Serialize()))
	fromaddr := from.ToString()

	client, err := NewClient(TestNet)
	assert.Nil(t, err)

	// Mint.
	{
		err := client.MintWithFaucetService(fromaddr, 100000000)
		assert.Nil(t, err)
	}
}

func TestClientQueryAccountStates(t *testing.T) {
	seed1 := bytes.Repeat([]byte("x"), 32)
	keypair1 := GenerateKeyPair(seed1)
	from := NewAddress([]byte(keypair1.Public.Value.Serialize()))
	fromaddr := from.ToString()

	seed2 := bytes.Repeat([]byte("d"), 32)
	keypair2 := GenerateKeyPair(seed2)
	from2 := NewAddress([]byte(keypair2.Public.Value.Serialize()))
	fromaddr2 := from2.ToString()

	client, err := NewClient(TestNet)
	assert.Nil(t, err)

	{
		states, err := client.QueryAccountStates([]string{fromaddr, fromaddr2})
		assert.Nil(t, err)
		t.Logf("states:%+v", states)
	}
}

func TestClientTransferCoins(t *testing.T) {
	seed1 := bytes.Repeat([]byte("x"), 32)
	keypair1 := GenerateKeyPair(seed1)
	from := NewAddress([]byte(keypair1.Public.Value.Serialize()))
	fromaddr := from.ToString()

	seed2 := bytes.Repeat([]byte("k"), 32)
	keypair2 := GenerateKeyPair(seed2)
	to := NewAddress([]byte(keypair2.Public.Value.Serialize()))
	toaddr := to.ToString()

	client, err := NewClient(TestNet)
	assert.Nil(t, err)

	// Sync.
	{
		t.Logf("sender[%v].to[%v]", fromaddr, toaddr)
		sender := &Account{
			Address: from,
			KeyPair: keypair1,
		}
		err := client.TransferCoins(sender, toaddr, 124356, 0, 100000)
		assert.Nil(t, err)
	}

	// Async.
	{
		t.Logf("sender[%v].to[%v]", fromaddr, toaddr)
		sender := &Account{
			Address: from,
			KeyPair: keypair1,
		}
		err := client.TransferCoinsAsync(sender, toaddr, 124356, 0, 100000)
		assert.Nil(t, err)
	}
}

func TestClientTwoPartyTransaction(t *testing.T) {
	// party1.
	seed1 := bytes.Repeat([]byte{'f'}, 31)
	keypair1 := GenerateKeyPair(seed1)
	prv1 := keypair1.Private.Value
	pub1 := keypair1.Public.Value
	party1 := xcrypto.NewEd25519Party(prv1)

	// party2.
	seed2 := bytes.Repeat([]byte{'u'}, 32)
	keypair2 := GenerateKeyPair(seed2)
	prv2 := keypair2.Private.Value
	pub2 := keypair2.Public.Value
	party2 := xcrypto.NewEd25519Party(prv2)

	// to.
	seed3 := bytes.Repeat([]byte{'c'}, 32)
	keypair3 := GenerateKeyPair(seed3)
	pub3 := keypair3.Public.Value
	to := NewAddress([]byte(pub3.Serialize())).ToString()

	// phase1.
	sharepub1 := party1.Phase1(pub2)
	sharepub2 := party2.Phase1(pub1)
	assert.Equal(t, sharepub1, sharepub2)
	from := NewAddress([]byte(sharepub1.Serialize()))
	fromaddr := from.ToString()
	t.Logf("from:%v", fromaddr)
	t.Logf("to:%v", to)

	client, err := NewClient(TestNet)
	assert.Nil(t, err)

	// Mint.
	{
		err := client.MintWithFaucetService(fromaddr, 100000000)
		assert.Nil(t, err)
	}

	// Transfer.
	{
		states, err := client.QueryAccountStates([]string{fromaddr})
		assert.Nil(t, err)
		accountState := states[0]
		t.Logf("from.state:%+v", accountState)

		tx, err := NewTransaction().
			From(from).
			To(to, 123456).
			SetSequence(accountState.SequenceNumber).
			SetGasUintPrice(0).
			SetMaxGasAmount(1000000).
			Build()
		assert.Nil(t, err)

		sighash := tx.GetSigHash()

		// phase2.
		localnonce1, err := party1.Phase2(sighash[:])
		assert.Nil(t, err)
		localnonce2, err := party2.Phase2(sighash[:])
		assert.Nil(t, err)

		// phase3.
		sharenonce1 := party1.Phase3(localnonce2)
		sharenonce2 := party2.Phase3(localnonce1)
		assert.Equal(t, sharenonce1, sharenonce2)

		// phase4.
		localsig1, err := party1.Phase4(sighash[:], sharepub1, sharenonce1)
		assert.Nil(t, err)
		localsig2, err := party2.Phase4(sighash[:], sharepub2, sharenonce2)
		assert.Nil(t, err)

		// phase5.
		final1, err := party1.Phase5(localsig1, localsig2)
		assert.Nil(t, err)

		ok := xcrypto.Ed25519Verify(&xcrypto.PublicKey{Value: sharepub1}, sighash, final1)
		assert.True(t, ok)

		signedTxn := &xproto.SignedTransaction{
			SenderPublicKey: sharepub1.Serialize(),
			RawTxnBytes:     tx.GetBytes(),
			SenderSignature: final1,
		}
		req := &xproto.SubmitTransactionRequest{
			SignedTxn: signedTxn,
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		res, err := client.client.SubmitTransaction(ctx, req, grpc.WaitForReady(true))
		assert.Nil(t, err)
		acStatus := res.GetAcStatus()
		assert.NotNil(t, acStatus)
		t.Logf("status:%v", acStatus)
	}
}
