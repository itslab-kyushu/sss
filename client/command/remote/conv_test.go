//
// client/command/remote/conv_test.go
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
	"testing"

	"github.com/itslab-kyushu/sss/sss"
)

func TestConvert(t *testing.T) {

	var err error
	secret := []byte("abcdefg")
	chunksize := 8
	size := 5
	threshold := 2

	shares, err := sss.Distribute(secret, chunksize, size, threshold)
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(shares) != size {
		t.Fatal("Distribute didn't make enough shares.")
	}

	v := ToValue(&shares[0])
	res, err := FromValue(v)
	if err != nil {
		t.Fatal(err.Error())
	}

	if shares[0].Field.Prime.Cmp(res.Field.Prime) != 0 {
		t.Error("Field is not same:", shares[0].Field, res.Field)
	}
	if shares[0].Key.Cmp(res.Key) != 0 {
		t.Error("Key is not same:", shares[0].Key, res.Key)
	}
	for i, v := range shares[0].Value {
		if v.Cmp(res.Value[i]) != 0 {
			t.Error("Shares are not same:", shares[0].Value, res.Value)
		}
	}

}
