// Copyright (c) 2015-2018 The Decred developers
// Copyright 2019 by KeyFuse Labs
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package edwards

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"math/big"
)

// nonceRFC6979 is a local instantiation of deterministic nonce generation
// by the standards of RFC6979.
func nonceRFC6979(privkey []byte, hash []byte, extra []byte, version []byte) []byte {
	pkD := new(big.Int).SetBytes(privkey)
	defer pkD.SetInt64(0)
	bigK := NonceRFC6979(pkD, hash, extra, version)
	defer bigK.SetInt64(0)
	k := bigIntToEncodedBytesNoReverse(bigK)
	return k[:]
}

// NonceRFC6979 generates an ECDSA nonce (`k`) deterministically according to
// RFC 6979. It takes a 32-byte hash as an input and returns 32-byte nonce to
// be used in ECDSA algorithm.
func NonceRFC6979(privkey *big.Int, hash []byte, extra []byte, version []byte) *big.Int {
	curve := Edwards()
	q := curve.Params().N
	x := privkey
	alg := sha256.New

	qlen := q.BitLen()
	holen := alg().Size()
	rolen := (qlen + 7) >> 3
	bx := append(int2octets(x, rolen), bits2octets(hash, rolen)...)
	if len(extra) == 32 {
		bx = append(bx, extra...)
	}
	if len(version) == 16 && len(extra) == 32 {
		bx = append(bx, extra...)
	}
	if len(version) == 16 && len(extra) != 32 {
		bx = append(bx, bytes.Repeat([]byte{0x00}, 32)...)
		bx = append(bx, version...)
	}

	// Step B
	v := bytes.Repeat(oneInitializer, holen)

	// Step C (Go zeroes the all allocated memory)
	k := make([]byte, holen)

	// Step D
	k = mac(alg, k, append(append(v, 0x00), bx...))

	// Step E
	v = mac(alg, k, v)

	// Step F
	k = mac(alg, k, append(append(v, 0x01), bx...))

	// Step G
	v = mac(alg, k, v)

	// Step H
	for {
		// Step H1
		var t []byte

		// Step H2
		for len(t)*8 < qlen {
			v = mac(alg, k, v)
			t = append(t, v...)
		}

		// Step H3
		secret := hashToInt(t, curve)
		if secret.Cmp(one) >= 0 && secret.Cmp(q) < 0 {
			return secret
		}
		k = mac(alg, k, append(v, 0x00))
		v = mac(alg, k, v)
	}
}

// hashToInt converts a hash value to an integer. There is some disagreement
// about how this is done. [NSA] suggests that this is done in the obvious
// manner, but [SECG] truncates the hash to the bit-length of the curve order
// first. We follow [SECG] because that's what OpenSSL does. Additionally,
// OpenSSL right shifts excess bits from the number if the hash is too large
// and we mirror that too.
// This is borrowed from crypto/ecdsa.
func hashToInt(hash []byte, c *TwistedEdwardsCurve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

// mac returns an HMAC of the given key and message.
func mac(alg func() hash.Hash, k, m []byte) []byte {
	h := hmac.New(alg, k)
	if _, err := h.Write(m); err != nil {
		panic(err)
	}
	return h.Sum(nil)
}

// https://tools.ietf.org/html/rfc6979#section-2.3.3
func int2octets(v *big.Int, rolen int) []byte {
	out := v.Bytes()

	// left pad with zeros if it's too short
	if len(out) < rolen {
		out2 := make([]byte, rolen)
		copy(out2[rolen-len(out):], out)
		return out2
	}

	// drop most significant bytes if it's too long
	if len(out) > rolen {
		out2 := make([]byte, rolen)
		copy(out2, out[len(out)-rolen:])
		return out2
	}

	return out
}

// https://tools.ietf.org/html/rfc6979#section-2.3.4
func bits2octets(in []byte, rolen int) []byte {
	curve := Edwards()
	z1 := hashToInt(in, curve)
	z2 := new(big.Int).Sub(z1, curve.Params().N)
	if z2.Sign() < 0 {
		return int2octets(z1, rolen)
	}
	return int2octets(z2, rolen)
}
