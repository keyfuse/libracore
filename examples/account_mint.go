// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package main

import (
	"log"

	"github.com/keyfuse/libracore/xcore"
)

func main() {
	addr := "596c41d9af1bedbeddac1c22dc00bf0809b4c809ccbd7f28c415a1f78fb2c5cd"

	client, err := xcore.NewClient(xcore.TestNet)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.MintWithFaucetService(addr, 100000000); err != nil {
		log.Fatal(err)
	}
}
