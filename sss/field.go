//
// sss/field.go
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
	"encoding/json"
	"fmt"
	"math/big"
)

// Field represents a finite field Z/pZ.
type Field struct {
	// Prime number.
	Prime *big.Int
	// Max is Prime - 1.
	Max *big.Int
}

// compactField defines a field to marshal/unmarshal.
type compactField struct {
	// Prime number.
	Prime string
}

// NewField creates a new finite field.
func NewField(prime *big.Int) *Field {

	return &Field{
		Prime: prime,
		Max:   new(big.Int).Sub(prime, big.NewInt(1)),
	}

}

// MarshalJSON implements Marshaler interface.
func (f *Field) MarshalJSON() ([]byte, error) {

	aux := compactField{
		Prime: f.Prime.Text(16),
	}
	return json.Marshal(aux)

}

// UnmarshalJSON implements Unmarshaler interface.
func (f *Field) UnmarshalJSON(data []byte) (err error) {

	var aux compactField
	if err = json.Unmarshal(data, &aux); err != nil {
		return
	}
	var ok bool
	if f.Prime, ok = new(big.Int).SetString(aux.Prime, 16); !ok {
		return fmt.Errorf("Given Field is broken: %v", aux)
	}
	f.Max = new(big.Int).Sub(f.Prime, big.NewInt(1))
	return

}
