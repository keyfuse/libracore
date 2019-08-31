// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"encoding/hex"

	"github.com/keyfuse/libracore/xbase"
)

const (
	// AccountStatePath --
	AccountStatePath = "01217da6c6b3e19f1825cfb2676daecce3bf3de03cf26647c78df00b371b25cc97"
)

// AccountState --
type AccountState struct {
	AuthenticationKey             string
	Balance                       uint64
	SequenceNumber                uint64
	SentEventsCount               uint64
	RecievedEventsCount           uint64
	DelegatedWithdrawalCapability bool
}

// Deserialize -- deserialize the payload to AccountState.
func (a *AccountState) Deserialize(payload []byte) error {
	var err error
	var authenticationKeyLen uint32
	buf := xbase.NewBufferReader(payload)

	var data []byte
	if authenticationKeyLen, err = buf.ReadU32(); err != nil {
		return err
	}
	if data, err = buf.ReadBytes(int(authenticationKeyLen)); err != nil {
		return err
	}
	a.AuthenticationKey = hex.EncodeToString(data)

	if a.Balance, err = buf.ReadU64(); err != nil {
		return err
	}
	if a.DelegatedWithdrawalCapability, err = buf.ReadBool(); err != nil {
		return err
	}
	if a.RecievedEventsCount, err = buf.ReadU64(); err != nil {
		return err
	}
	if a.SentEventsCount, err = buf.ReadU64(); err != nil {
		return err
	}
	if a.SequenceNumber, err = buf.ReadU64(); err != nil {
		return err
	}
	return nil
}
