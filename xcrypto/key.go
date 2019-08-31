// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcrypto

import (
	"github.com/keyfuse/libracore/xcrypto/edwards"
)

// PrivateKey --
type PrivateKey struct {
	Value *edwards.PrivateKey
}

// PublicKey --
type PublicKey struct {
	Value *edwards.PublicKey
}

// KeyPair --
type KeyPair struct {
	Private *PrivateKey
	Public  *PublicKey
}
