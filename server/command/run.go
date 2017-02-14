//
// server/command/run.go
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
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"

	"github.com/itslab-kyushu/sss/kvs"
	"github.com/urfave/cli"
)

// CmdRun runs a simple KVS server.
func CmdRun(c *cli.Context) (err error) {

	var log io.Writer
	if c.Bool("quiet") {
		log = ioutil.Discard
	} else {
		log = os.Stderr
	}

	port := c.GlobalInt("port")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}

	root, err := filepath.Abs(filepath.ToSlash(c.GlobalString("root")))
	if err != nil {
		return
	}
	info, err := os.Stat(root)
	if err != nil {
		if err = os.MkdirAll(root, 0755); err != nil {
			return err
		}
	} else if !info.IsDir() {
		return fmt.Errorf("Given document root isn't a directory:", root)
	}
	fmt.Fprintln(log, "Document root is set to", root)

	s := grpc.NewServer(
		grpc.RPCCompressor(grpc.NewGZIPCompressor()),
		grpc.RPCDecompressor(grpc.NewGZIPDecompressor()),
		grpc.MaxMsgSize(c.Int("max-message-size")),
	)
	kvs.RegisterKvsServer(s, &Server{
		Root:     root,
		Compress: !c.Bool("no-compress"),
		Log:      log,
	})

	fmt.Fprintln(log, "Start listening:", port)
	return s.Serve(listen)

}
