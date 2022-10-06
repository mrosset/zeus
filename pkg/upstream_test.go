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
	"testing"
)

const uri = "http://10.119.176.16/bitcoin-23.0-x86_64-linux-gnu.tar.gz"

func TestFetchExtract(t *testing.T) {
	tmp, err := ioutil.TempDir(os.TempDir(), "raijin")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmp)
	gzfile, err := Fetch(tmp, uri)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(gzfile)
	if !Verify(gzfile) {
		t.Error("File could not be verifed")
	}
	index, err := TarDir(gzfile)
	if err != nil {
		t.Fatal(err)
	}
	if index != "bitcoin-23.0" {
		t.Errorf("Expected tar index bitcoin-23.0 got %s", index)
	}
	if err = Extract(tmp, gzfile); err != nil {
		t.Fatal(err)
	}
	expect := filepath.Join(tmp, "bitcoin-23.0", "README.md")
	if !Exists(expect) {
		t.Errorf("File %s does not", expect)
	}
}
