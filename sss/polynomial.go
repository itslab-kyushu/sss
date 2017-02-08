//
// sss/polynomial.go
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
	"math/big"
)

// Polynomial represents a polynomial defined on a finite field.
type Polynomial struct {
	Field        *Field
	Dimension    int
	Coefficients []*big.Int
	Const        *big.Int
}

// NewPolynomial creates a new random polynomial on the given finite field.
// The dimension of the polynomial is the given dim, and it has a given
// constant.
func NewPolynomial(field *Field, c *big.Int, dim int) (*Polynomial, error) {

	coeffs := make([]*big.Int, dim)
	for i := 0; i < dim; i++ {
		v, err := rand.Int(rand.Reader, field.Max)
		if err != nil {
			return nil, err
		}
		coeffs[i] = v
	}

	return &Polynomial{
		Field:        field,
		Dimension:    dim,
		Coefficients: coeffs,
		Const:        c,
	}, nil

}

// Call computes a value F(x) where x is the given number.
func (p *Polynomial) Call(x *big.Int) *big.Int {

	res := big.NewInt(0)
	for i, a := range p.Coefficients {
		y := new(big.Int).Exp(x, big.NewInt(int64(i+1)), p.Field.Prime)
		res.Add(res, y.Mul(y, a))
	}
	res.Add(res, p.Const)
	return res.Mod(res, p.Field.Prime)

}
