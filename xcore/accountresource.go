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

// AccountResource --
type AccountResource struct {
	State map[string][]byte
}

// Deserialize -- deserialize the payload to AccountResource.
func (a *AccountResource) Deserialize(payload []byte) error {
	var err error
	buf := xbase.NewBufferReader(payload)
	a.State = make(map[string][]byte)

	var cnts uint32
	if cnts, err = buf.ReadU32(); err != nil {
		return err
	}

	for i := 0; i < int(cnts); i++ {
		var keyLen uint32
		if keyLen, err = buf.ReadU32(); err != nil {
			return err
		}

		var key []byte
		if key, err = buf.ReadBytes(int(keyLen)); err != nil {
			return err
		}

		var valLen uint32
		if valLen, err = buf.ReadU32(); err != nil {
			return err
		}

		var val []byte
		if val, err = buf.ReadBytes(int(valLen)); err != nil {
			return err
		}
		a.State[hex.EncodeToString(key)] = val[:]
	}
	return nil
}
