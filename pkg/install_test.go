/*
Copyright © 2022 Michael Rosset <mike.rosset@gmail.com>

# This file is part of Raijin

PROGRAM is free software: you can redistribute it and/or modify it
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
	"reflect"
	"testing"
)

func TestInstallerType(t *testing.T) {
	var (
		gzExpect = "temp/gz"
		gzPath   = "temp/gz/bitcoin-23.0-x86_64-linux-gnu.tar.gz"
		got      = NewBitcoinInstaller("amd64", "linux", "temp", LAN)
		expect   = &Installer{
			Description: "Bitcoin Core",
			hash:        "2CCA490C1F2842884A3C5B0606F179F9F937177DA4EADD628E3F7FD7E25D26D0",
			arch:        "amd64",
			os:          "linux",
			prefix:      "temp",
			commands: []string{
				"bin/test_bitcoin",
				"bin/bitcoind",
				"bin/bitcoin-wallet",
				"bin/bitcoin-qt",
				"bin/bitcoin-tx",
				"bin/bitcoin-util",
				"bin/bitcoin-cli"},
			tarDir: "bitcoin-23.0",
			uri:    "http://10.119.176.16/bitcoin-23.0-x86_64-linux-gnu.tar.gz"}
	)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v expect %v", got, expect)
	}
	if got.GzDir() != gzExpect {
		t.Errorf("expect %s got %s", gzExpect, got.GzDir())
	}
	if got.GzPath() != gzPath {
		t.Errorf("expect %s got %s", gzPath, got.GzPath())
	}
}

func TestInstall(t *testing.T) {
	prefix, err := ioutil.TempDir(os.TempDir(), "raijin")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(prefix)
	var (
		installers = []*Installer{
			NewBitcoinInstaller("amd64", "linux", prefix, LAN),
			NewLNDInstaller("amd64", "linux", prefix, LAN),
		}
	)
	for _, i := range installers {
		if err := i.Install(); err != nil {
			t.Fatal(err)
		}
		if !Exists(i.GzDir()) {
			t.Errorf("directory %s expect to exist", i.GzDir())
		}
		if !Exists(i.GzPath()) {
			t.Errorf("file %s expect to exist", i.GzPath())
		}
		for _, c := range i.commands {
			if !Exists(filepath.Join(i.prefix, "bin", filepath.Base(c))) {
				t.Errorf("%s does not exist in prefix/bin", c)
			}
		}
		if !Exists(i.GzPath()) {
			t.Errorf("file %s expect to exist", i.GzPath())
		}
	}
	for _, i := range installers {
		if err := i.UnInstall(); err != nil {
			t.Fatal(err)
		}
		for _, c := range i.commands {
			file := filepath.Join(i.prefix, "bin", filepath.Base(c))
			if Exists(file) {
				t.Errorf("%s should not exist", file)
			}
		}
	}
}
