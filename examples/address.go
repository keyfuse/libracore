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
)

func main() {
	{
		pub, _, _ := xcrypto.GenerateEd25519KeyPair()

		addr := xcore.NewAddress([]byte(pub.Value.Serialize()))
		addrStr := addr.ToString()
		log.Printf("random.address:%v", addrStr)
	}

	{
		seed1 := bytes.Repeat([]byte{'l'}, 31)
		keypair1 := xcore.GenerateKeyPair(seed1)
		addr := xcore.NewAddress([]byte(keypair1.Public.Value.Serialize()))
		addrStr := addr.ToString()
		log.Printf("seed.address:%v", addrStr)
	}
}
