// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcrypto

import (
	"bytes"
	"testing"

	"crypto/sha256"
	"github.com/keyfuse/libracore/xcrypto/edwards"

	"github.com/stretchr/testify/assert"
)

func TestMPCEd25519(t *testing.T) {
	hash := sha256.Sum256([]byte{0x01, 0x02, 0x03, 0x04})

	// party1.
	seed1 := bytes.Repeat([]byte{'x'}, 31)
	prv1, pub1, err := edwards.NewKeyFromSeed(seed1)
	assert.Nil(t, err)
	party1 := NewEd25519Party(prv1)

	// party2.
	seed2 := bytes.Repeat([]byte{'y'}, 32)
	prv2, pub2, err := edwards.NewKeyFromSeed(seed2)
	assert.Nil(t, err)
	party2 := NewEd25519Party(prv2)

	// phase1.
	sharepub1 := party1.Phase1(pub2)
	sharepub2 := party2.Phase1(pub1)
	assert.Equal(t, sharepub1, sharepub2)

	// phase2.
	localnonce1, err := party1.Phase2(hash[:])
	assert.Nil(t, err)
	localnonce2, err := party2.Phase2(hash[:])
	assert.Nil(t, err)

	// phase3.
	sharenonce1 := party1.Phase3(localnonce2)
	sharenonce2 := party2.Phase3(localnonce1)
	assert.Equal(t, sharenonce1, sharenonce2)

	// phase4.
	localsig1, err := party1.Phase4(hash[:], sharepub1, sharenonce1)
	assert.Nil(t, err)
	assert.NotNil(t, localsig1)
	localsig2, err := party2.Phase4(hash[:], sharepub2, sharenonce2)
	assert.Nil(t, err)
	assert.NotNil(t, localsig2)

	// phase5.
	final1, err := party1.Phase5(localsig1, localsig2)
	assert.Nil(t, err)
	final2, err := party2.Phase5(localsig1, localsig2)
	assert.Nil(t, err)
	assert.Equal(t, final1, final2)

	sigv, err := edwards.ParseSignature(final1)
	assert.Nil(t, err)
	ok := edwards.Verify(sharepub1, hash[:], sigv.GetR(), sigv.GetS())
	assert.True(t, ok)
}
