//
// client/command/local/distribute.go
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

package local

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"

	pb "gopkg.in/cheggaaa/pb.v1"

	"golang.org/x/sync/errgroup"

	"github.com/itslab-kyushu/sss/sss"
	"github.com/ulikunitz/xz"
	"github.com/urfave/cli"
)

type distributeOpt struct {
	Filename  string
	ChunkSize int
	Size      int
	Threshold int
	Log       io.Writer
}

// CmdDistribute executes distribute command.
func CmdDistribute(c *cli.Context) (err error) {

	if c.NArg() != 3 {
		return cli.ShowSubcommandHelp(c)
	}

	threshold, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return
	}
	size, err := strconv.Atoi(c.Args().Get(2))
	if err != nil {
		return
	}
	var log io.Writer
	if c.Bool("quiet") {
		log = ioutil.Discard
	} else {
		log = os.Stderr
	}

	return cmdDistribute(&distributeOpt{
		Filename:  c.Args().Get(0),
		ChunkSize: c.Int("chunk"),
		Size:      size,
		Threshold: threshold,
		Log:       log,
	})
}

func cmdDistribute(opt *distributeOpt) (err error) {

	fmt.Fprintln(opt.Log, "Reading the secret file.")
	secret, err := ioutil.ReadFile(opt.Filename)
	if err != nil {
		return
	}

	fmt.Fprintln(opt.Log, "Computing shares.")
	shares, err := sss.Distribute(secret, opt.ChunkSize, opt.Size, opt.Threshold)
	if err != nil {
		return
	}

	fmt.Fprint(opt.Log, "Writing share files.")
	bar := pb.New(len(shares))
	bar.Output = opt.Log
	bar.Prefix("Files")
	bar.Start()
	defer bar.Finish()

	wg, ctx := errgroup.WithContext(context.Background())
	semaphore := make(chan struct{}, runtime.NumCPU())
	for i, s := range shares {

		func(i int, s sss.Share) {

			select {
			case <-ctx.Done():
				return
			case semaphore <- struct{}{}:
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()
					defer bar.Increment()

					data, err := json.Marshal(s)
					if err != nil {
						return
					}

					fp, err := os.OpenFile(fmt.Sprintf("%s.%d.xz", opt.Filename, i), os.O_WRONLY|os.O_CREATE, 0644)
					if err != nil {
						return
					}
					defer fp.Close()

					w, err := xz.NewWriter(fp)
					if err != nil {
						return
					}
					defer w.Close()

					for {
						select {
						case <-ctx.Done():
							return ctx.Err()
						default:
							n, err := w.Write(data)
							if err != nil {
								return err
							}
							if n == len(data) {
								return nil
							}
							data = data[n:]
						}
					}

				})
			}

		}(i, s)

	}

	return wg.Wait()
}
