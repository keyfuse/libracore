// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcrypto

import (
	"github.com/keyfuse/libracore/xcrypto/edwards"
)

// GenerateEd25519KeyPair -- generate prv/pub pair by random.
func GenerateEd25519KeyPair() (*PrivateKey, *PublicKey, error) {
	prv, pub, err := edwards.GenerateKeyPair()
	if err != nil {
		return nil, nil, err
	}
	return &PrivateKey{Value: prv}, &PublicKey{Value: pub}, nil
}

// NewEd25519KeyFromSeed -- generate prv/pub pair by the seed hash.
func NewEd25519KeyFromSeed(seed []byte) (*PrivateKey, *PublicKey) {
	prv, pub, err := edwards.NewKeyFromSeed(seed)
	if err != nil {
		return nil, nil
	}
	return &PrivateKey{Value: prv}, &PublicKey{Value: pub}
}

// Ed25519Sign -- sign.
func Ed25519Sign(prv *PrivateKey, message []byte) ([]byte, error) {
	sig, err := prv.Value.Sign(message)
	if err != nil {
		return nil, err
	}
	return sig.Serialize(), nil
}

// Ed25519Verify -- verfiy.
func Ed25519Verify(pub *PublicKey, message []byte, sig []byte) bool {
	sigv, err := edwards.ParseSignature(sig)
	if err != nil {
		return false
	}
	return edwards.Verify(pub.Value, message, sigv.GetR(), sigv.GetS())
}
