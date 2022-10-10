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
	"testing"
)

func TestEntry(t *testing.T) {
	entry, err := bitcoinUpstream.Entry("amd64", "linux")
	if err != nil {
		t.Fatal(err)
	}
	var (
		expect = "2CCA490C1F2842884A3C5B0606F179F9F937177DA4EADD628E3F7FD7E25D26D0"
		got    = entry.Hash
	)
	if expect != got {
		t.Errorf("Expected Hash %s got %s", expect, got)
	}
	entry, err = bitcoinUpstream.Entry("amd64", "bork")
	if err == nil {
		t.Errorf("Entry() should fail")
	}
}
