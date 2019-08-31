// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// Address --
type Address struct {
	hash []byte
}

// NewAddress -- creates new address.
func NewAddress(pub []byte) *Address {
	digest := sha3.Sum256(pub)
	return &Address{
		hash: digest[:],
	}
}

// Hash -- returns the address hash.
func (a *Address) Hash() []byte {
	return a.hash
}

// ToString -- returns the address hex string.
func (a *Address) ToString() string {
	return hex.EncodeToString(a.hash)
}
