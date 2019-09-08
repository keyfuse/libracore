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

// Event --
type Event struct {
	Count uint64
	Key   []byte
}

// AccountState --
type AccountState struct {
	AuthenticationKey             string
	Balance                       uint64
	SequenceNumber                uint64
	SentEvents                    Event
	RecievedEvents                Event
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

	if a.RecievedEvents.Count, err = buf.ReadU64(); err != nil {
		return err
	}
	var rel uint32
	if rel, err = buf.ReadU32(); err != nil {
		return err
	}
	if a.RecievedEvents.Key, err = buf.ReadBytes(int(rel)); err != nil {
		return err
	}
	if a.SentEvents.Count, err = buf.ReadU64(); err != nil {
		return err
	}
	var sel uint32
	if sel, err = buf.ReadU32(); err != nil {
		return err
	}
	if a.SentEvents.Key, err = buf.ReadBytes(int(sel)); err != nil {
		return err
	}
	if a.SequenceNumber, err = buf.ReadU64(); err != nil {
		return err
	}
	return nil
}
