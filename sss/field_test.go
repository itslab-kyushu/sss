//
// sss/field_test.go
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

func TestNewField(t *testing.T) {

	prime, err := rand.Prime(rand.Reader, 258)
	if err != nil {
		t.Fatal(err.Error())
	}

	field := NewField(prime)
	if field.Prime.Cmp(prime) != 0 {
		t.Error("The field has different prime number:", field)
	}
	if field.Max.Cmp(new(big.Int).Sub(prime, big.NewInt(1))) != 0 {
		t.Error("Max of the field is wrong:", field)
	}

}

func TestMarshal(t *testing.T) {

	var err error
	field := NewField(big.NewInt(12345))
	data, err := json.Marshal(field)
	if err != nil {
		t.Fatal(err.Error())
	}

	var res Field
	if err = json.Unmarshal(data, &res); err != nil {
		t.Fatal(err.Error())
	}
	if field.Prime.Cmp(res.Prime) != 0 {
		t.Error("Unmarshaled prime is wrong:", res)
	}
	if field.Max.Cmp(res.Max) != 0 {
		t.Error("Unmarshaled max is wrong:", res)
	}

}
