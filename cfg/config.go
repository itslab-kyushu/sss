//
// cfg/config.go
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

package cfg

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config defines a set of group servers to distribute shares.
type Config struct {
	Servers []Server
}

// Server defines server information.
type Server struct {
	Address string
	Port    int
}

// ReadConfig reads a YAML formatted config file.
func ReadConfig(filename string) (conf *Config, err error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	conf = new(Config)
	err = yaml.Unmarshal(data, conf)
	return

}

// NServers returns the number of all servers.
func (c *Config) NServers() int {
	return len(c.Servers)
}
