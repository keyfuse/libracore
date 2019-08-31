// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcrypto

import (
	"math/big"

	"github.com/keyfuse/libracore/xcrypto/edwards"
)

// Ed25519Party -- Ed25519 party struct.
type Ed25519Party struct {
	prv      *edwards.PrivateKey
	curve    *edwards.TwistedEdwardsCurve
	prvnonce *edwards.PrivateKey
	pubnonce *edwards.PublicKey
}

// NewEd25519Party -- create new Ed25519Party.
func NewEd25519Party(prv *edwards.PrivateKey) *Ed25519Party {
	return &Ed25519Party{
		prv:   prv,
		curve: edwards.Edwards(),
	}
}

// Phase1 -- return the aggregate public key.
func (party *Ed25519Party) Phase1(pub2 *edwards.PublicKey) *edwards.PublicKey {
	prv := party.prv
	return edwards.PubkeyAdd([]*edwards.PublicKey{pub2, prv.PubKey()})
}

// Phase2 -- return the local RFC6979 nonce*G.
func (party *Ed25519Party) Phase2(hash []byte) (*edwards.PublicKey, error) {
	var err error

	prv := party.prv
	curve := party.curve

	nonce := edwards.GetRFC6979Nonce(prv, hash)
	nonceBig := new(big.Int).SetBytes(nonce)
	nonceBig.Mod(nonceBig, curve.N)
	nonce = copyBytes(nonceBig.Bytes())[:]
	nonce[31] &= 248

	if party.prvnonce, party.pubnonce, err = edwards.PrivKeyFromScalar(nonce); err != nil {
		return nil, err
	}
	return party.pubnonce, nil
}

// Phase3 -- returns the share nonce.
func (party *Ed25519Party) Phase3(pubnonce2 *edwards.PublicKey) *edwards.PublicKey {
	pubnonce := party.pubnonce
	return edwards.PubkeyAdd([]*edwards.PublicKey{pubnonce, pubnonce2})
}

// Phase4 -- returns the local sign.
func (party *Ed25519Party) Phase4(hash []byte, groupPubkey *edwards.PublicKey, groupPubnonce *edwards.PublicKey) ([]byte, error) {
	prv := party.prv
	prvnonce := party.prvnonce

	r, s, err := edwards.PartialSign(hash, prv, groupPubkey, prvnonce, groupPubnonce)
	if err != nil {
		return nil, err
	}
	localSig := edwards.NewSignature(r, s)
	return localSig.Serialize(), nil
}

// Phase5 -- returns the final sign.
func (party *Ed25519Party) Phase5(sigs ...[]byte) ([]byte, error) {
	var signatures []*edwards.Signature

	for _, sig := range sigs {
		s, err := edwards.ParseSignature(sig)
		if err != nil {
			return nil, err
		}
		signatures = append(signatures, s)
	}

	final, err := edwards.CombinePartialSigs(signatures)
	if err != nil {
		return nil, err
	}
	return final.Serialize(), nil
}

func copyBytes(aB []byte) *[32]byte {
	fieldIntSize := 32
	if aB == nil {
		return nil
	}
	s := new([32]byte)

	// If we have a short byte string, expand
	// it so that it's long enough.
	aBLen := len(aB)
	if aBLen < fieldIntSize {
		diff := fieldIntSize - aBLen
		for i := 0; i < diff; i++ {
			aB = append([]byte{0x00}, aB...)
		}
	}
	for i := 0; i < fieldIntSize; i++ {
		s[i] = aB[i]
	}
	return s
}
