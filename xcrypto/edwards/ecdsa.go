// Copyright (c) 2015-2018 The Decred developers
// Copyright 2019 by KeyFuse Labs
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package edwards

import (
	"crypto/sha512"
	"fmt"
	"math/big"

	"github.com/agl/ed25519"
	"github.com/agl/ed25519/edwards25519"
)

// BIG CAVEAT
// Memory management is kind of sloppy and whether or not your keys or
// nonces can be found in memory later is likely a product of when the
// garbage collector runs.
// Signing/EC mult is also not constant side, so don't use this in any
// application where you think you might be vulnerable to side channel
// attacks.

var (
	// oneInitializer is used to fill a byte slice with byte 0x01.  It is provided
	// here to avoid the need to create it multiple times.
	oneInitializer = []byte{0x01}
)

// SignFromSecretNoReader signs a message 'hash' using the given private key
// priv. It doesn't actually user the random reader.
func SignFromSecretNoReader(priv *PrivateKey, hash []byte) (r, s *big.Int, err error) {
	privBytes := priv.SerializeSecret()
	privArray := copyBytes64(privBytes)
	sig := ed25519.Sign(privArray, hash)

	// The signatures are encoded as
	//   sig[0:32]  R, a point encoded as little endian
	//   sig[32:64] S, scalar multiplication/addition results = (ab+c) mod l
	//     encoded also as little endian
	rBytes := copyBytes(sig[0:32])
	r = encodedBytesToBigInt(rBytes)
	sBytes := copyBytes(sig[32:64])
	s = encodedBytesToBigInt(sBytes)

	return
}

// SignFromScalar signs a message 'hash' using the given private scalar priv.
// It uses RFC6979 to generate a deterministic nonce. Considered experimental.
// r = kG, where k is the RFC6979 nonce
// s = r + hash512(k || A || M) * a
func SignFromScalar(priv *PrivateKey, nonce []byte, hash []byte) (r, s *big.Int, err error) {
	publicKey := new([PubKeyBytesLen]byte)
	var A edwards25519.ExtendedGroupElement
	privateScalar := copyBytes(priv.Serialize())
	reverse(privateScalar) // BE --> LE
	edwards25519.GeScalarMultBase(&A, privateScalar)
	A.ToBytes(publicKey)

	// For signing from a scalar, r = nonce.
	nonceLE := copyBytes(nonce)
	reverse(nonceLE)
	var R edwards25519.ExtendedGroupElement
	edwards25519.GeScalarMultBase(&R, nonceLE)

	var encodedR [32]byte
	R.ToBytes(&encodedR)

	// h = hash512(k || A || M)
	h := sha512.New()
	h.Reset()
	if _, err := h.Write(encodedR[:]); err != nil {
		return nil, nil, err
	}
	if _, err := h.Write(publicKey[:]); err != nil {
		return nil, nil, err
	}
	if _, err := h.Write(hash); err != nil {
		return nil, nil, err
	}

	// s = r + h * a
	var hramDigest [64]byte
	h.Sum(hramDigest[:0])
	var hramDigestReduced [32]byte
	edwards25519.ScReduce(&hramDigestReduced, &hramDigest)

	var localS [32]byte
	edwards25519.ScMulAdd(&localS, &hramDigestReduced, privateScalar,
		nonceLE)

	signature := new([64]byte)
	copy(signature[:], encodedR[:])
	copy(signature[32:], localS[:])
	sigEd, err := ParseSignature(signature[:])
	if err != nil {
		return nil, nil, err
	}

	return sigEd.GetR(), sigEd.GetS(), nil
}

// SignThreshold signs a message 'hash' using the given private scalar priv in
// a threshold group signature. It uses RFC6979 to generate a deterministic nonce.
// Considered experimental.
// As opposed to the threshold signing function for secp256k1, this function
// takes the entirety of the public nonce point (all points added) instead of
// the public nonce point with n-1 keys added.
// r = K_Sum
// s = r + hash512(k || A || M) * a
func SignThreshold(priv *PrivateKey, groupPub *PublicKey, hash []byte, privNonce *PrivateKey,
	pubNonceSum *PublicKey) (r, s *big.Int, err error) {

	if priv == nil || hash == nil || privNonce == nil || pubNonceSum == nil {
		return nil, nil, fmt.Errorf("nil input")
	}

	privateScalar := copyBytes(priv.Serialize())
	reverse(privateScalar) // BE --> LE

	// Threshold variant scheme:
	// R = K_Sum
	// Where K_Sum is the sum of the public keys corresponding to
	// the private nonce scalars of each group signature member.
	// That is, R = k1G + ... + knG.
	encodedGroupR := bigIntPointToEncodedBytes(pubNonceSum.GetX(),
		pubNonceSum.GetY())

	// h = hash512(k || A || M)
	var hramDigest [64]byte
	h := sha512.New()
	h.Reset()
	if _, err := h.Write(encodedGroupR[:]); err != nil {
		return nil, nil, err
	}
	if _, err := h.Write(groupPub.Serialize()[:]); err != nil {
		return nil, nil, err
	}
	if _, err := h.Write(hash); err != nil {
		return nil, nil, err
	}
	h.Sum(hramDigest[:0])
	var hramDigestReduced [32]byte
	edwards25519.ScReduce(&hramDigestReduced, &hramDigest)

	// s = r + h * a
	var localS [32]byte
	privNonceLE := copyBytes(privNonce.Serialize())
	reverse(privNonceLE) // BE --> LE
	edwards25519.ScMulAdd(&localS, &hramDigestReduced, privateScalar,
		privNonceLE)

	signature := new([64]byte)
	copy(signature[:], encodedGroupR[:])
	copy(signature[32:], localS[:])
	sigEd, err := ParseSignature(signature[:])
	if err != nil {
		return nil, nil, err
	}

	return sigEd.GetR(), sigEd.GetS(), nil
}

// Sign is the generalized and exported version of Ed25519 signing, that
// handles both standard private secrets and non-standard scalars.
func Sign(priv *PrivateKey, hash []byte) (r, s *big.Int, err error) {
	if priv == nil {
		return nil, nil, fmt.Errorf("private key is nil")
	}
	if hash == nil {
		return nil, nil, fmt.Errorf("message key is nil")
	}

	if priv.secret == nil {
		privLE := copyBytes(priv.Serialize())
		reverse(privLE)
		nonce := nonceRFC6979(privLE[:], hash, nil, nil)
		return SignFromScalar(priv, nonce, hash)
	}

	return SignFromSecretNoReader(priv, hash)
}

// Verify verifies a message 'hash' using the given public keys and signature.
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool {
	if pub == nil || hash == nil || r == nil || s == nil {
		return false
	}

	pubBytes := pub.Serialize()
	sig := &Signature{r, s}
	sigBytes := sig.Serialize()
	pubArray := copyBytes(pubBytes)
	sigArray := copyBytes64(sigBytes)
	return ed25519.Verify(pubArray, hash, sigArray)
}
