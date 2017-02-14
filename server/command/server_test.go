//
// server/command/server_test.go
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

package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/grpc/metadata"

	"github.com/itslab-kyushu/sss/kvs"
	context "golang.org/x/net/context"
)

func TestServer(t *testing.T) {

	var err error
	root, err := ioutil.TempDir("", "test_server")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.RemoveAll(root)

	ctx := context.Background()
	server := &Server{
		Root: root,
		Log:  os.Stdout,
	}

	entry := &kvs.Entry{
		Key: &kvs.Key{
			Name: "test",
		},
		Value: &kvs.Value{
			Key: "data-key",
		},
	}

	_, err = server.Put(ctx, entry)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = os.Stat(filepath.Join(root, entry.Key.Name))
	if err != nil {
		t.Error(err.Error())
	}

	res, err := server.Get(ctx, entry.Key)
	if err != nil {
		t.Error(err.Error())
	} else if res.Key != "data-key" {
		t.Error("Get returns a wrong share:", res)
	}

	_, err = server.Delete(ctx, entry.Key)
	if err != nil {
		t.Error(err.Error())
	}
	_, err = os.Stat(filepath.Join(root, entry.Key.Name))
	if err == nil {
		t.Error("Delete doesn't delete the given share")
	}

}

type mockListServer struct {
	Messages []*kvs.Key
}

func (s *mockListServer) Send(key *kvs.Key) error {
	s.Messages = append(s.Messages, key)
	return nil
}

func (s *mockListServer) Context() context.Context {
	return context.Background()
}

func (s *mockListServer) SendMsg(m interface{}) error {
	return fmt.Errorf("Not implemented")
}

func (s *mockListServer) RecvMsg(m interface{}) error {
	return fmt.Errorf("Not implemented")
}

func (s *mockListServer) SetHeader(metadata.MD) error {
	return fmt.Errorf("Not implemented")
}

func (s *mockListServer) SendHeader(metadata.MD) error {
	return fmt.Errorf("Not implemented")
}

func (s *mockListServer) SetTrailer(metadata.MD) {

}

func TestList(t *testing.T) {

	var err error
	name := "test_file"

	root, err := ioutil.TempDir("", "test_list")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.RemoveAll(root)

	err = ioutil.WriteFile(filepath.Join(root, name), []byte(name), 0644)
	if err != nil {
		t.Fatal(err.Error())
	}

	bucket, err := ioutil.TempDir(root, "test_bucket")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = ioutil.WriteFile(filepath.Join(bucket, name), []byte(name), 0644)
	if err != nil {
		t.Fatal(err.Error())
	}

	server := Server{
		Root: root,
		Log:  os.Stdout,
	}
	mock := mockListServer{}
	err = server.List(&kvs.ListRequest{}, &mock)
	if err != nil {
		t.Error("List returns an error:", err.Error())
	}

	if len(mock.Messages) != 2 {
		t.Error("List sent wrong messages:", mock.Messages)
	}

	for _, v := range mock.Messages {
		if v.Name != name && v.Name != fmt.Sprintf("%s/%s", filepath.Base(bucket), name) {
			t.Error("List sent wrong messages:", v.Name)
		}
	}

}
