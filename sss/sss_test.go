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

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestSSS(t *testing.T) {

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

func TestSSSWord(t *testing.T) {

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

// Example code for Distribute function.
func ExampleDistribute() {
	chunksize := 256
	totalShares := 10
	threshold := 5

	secret, err := ioutil.ReadFile("secret-file")
	if err != nil {
		log.Fatal(err)
	}

	shares, err := Distribute(secret, chunksize, totalShares, threshold)
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range shares {
		data, err := json.Marshal(s)
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("%s.%d.json", "share-", i)
		if err = ioutil.WriteFile(filename, data, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

// Example code for Reconstruct function.
func ExampleReconstruct() {
	// filenames is a slice of file names of shares.
	filenames := []string{"share1.dat", "share2.dat", "share3.dat"}

	shares := make([]Share, len(filenames))
	for i, f := range filenames {

		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		if err = json.Unmarshal(data, &shares[i]); err != nil {
			log.Fatal(err)
		}

	}

	secret, err := Reconstruct(shares)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("secret-file", secret, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
