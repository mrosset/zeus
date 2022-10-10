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
	"os/exec"
	"path/filepath"
)

// Creates a new bitcoind Cmd with binary from prefix
func NewBitcoind(prefix string) *exec.Cmd {
	var (
		bin = filepath.Join(prefix, "bin", "bitcoind")
	)
	return &exec.Cmd{
		Path: bin,
		Args: []string{
			bin,
			"-conf=" + filepath.Join(prefix, "bitcoin.conf")},
		Stdout: os.Stdout,
		Stderr: os.Stderr}
}
