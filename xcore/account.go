// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"github.com/keyfuse/libracore/xcrypto"
)

// Account --
type Account struct {
	Address *Address
	KeyPair *xcrypto.KeyPair
}

// GenerateKeyPair -- generate the prv/pub key pair.
func GenerateKeyPair(seed []byte) *xcrypto.KeyPair {
	prv, pub := xcrypto.NewEd25519KeyFromSeed(seed)
	return &xcrypto.KeyPair{
		Private: prv,
		Public:  pub,
	}
}
