// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package edwards

import (
	"math/big"
)

// PubkeyAdd -- returns the aggregate of the pubkeys.
func PubkeyAdd(pks []*PublicKey) *PublicKey {
	return combinePubkeys(pks)
}

// GetRFC6979Nonce -- returns the deterministic nonce.
func GetRFC6979Nonce(prv *PrivateKey, hash []byte) []byte {
	return nonceRFC6979(prv.Serialize(), hash, nil, Sha512VersionStringRFC6979)
}

// CombinePartialSigs -- returns the signature of aggragation.
func CombinePartialSigs(sigs []*Signature) (*Signature, error) {
	return schnorrCombinePartialSigs(sigs)
}

// PartialSign -- returns the partial signature of this private key.
func PartialSign(hash []byte, prv *PrivateKey, groupPubkey *PublicKey, prvNonce *PrivateKey, pubNonce *PublicKey) (*big.Int, *big.Int, error) {
	return schnorrPartialSign(hash, prv.Serialize(), groupPubkey.Serialize(), prvNonce.Serialize(), pubNonce.Serialize())
}
