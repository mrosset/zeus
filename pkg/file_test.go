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

var (
	testTarFile = "testdata/test-0.1.0.tar.gz"
)

func TestTarIndex(t *testing.T) {
	var (
		got, err = TarIndex(testTarFile)
		expect   = "top"
	)
	if err != nil {
		t.Fatal(err)
	}
	if expect != got {
		t.Errorf("Expect %s got %s", expect, got)
	}
}
