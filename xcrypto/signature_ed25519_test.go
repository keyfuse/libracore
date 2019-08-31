// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcrypto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEd25519(t *testing.T) {
	{
		prv, pub, _ := GenerateEd25519KeyPair()

		message := []byte("test message")
		sig, err := Ed25519Sign(prv, message)
		assert.Nil(t, err)
		verify := Ed25519Verify(pub, message, sig)
		assert.True(t, verify)
	}

	{
		seed := bytes.Repeat([]byte("x"), 32)
		NewEd25519KeyFromSeed(seed)
	}
}
