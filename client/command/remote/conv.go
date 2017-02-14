//
// client/command/remote/conv.go
//
// Copyright (c) 2017 Junpei Kawamoto
//
// This file is part of sss.
//
// sss is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// sss is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with sss.  If not, see <http://www.gnu.org/licenses/>.
//

package remote

import (
	"fmt"
	"math/big"

	"github.com/itslab-kyushu/sss/kvs"
	"github.com/itslab-kyushu/sss/sss"
)

// ToValue converts a share to a value.
func ToValue(share *sss.Share) (value *kvs.Value) {

	value = &kvs.Value{
		Field:  share.Field.Prime.Text(16),
		Key:    share.Key.Text(16),
		Shares: make([]string, len(share.Value)),
	}
	for i, v := range share.Value {
		value.Shares[i] = v.Text(16)
	}
	return

}

// FromValue converts a value to a share.
func FromValue(value *kvs.Value) (*sss.Share, error) {

	var ok bool
	field, ok := new(big.Int).SetString(value.Field, 16)
	if !ok {
		return nil, fmt.Errorf("Cannot convert the field: %v", value.Field)
	}
	key, ok := new(big.Int).SetString(value.Key, 16)
	if !ok {
		return nil, fmt.Errorf("Cannot convert the key: %v", value.Key)
	}

	res := &sss.Share{
		Field: sss.NewField(field),
		Key:   key,
		Value: make([]*big.Int, len(value.Shares)),
	}
	for i, v := range value.Shares {
		if res.Value[i], ok = new(big.Int).SetString(v, 16); !ok {
			return nil, fmt.Errorf("Cannot convert a group share: %v", v)
		}
	}

	return res, nil

}
