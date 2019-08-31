// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"testing"

	"github.com/keyfuse/libracore/xcrypto"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	prvf, pubf, _ := xcrypto.GenerateEd25519KeyPair()
	pubt, _, _ := xcrypto.GenerateEd25519KeyPair()

	tx, err := NewTransaction().
		From(NewAddress([]byte(pubf.Value.Serialize()))).
		AddKey(prvf).
		To(NewAddress([]byte(pubt.Value.Serialize())).ToString(), 1000).
		Sign().
		Build()

	assert.Nil(t, err)
	t.Logf("tx:%v", tx)
}
