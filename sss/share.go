//
// sss/share.go
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

// Share defines a share of Shamir's Secret Sharing scheme.
type Share struct {
	Key   *big.Int
	Value []*big.Int
	Field *Field
}

type compactShare struct {
	Key   string
	Value []string
	Field *Field
}

// MarshalJSON implements Marshaler interface.
func (s Share) MarshalJSON() ([]byte, error) {

	aux := compactShare{
		Key:   s.Key.Text(16),
		Value: make([]string, len(s.Value)),
		Field: s.Field,
	}
	for i, v := range s.Value {
		aux.Value[i] = v.Text(16)
	}
	return json.Marshal(aux)

}

// UnmarshalJSON implements Unmarshaler interface.
func (s *Share) UnmarshalJSON(data []byte) (err error) {

	var aux compactShare
	if err = json.Unmarshal(data, &aux); err != nil {
		return
	}
	var ok bool
	if s.Key, ok = new(big.Int).SetString(aux.Key, 16); !ok {
		return fmt.Errorf("Given share is broken: %v", aux)
	}

	s.Value = make([]*big.Int, len(aux.Value))
	for i, v := range aux.Value {
		if s.Value[i], ok = new(big.Int).SetString(v, 16); !ok {
			return fmt.Errorf("Given share is broken: %v", aux)
		}
	}
	s.Field = aux.Field
	return

}
