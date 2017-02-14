//
// server/commands.go
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

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// GlobalFlags defines a set of global flags.
var GlobalFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "port",
		Usage: "Listening `port` number",
		Value: 13009,
	},
	cli.StringFlag{
		Name:  "root",
		Usage: "Document root `path`",
		Value: "data",
	},
	cli.BoolFlag{
		Name:  "no-compress",
		Usage: "Store data files without compression",
	},
	cli.IntFlag{
		Name:  "max-message-size",
		Usage: "Maximum acceptable message `byte` size",
		Value: 1024 * 1024 * 256,
	},
	cli.BoolFlag{
		Name:  "quiet",
		Usage: "Omit printing logging information.",
	},
}

// Commands defines a set of commands.
var Commands = cli.Commands{}

// CommandNotFound handles an error that the given command is not found.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
