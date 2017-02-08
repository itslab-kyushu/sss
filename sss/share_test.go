//
// sss/share_test.go
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

package sss

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"testing"
)

func TestMarshalShare(t *testing.T) {

	var err error
	share := Share{}

	share.Key, err = rand.Int(rand.Reader, big.NewInt(2048))
	if err != nil {
		t.Error(err.Error())
	}
	share.Value = make([]*big.Int, 10)
	for i := 0; i < 10; i++ {
		share.Value[i], err = rand.Int(rand.Reader, big.NewInt(2048))
		if err != nil {
			t.Error(err.Error())
		}
	}
	share.Field = NewField(big.NewInt(2049))

	data, err := json.Marshal(share)
	if err != nil {
		t.Fatal(err.Error())
	}

	var res Share
	if err = json.Unmarshal(data, &res); err != nil {
		t.Error(err.Error())
	}

	if share.Key.Cmp(res.Key) != 0 {
		t.Error("Unmarshalled Key is wrong:", res)
	}
	if len(share.Value) != len(res.Value) {
		t.Error("Unmarshalled Values are wrong:", res)
	}
	for i, v := range share.Value {
		if v.Cmp(res.Value[i]) != 0 {
			t.Error("Unmarshalled Values are wrong:", res)
		}
	}

}
