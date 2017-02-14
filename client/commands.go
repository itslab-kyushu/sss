//
// commands.go
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

	"github.com/itslab-kyushu/sss/client/command/local"
	"github.com/itslab-kyushu/sss/client/command/remote"
	"github.com/urfave/cli"
)

// DefaultConfFile defines the default configuration file name.
const DefaultConfFile = "sss.yml"

// GlobalFlags defines a set of global flags.
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "quiet",
		Usage: "not output logging infroamtion",
	},
}

// Commands defines a set of commands.
var Commands = []cli.Command{
	{
		Name:  "remote",
		Usage: "Access remote SSS servers",
		Subcommands: cli.Commands{
			{
				Name:        "get",
				Usage:       "Download shares and reconstruct a file",
				Description: "Download shares from a given set of servers and reconstruct original file.",
				ArgsUsage:   "<file name>",
				Action:      remote.CmdGet,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "config",
						Usage: "Server configuration `FILE`.",
						Value: DefaultConfFile,
					},
					cli.StringFlag{
						Name:  "output",
						Usage: "Store the reconstructed secret to the `FILE`.",
					},
					cli.IntFlag{
						Name:  "max-connection",
						Usage: "Maximum connections",
						Value: 10,
					},
				},
			},
			{
				Name:        "put",
				Usage:       "Distribute and store shares",
				Description: "Create shares of the given file and upload them to servers.",
				ArgsUsage:   "<file> <threshold>",
				Action:      remote.CmdPut,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "config",
						Usage: "Server configuration `FILE`.",
						Value: DefaultConfFile,
					},
					cli.IntFlag{
						Name:  "chunk",
						Usage: "Byte `size` of each chunk.",
						Value: 256,
					},
					cli.IntFlag{
						Name:  "max-connection",
						Usage: "Maximum connections",
						Value: 10,
					},
				},
			},
			{
				Name:        "delete",
				Usage:       "Delete a file from all servers",
				Description: "Delete a file from all servers.",
				ArgsUsage:   "<file name>",
				Action:      remote.CmdDelete,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "config",
						Usage: "Server configuration `FILE`.",
						Value: DefaultConfFile,
					},
				},
			},
			{
				Name:        "list",
				Usage:       "Get a list of files stored in servers",
				Description: "Receive a list of files stored in a random server.",
				ArgsUsage:   " ",
				Action:      remote.CmdList,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "config",
						Usage: "Server configuration `FILE`.",
						Value: DefaultConfFile,
					},
				},
			},
		},
	},
	{
		Name:  "local",
		Usage: "Run local file based on a Secret Sharing scheme",
		Subcommands: cli.Commands{
			{
				Name:        "distribute",
				Usage:       "Distribute a file",
				ArgsUsage:   "<file> <share size> <threshold>",
				Description: "distribute command makes a set of shares of a given file.",
				Action:      local.CmdDistribute,
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:  "chunk",
						Usage: "Byte `size` of each chunk.",
						Value: 256,
					},
				},
			},
			{
				Name:        "reconstruct",
				Usage:       "Reconstruct a file from a set of secrets",
				ArgsUsage:   "<file>...",
				Description: "reconstruct command reconstructs a file from a given set of shares.",
				Action:      local.CmdReconstruct,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "output",
						Usage: "Store the reconstructed secret to the `FILE`.",
					},
				},
			},
		},
	},
}

// CommandNotFound handles an error that the given command is not found.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
