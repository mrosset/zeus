/*
Copyright © 2022 Michael Rosset <mike.rosset@gmail.com>

# This file is part of Zeus

Zeus is free software: you can redistribute it and/or modify it
under the terms of the GNU General Public License as published by the
Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Zeus is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package zeus

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/subpop/go-ini"
)

type Config struct {
	Server   int    `ini:"server"`
	NoListen int    `ini:"nolisten"`
	DBCache  int    `ini:"dbcache"`
	TXIndex  int    `ini:"txindex"`
	DataDir  string `ini:"datadir"`
	Regtest  int    `ini:"regtest"`
}

// Returns a new default *Config with PREFIX substitution
func NewDefaultConfig(prefix string) *Config {
	return &Config{
		Server:   1,
		NoListen: 1,
		DBCache:  1000,
		TXIndex:  1,
		DataDir:  filepath.Join(prefix, "data"),
		Regtest:  1,
	}
}

// Writes *Config to INI PATH
func (c *Config) Write(path string) error {
	fi, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fi.Close()
	b, err := ini.Marshal(c)
	if err != nil {
		return err
	}
	_, err = fi.Write(b)
	return err
}

// Reads *Config from INI PATH
func ReadConfig(path string) (*Config, error) {
	var (
		config = new(Config)
	)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = ini.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}
	return config, err
}
