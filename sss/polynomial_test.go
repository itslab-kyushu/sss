//
// sss/polynomial_test.go
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
	"math/big"
	"testing"
)

func TestPolynomial(t *testing.T) {

	field := NewField(big.NewInt(37))
	threshold := 2
	polynomial, err := NewPolynomial(field, big.NewInt(0), threshold-1)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(polynomial.Coefficients) != 1 {
		t.Fatal("Dimension of the polynomial is wrong:", polynomial)
	}

	var res *big.Int
	if res = polynomial.Call(big.NewInt(0)); res.Int64() != 0 {
		t.Error("F(0) returns a wrong value:")
	}

	if res = polynomial.Call(big.NewInt(1)); res.Cmp(polynomial.Coefficients[0]) != 0 {
		t.Error("F(1) returns a wrong value:", res)
	}

}
