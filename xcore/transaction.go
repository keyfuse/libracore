// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"encoding/binary"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/golang/protobuf/proto"
	"github.com/keyfuse/libracore/xcrypto"
	"github.com/keyfuse/libracore/xmove"
	"github.com/keyfuse/libracore/xproto"
)

const (
	// LIBRA_HASH_PREFIX_HEX -- sha3.256("@@$$LIBRA$$@@")
	LIBRA_HASH_PREFIX_HEX = "46f174df6ca8de5ad29745f91584bb913e7df8dd162e3e921a5c1d8637c88d16"
)

// Transaction -- txn struct.
type Transaction struct {
	from           *Address
	to             string
	key            *xcrypto.PrivateKey
	needSign       bool
	sendValue      uint64
	gasUnitPrice   uint64
	maxGasAmount   uint64
	expirationTime uint64
	sequence       uint64
	rawtx          *xproto.RawTransaction
	txproto        []byte
	sighash        []byte
	signature      []byte
}

// NewTransaction -- creates new txn.
func NewTransaction() *Transaction {
	return &Transaction{
		maxGasAmount:   10000,
		expirationTime: 100,
	}
}

// From -- coin which from
func (tx *Transaction) From(addr *Address) *Transaction {
	tx.from = addr
	return tx
}

// AddKey -- the from coin private key.
func (tx *Transaction) AddKey(key *xcrypto.PrivateKey) *Transaction {
	tx.key = key
	return tx
}

// To -- transfer the coins to.
func (tx *Transaction) To(addr string, value uint64) *Transaction {
	tx.to = addr
	tx.sendValue = value
	return tx
}

// Sign -- sign the txn.
func (tx *Transaction) Sign() *Transaction {
	tx.needSign = true
	return tx
}

// SetSequence -- set the sequence of this txn.
func (tx *Transaction) SetSequence(seq uint64) *Transaction {
	tx.sequence = seq
	return tx
}

// SetGasUintPrice -- set the gas price.
func (tx *Transaction) SetGasUintPrice(price uint64) *Transaction {
	tx.gasUnitPrice = price
	return tx
}

// SetMaxGasAmount -- set the max gas amount.
func (tx *Transaction) SetMaxGasAmount(amount uint64) *Transaction {
	tx.maxGasAmount = amount
	return tx
}

// GetBytes -- get the txproto bytes.
func (tx *Transaction) GetBytes() []byte {
	return tx.txproto
}

// GetSigHash -- get the unsigned sighash.
func (tx *Transaction) GetSigHash() []byte {
	return tx.sighash
}

// GetSignature -- get the signed signature.
func (tx *Transaction) GetSignature() []byte {
	return tx.signature
}

// Build -- build the txn.
func (tx *Transaction) Build() (*Transaction, error) {
	var err error
	var program *xproto.Program

	// Program.
	{
		to, err := hex.DecodeString(tx.to)
		if err != nil {
			return nil, err
		}
		code, _ := hex.DecodeString(xmove.PEER_TO_PEER_TRANSFER_OPCODE)
		amount := make([]byte, 8)
		binary.LittleEndian.PutUint64(amount, tx.sendValue)

		arg1 := &xproto.TransactionArgument{Type: xproto.TransactionArgument_ADDRESS, Data: to}
		arg2 := &xproto.TransactionArgument{Type: xproto.TransactionArgument_U64, Data: amount}
		arg := []*xproto.TransactionArgument{arg1, arg2}

		modules := [][]byte{}
		program = &xproto.Program{Code: code, Arguments: arg, Modules: modules}
	}

	// Raw transaction hash.
	{
		tx.rawtx = &xproto.RawTransaction{
			SequenceNumber: tx.sequence,
			SenderAccount:  tx.from.Hash(),
			MaxGasAmount:   tx.maxGasAmount,
			GasUnitPrice:   tx.gasUnitPrice,
			Payload:        &xproto.RawTransaction_Program{program},
			ExpirationTime: uint64(time.Now().Unix()) + tx.expirationTime,
		}

		if tx.txproto, err = proto.Marshal(tx.rawtx); err != nil {
			return nil, err
		}
		hashPrefix, _ := hex.DecodeString(LIBRA_HASH_PREFIX_HEX)
		hasher := sha3.New256()
		hasher.Write(hashPrefix)
		hasher.Write(tx.txproto[:])
		tx.sighash = hasher.Sum(tx.sighash[:0])
	}

	// Sign.
	if tx.needSign {
		if tx.signature, err = xcrypto.Ed25519Sign(tx.key, tx.sighash); err != nil {
			return nil, err
		}
	}
	return tx, nil
}
