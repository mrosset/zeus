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
	"context"
	"fmt"
	"github.com/codeclysm/extract"
	"github.com/mrosset/gurl"
	"os"
	"path"
)

const (
	amd64_linux_gnu = "2CCA490C1F2842884A3C5B0606F179F9F937177DA4EADD628E3F7FD7E25D26D0"
)

// Download URI to DIR path. Returns downloaded file path
func Fetch(dir, uri string) (string, error) {
	if !Exists(dir) {
		return "", fmt.Errorf("Directory %s does not exist", dir)
	}
	file := path.Join(dir, path.Base(uri))
	if err := gurl.Download(dir, uri); err != nil {
		return "", err
	}
	return file, nil
}

// Verify the sha256 sum for PATH. Returns true if verification
// passes. False if it does not pass
//
// TODO: add support for other OS and ARCHs
func Verify(path string) bool {
	hash, err := Sha256sum(path)
	if err != nil {
		return false
	}
	if hash == amd64_linux_gnu {
		return true
	}
	fmt.Println(hash, amd64_linux_gnu)
	return false
}

// Extracts tarball FILE to DIR. Returns error
func Extract(dir, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return extract.Gz(context.Background(), f, dir, nil)
}
