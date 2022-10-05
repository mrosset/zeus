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
	"os"
	"testing"
)

var config = NewDefaultConfig("./testdata")

func TestConfig(t *testing.T) {
	var (
		expect = "testdata/data"
		got    = config.DataDir
	)
	if got != expect {
		t.Errorf("expect %s got %s", expect, got)
	}
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

func TestRead(t *testing.T) {
	var (
		file   = "testdata/default.conf"
		expect = "testdata/data"
	)
	config, err := ReadConfig(file)
	if err != nil {
		t.Fatal(err)
	}
	if expect != config.DataDir {
		t.Fatalf("Expected %s got %s", expect, config.DataDir)
	}
}
