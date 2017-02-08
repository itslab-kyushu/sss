//
// sss/sss.go
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
	"fmt"
	"math"
	"math/big"
)

// Distribute computes shares having a given secret.
func Distribute(secret []byte, chunkByte, size, threshold int) (shares []Share, err error) {

	prime, err := rand.Prime(rand.Reader, chunkByte*8+1)
	if err != nil {
		return
	}
	field := NewField(prime)

	nvalue := int(math.Ceil(float64(len(secret)) / float64(chunkByte)))
	shares = make([]Share, size)
	for i := range shares {
		key := big.NewInt(int64(i + 1))
		shares[i] = Share{
			Key:   key,
			Value: make([]*big.Int, nvalue),
			Field: field,
		}
	}

	var value *big.Int
	for chunk := 0; chunk < nvalue; chunk++ {
		if len(secret) > chunkByte {
			value = new(big.Int).SetBytes(secret[:chunkByte])
			secret = secret[chunkByte:]
		} else {
			value = new(big.Int).SetBytes(secret)
			secret = nil
		}

		polynomial, err := NewPolynomial(field, value, threshold-1)
		if err != nil {
			return nil, err
		}

		for i := range shares {
			key := big.NewInt(int64(i + 1))
			shares[i].Value[chunk] = polynomial.Call(key)
		}

	}

	return
}

// Reconstruct computes the secret value from a set of shares.
func Reconstruct(shares []Share) (bytes []byte, err error) {

	if len(shares) == 0 {
		err = fmt.Errorf("No shares are given")
		return
	}

	bytes = []byte{}
	for chunk := 0; chunk < len(shares[0].Value); chunk++ {

		value := big.NewInt(0)
		field := shares[0].Field
		for i, s := range shares {
			value.Add(value, new(big.Int).Mul(s.Value[chunk], beta(field, shares, i)))
		}
		value.Mod(value, field.Prime)

		bytes = append(bytes, value.Bytes()...)

	}
	return

}

// beta computes the following value:
//   \mul_{i<=u<=k, u!=t} \frac{u-th key}{(u-th key) - (t-th key)}
func beta(field *Field, shares []Share, t int) *big.Int {

	res := big.NewInt(1)
	for i, s := range shares {
		if i == t {
			continue
		}
		sub := new(big.Int).Mod(new(big.Int).Sub(s.Key, shares[t].Key), field.Prime)
		v := new(big.Int).Mul(s.Key, new(big.Int).ModInverse(sub, field.Prime))
		res.Mul(res, v)
		res.Mod(res, field.Prime)
	}

	return res.Mod(res, field.Prime)

}
