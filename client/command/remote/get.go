//
// client/command/remote/get.go
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

	pb "gopkg.in/cheggaaa/pb.v1"

	"google.golang.org/grpc"

	"golang.org/x/sync/errgroup"

	"github.com/itslab-kyushu/sss/cfg"
	"github.com/itslab-kyushu/sss/kvs"
	"github.com/itslab-kyushu/sss/sss"
	"github.com/urfave/cli"
)

type getOpt struct {
	Config      *cfg.Config
	Name        string
	OutputFile  string
	NConnection int
	Log         io.Writer
}

// CmdGet prepares get command and run cmdGet.
func CmdGet(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}

	conf, err := cfg.ReadConfig(c.String("config"))
	if err != nil {
		return
	}

	output := c.String("output")
	if output == "" {
		output = c.Args().First()
	}

	var log io.Writer
	if c.GlobalBool("quiet") {
		log = ioutil.Discard
	} else {
		log = os.Stderr
	}

	return cmdGet(&getOpt{
		Config:      conf,
		Name:        c.Args().First(),
		OutputFile:  output,
		NConnection: c.Int("max-connection"),
		Log:         log,
	})

}

func cmdGet(opt *getOpt) (err error) {

	if opt.Config.NServers() == 0 {
		return fmt.Errorf("No server information is given.")
	}

	fmt.Fprintln(opt.Log, "Downloading shares")
	bar := pb.New(opt.Config.NServers())
	bar.Output = opt.Log
	bar.Prefix("Server")
	bar.Start()

	shares := make([]sss.Share, opt.Config.NServers())
	wg, ctx := errgroup.WithContext(context.Background())
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
				value, err := client.Get(ctx, &kvs.Key{
					Name: opt.Name,
				})
				if err != nil {
					return
				}

				share, err := FromValue(value)
				if err != nil {
					return
				}
				shares[i] = *share
				return

			})

		}(&server, i)
		i++

	}

	err = wg.Wait()
	bar.Finish()
	if err != nil {
		return
	}

	fmt.Fprintln(opt.Log, "Reconstructing the secret")
	secret, err := sss.Reconstruct(shares)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(opt.OutputFile, secret, 0644)

}
