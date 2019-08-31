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
	addrs := []string{
		"596c41d9af1bedbeddac1c22dc00bf0809b4c809ccbd7f28c415a1f78fb2c5cd",
		"a516f5626e160e4408c82844bf30df0087e3d31850c8c7df49d69480f13ad0ee",
		"1ad1a1d56e160e4408c82844bf30df0087e3d31850c8c7df49d69480f13ad0ee",
	}

	client, err := xcore.NewClient(xcore.TestNet)
	if err != nil {
		log.Fatal(err)
	}

	states, err := client.QueryAccountStates(addrs)
	if err != nil {
		log.Fatal(err)
	}

	for _, state := range states {
		log.Printf("%+v", state)
	}
}
