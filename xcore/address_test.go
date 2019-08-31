// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"testing"

	"github.com/keyfuse/libracore/xcrypto"
)

func TestAddress(t *testing.T) {
	pub, _, _ := xcrypto.GenerateEd25519KeyPair()

	addr := NewAddress([]byte(pub.Value.Serialize()))
	addrStr := addr.ToString()
	t.Logf("%v", addrStr)
}
