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
)

func main() {
	seed1 := bytes.Repeat([]byte("x"), 32)
	keypair1 := xcore.GenerateKeyPair(seed1)
	from := xcore.NewAddress([]byte(keypair1.Public.Value.Serialize()))

	seed2 := bytes.Repeat([]byte("k"), 32)
	keypair2 := xcore.GenerateKeyPair(seed2)
	to := xcore.NewAddress([]byte(keypair2.Public.Value.Serialize()))
	toaddr := to.ToString()

	tx, err := xcore.NewTransaction().
		From(from).
		To(toaddr, 123456).
		AddKey(keypair1.Private).
		SetSequence(1).
		SetGasUintPrice(0).
		SetMaxGasAmount(1000000).
		Sign().
		Build()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SigHash: %x", tx.GetSigHash())
	log.Printf("Signature: %x", tx.GetSignature())
}
