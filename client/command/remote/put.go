//
// client/command/remote/put.go
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

package remote

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	pb "gopkg.in/cheggaaa/pb.v1"

	"google.golang.org/grpc"

	"golang.org/x/sync/errgroup"

	"github.com/itslab-kyushu/sss/cfg"
	"github.com/itslab-kyushu/sss/kvs"
	"github.com/itslab-kyushu/sss/sss"
	"github.com/urfave/cli"
)

// putOpt defines option values for cmdPut.
type putOpt struct {
	Filename    string
	Size        int
	Threshold   int
	Config      *cfg.Config
	NConnection int
	ChunkSize   int
	Log         io.Writer
}

// CmdPut prepares put command and run cmdPut.
func CmdPut(c *cli.Context) (err error) {

	if c.NArg() != 3 {
		return cli.ShowSubcommandHelp(c)
	}
	conf, err := cfg.ReadConfig(c.String("config"))
	if err != nil {
		return
	}

	size, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return
	}

	threshold, err := strconv.Atoi(c.Args().Get(2))
	if err != nil {
		return
	}

	var output io.Writer
	if c.GlobalBool("quiet") {
		output = ioutil.Discard
	} else {
		output = os.Stderr
	}

	return cmdPut(&putOpt{
		Filename:    c.Args().First(),
		Size:        size,
		Threshold:   threshold,
		Config:      conf,
		NConnection: c.Int("max-connection"),
		ChunkSize:   c.Int("chunk"),
		Log:         output,
	})

}

func cmdPut(opt *putOpt) (err error) {

	secret, err := ioutil.ReadFile(opt.Filename)
	if err != nil {
		return
	}

	fmt.Fprintln(opt.Log, "Creating shares")
	ctx := context.Background()
	shares, err := sss.Distribute(secret, opt.ChunkSize, opt.Size, opt.Threshold)
	if err != nil {
		return
	}

	fmt.Fprintln(opt.Log, "Uploading shares")
	bar := pb.New(opt.Config.NServers())
	bar.Output = opt.Log
	bar.Prefix("Server")
	bar.Start()
	defer bar.Finish()

	wg, ctx := errgroup.WithContext(ctx)
	semaphore := make(chan struct{}, opt.NConnection)
	var i int
	for _, server := range opt.Config.Servers {

		// Check the current context.
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		func(server *cfg.Server, i int) {

			semaphore <- struct{}{}
			wg.Go(func() (err error) {
				defer func() { <-semaphore }()
				defer bar.Increment()

				conn, err := grpc.Dial(
					fmt.Sprintf("%s:%d", server.Address, server.Port),
					grpc.WithInsecure(),
					grpc.WithCompressor(grpc.NewGZIPCompressor()),
					grpc.WithDecompressor(grpc.NewGZIPDecompressor()),
				)
				if err != nil {
					return
				}
				defer conn.Close()

				client := kvs.NewKvsClient(conn)
				_, err = client.Put(ctx, &kvs.Entry{
					Key: &kvs.Key{
						Name: opt.Filename,
					},
					Value: ToValue(&shares[i]),
				})
				return

			})

		}(&server, i)
		i++

	}
	return wg.Wait()

}
