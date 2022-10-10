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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

var (
	server = httptest.NewServer(&testHandler{})
)

type testHandler struct {
}

func NewTestInstaller(arch, os, prefix string) *Installer {
	return &Installer{
		Description: "Test Installer",
		arch:        arch,
		os:          os,
		prefix:      prefix,
		commands: []string{
			"bin/lnd"},
		hash: "4906F5BC7569674997A08D801809A6D3F829891298EF45331FF2FFB12BCE6325",
		uri:  fmt.Sprintf("%s/%s", server.URL, testTarFile)}
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	size, err := Size(testTarFile)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(testTarFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Content-Length", strconv.Itoa(int(size)))
	io.Copy(w, file)
}

func TestInstallerType(t *testing.T) {
	var (
		gzExpect = "temp/gz"
		gzPath   = "temp/gz/bitcoin-23.0-x86_64-linux-gnu.tar.gz"
		got      = Installer(NewBitcoinInstaller("amd64", "linux", "temp", LAN))
		expect   = Installer{
			Description: "Bitcoin Core",
			uri:         "http://10.119.176.16/bitcoin-23.0-x86_64-linux-gnu.tar.gz",
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
				"bin/bitcoin-cli"}}
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
	prefix, err := ioutil.TempDir(os.TempDir(), "zeus")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(prefix)
	i := NewTestInstaller(runtime.GOARCH, runtime.GOOS, prefix)
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
			t.Errorf("%s does not exist in %s/bin", c, i.prefix)
		}
	}
}
