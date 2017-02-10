//
// command/reconstruct.go
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
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/itslab-kyushu/sss/sss"
	"github.com/ulikunitz/xz"
	"github.com/urfave/cli"
)

type reconstructOpt struct {
	ShareFiles []string
	OutputFile string
}

// CmdReconstruct executes reconstruct command.
func CmdReconstruct(c *cli.Context) error {

	if c.NArg() == 0 {
		return cli.ShowSubcommandHelp(c)
	}

	opt := &reconstructOpt{
		ShareFiles: append([]string{c.Args().First()}, c.Args().Tail()...),
		OutputFile: c.String("output"),
	}
	if opt.OutputFile == "" {
		opt.OutputFile = outputFile(opt.ShareFiles[0])
	}

	return cmdReconstruct(opt)

}

func cmdReconstruct(opt *reconstructOpt) (err error) {

	wg, ctx := errgroup.WithContext(context.Background())
	semaphore := make(chan struct{}, runtime.NumCPU())

	shares := make([]sss.Share, len(opt.ShareFiles))
	for i, f := range opt.ShareFiles {

		func(i int, f string) {

			select {
			case <-ctx.Done():
				return
			case semaphore <- struct{}{}:
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()

					fp, err := os.Open(f)
					if err != nil {
						return
					}
					defer fp.Close()

					r, err := xz.NewReader(fp)
					if err != nil {
						return
					}
					data, err := ioutil.ReadAll(r)
					if err != nil {
						return
					}
					return json.Unmarshal(data, &shares[i])

				})

			}

		}(i, f)

	}

	if err = wg.Wait(); err != nil {
		return err
	}

	secret, err := sss.Reconstruct(shares)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(opt.OutputFile, secret, 0644)

}

// outputFile returns a filename associated with the given share file name.
func outputFile(sharename string) string {

	components := strings.Split(sharename, ".")
	if len(components) < 2 {
		return ""
	}
	return strings.Join(components[:len(components)-2], ".")

}
