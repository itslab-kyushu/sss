//
// sss/sss_test.go
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

import "testing"

func TestSS1(t *testing.T) {

	var err error

	secret := "abcdefghijklmnopqrstuvwxyz"

	size := 10
	threshold := 3

	shares, err := Distribute([]byte(secret), 8, size, threshold)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(shares)

	res, err := Reconstruct(shares[:threshold])
	if err != nil {
		t.Error(err.Error())
	}
	if string(res) != secret {
		t.Error("SS1 is broken:", res, secret)
	}

}

func TestSS1Word(t *testing.T) {

	var err error

	secret := "a"

	size := 10
	threshold := 3

	shares, err := Distribute([]byte(secret), 256, size, threshold)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(shares)

	res, err := Reconstruct(shares[:threshold])
	if err != nil {
		t.Error(err.Error())
	}
	if string(res) != secret {
		t.Error("SS1 is broken:", res, secret)
	}

}
