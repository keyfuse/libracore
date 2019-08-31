// Copyright (c) 2019 The Decred developers
// Copyright 2019 by KeyFuse Labs
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package edwards

import (
	"testing"
)

func TestPrivKeySign(t *testing.T) {
	message := []byte("the quick brown fox jumps over the lazy dog")

	{
		privkey, _, err := NewKeyFromSeed([]byte{0x01, 0x02})
		if err != nil {
			t.Fatalf("GeneratePrivateKey: %v", err)
		}
		r, s, err := Sign(privkey, message)
		if err != nil {
			t.Fatalf("Sign: %v", err)
		}
		sig1 := NewSignature(r, s)

		sig2, err := privkey.Sign(message)
		if err != nil {
			t.Fatalf("privkey.Sign: %v", err)
		}
		if !sig1.IsEqual(sig2) {
			t.Fatalf("Sign and (PrivateKey).Sign differ")
		}
	}
}

func TestPrivKeyGenerate(t *testing.T) {
	_, _, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("GenerateKeyPair: %v", err)
	}
}
