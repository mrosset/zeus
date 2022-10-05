/*
Copyright © 2022 Michael Rosset <mike.rosset@gmail.com>

# This file is part of Raijin

Raijin is free software: you can redistribute it and/or modify it
under the terms of the GNU General Public License as published by the
Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Raijin is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package raijin

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/subpop/go-ini"
)

type Config struct {
	DataDir string `ini:"datadir"`
}

// Writes *Config to INI PATH
func (config *Config) Write(path string) error {
	fi, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fi.Close()
	b, err := ini.Marshal(config)
	if err != nil {
		return err
	}
	_, err = fi.Write(b)
	return err
}

// Returns a new *Config with PREFIX substitution
func NewBitcoinConfig(prefix string) *Config {
	return &Config{
		DataDir: filepath.Join(prefix, "data"),
	}
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
