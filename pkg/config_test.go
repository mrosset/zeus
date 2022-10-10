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
	"os"
	"testing"
)

var config = NewDefaultConfig("./testdata")

func TestConfig(t *testing.T) {
	checkConfig(config, t)
}

func TestDefault(t *testing.T) {
	var (
		file = "testdata/default.conf"
	)
	if !Exists(file) {
		if err := config.Write(file); err != nil {
			t.Fatal(err)
		}
	}

}

func TestWrite(t *testing.T) {
	var (
		file = "testdata/write.conf"
	)
	defer os.Remove(file)
	if err := config.Write(file); err != nil {
		t.Fatal(err)
	}

}

func checkConfig(config *Config, t *testing.T) {
	var (
		expect = "testdata/data"
	)
	if config.Server != 1 {
		t.Fatal("config.Server should be 1")
	}
	if config.NoListen != 1 {
		t.Fatal("config.NoListen should be 1")
	}
	if config.DBCache != 1000 {
		t.Fatal("config.DBCache should be 1000")
	}
	if config.TXIndex != 1 {
		t.Fatal("config.TXIndex should be ")
	}
	if expect != config.DataDir {
		t.Fatalf("Expected %s got %s", expect, config.DataDir)
	}
	if config.Regtest != 1 {
		t.Fatal("config.Regtest should be 1")
	}
}

func TestRead(t *testing.T) {
	var (
		file = "testdata/default.conf"
	)
	config, err := ReadConfig(file)
	if err != nil {
		t.Fatal(err)
	}
	checkConfig(config, t)
}
