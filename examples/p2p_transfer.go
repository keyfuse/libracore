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
	fromaddr := from.ToString()

	seed2 := bytes.Repeat([]byte("k"), 32)
	keypair2 := xcore.GenerateKeyPair(seed2)
	to := xcore.NewAddress([]byte(keypair2.Public.Value.Serialize()))
	toaddr := to.ToString()

	client, err := xcore.NewClient(xcore.TestNet)
	if err != nil {
		log.Fatal(err)
	}

	var amount uint64 = 123456
	// Sync.
	{
		log.Printf("transfer.from[%v].to[%v].sync...", fromaddr, toaddr)
		sender := &xcore.Account{
			Address: from,
			KeyPair: keypair1,
		}
		if err := client.TransferCoins(sender, toaddr, amount, 0, 100000); err != nil {
			log.Fatal(err)
		}
	}

	// Async.
	{
		log.Printf("transfer.from[%v].to[%v].async...", fromaddr, toaddr)
		sender := &xcore.Account{
			Address: from,
			KeyPair: keypair1,
		}
		if err := client.TransferCoinsAsync(sender, toaddr, amount, 0, 100000); err != nil {
			log.Fatal(err)
		}
	}
}
