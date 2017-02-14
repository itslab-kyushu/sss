//
// server/command/server.go
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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	proto "github.com/golang/protobuf/proto"
	"github.com/ulikunitz/xz"
	context "golang.org/x/net/context"

	"github.com/itslab-kyushu/sss/kvs"
)

// DateFormat defines a format to output logging information.
const DateFormat = "2006-01-02 15:04:05"

// Server defines a KVS server.
type Server struct {
	// Root is a path to the document root.
	Root string
	// Compress stored data.
	Compress bool
	// Log is a writer to output logging information.
	Log io.Writer
}

// Get returns a value associated with the given key.
func (s *Server) Get(ctx context.Context, key *kvs.Key) (res *kvs.Value, err error) {

	target := filepath.Join(s.Root, filepath.ToSlash(key.Name))
	info, err := os.Stat(target)
	if err != nil {
		return
	} else if info.IsDir() {
		return nil, fmt.Errorf("The given key is a bucket name")
	}

	var data []byte
	if s.Compress {
		fp, err := os.Open(target)
		if err != nil {
			return nil, err
		}
		defer fp.Close()

		r, err := xz.NewReader(fp)
		if err != nil {
			return nil, err
		}
		data, err = ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}

	} else {
		data, err = ioutil.ReadFile(target)
		if err != nil {
			return nil, err
		}
	}

	res = &kvs.Value{}
	if err = proto.Unmarshal(data, res); err != nil {
		return
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		fmt.Fprintln(s.Log, time.Now().Local().Format(DateFormat), "GET", key.Name)
		return res, nil
	}

}

// Put stores a given entry as a file.
func (s *Server) Put(ctx context.Context, entry *kvs.Entry) (*kvs.PutResponse, error) {

	var err error

	target := filepath.Join(s.Root, filepath.ToSlash(entry.Key.Name))
	if err = os.MkdirAll(filepath.Dir(target), 0755); err != nil {
		return nil, err
	}
	info, err := os.Stat(target)
	if err == nil && info.IsDir() {
		return nil, fmt.Errorf("The given key is used as a bucket name")
	}

	data, err := proto.Marshal(entry.Value)
	if err != nil {
		return nil, err
	}

	if s.Compress {
		fp, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		defer fp.Close()

		w, err := xz.NewWriter(fp)
		if err != nil {
			return nil, err
		}
		defer w.Close()

		for {
			n, err := w.Write(data)
			if err != nil {
				return nil, err
			}
			if n == len(data) {
				break
			}
			data = data[n:]
		}
		return &kvs.PutResponse{}, nil
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		fmt.Fprintln(s.Log, time.Now().Local().Format(DateFormat), "PUT", entry.Key.Name)
		return &kvs.PutResponse{}, ioutil.WriteFile(target, data, 0644)
	}

}

// Delete deletes a given file.
func (s *Server) Delete(ctx context.Context, key *kvs.Key) (res *kvs.DeleteResponse, err error) {

	var info os.FileInfo
	target := filepath.Join(s.Root, filepath.ToSlash(key.Name))
	info, err = os.Stat(target)
	if err != nil {
		return
	} else if info.IsDir() {
		return nil, fmt.Errorf("Given key is not associated with any items")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		fmt.Fprintln(s.Log, time.Now().Local().Format(DateFormat), "DELETE", target)
		if err = os.Remove(target); err != nil {
			return
		}

		dir := filepath.Dir(target)
		for {
			if dir == s.Root {
				break
			}

			info, err = os.Stat(dir)
			if err != nil {
				return
			}
			if !info.IsDir() && info.Size() != 0 {
				break
			}
			fmt.Fprintln(s.Log, time.Now().Local().Format(DateFormat), "DELETE", dir)
			if err = os.Remove(dir); err != nil {
				return
			}
			dir = filepath.Dir(dir)

		}

	}
	return &kvs.DeleteResponse{}, err

}

// List lists up items stored in this KVS.
func (s *Server) List(_ *kvs.ListRequest, server kvs.Kvs_ListServer) error {

	fmt.Fprintln(s.Log, time.Now().Local().Format(DateFormat), "LIST")
	ctx := server.Context()
	return filepath.Walk(s.Root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			if info.IsDir() {
				return nil
			}
			item, err := filepath.Rel(s.Root, path)
			if err != nil {
				return err
			}
			return server.Send(&kvs.Key{
				Name: item,
			})
		}

	})

}
