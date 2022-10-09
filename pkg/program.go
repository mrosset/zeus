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

package zeus

import (
	"bytes"
	"strings"
	"text/template"
)

const (
	ShortName = "zeus"
)

type Program struct {
	ShortName  string
	TitledName string
}

// The default program struct
var DefaultProgram = Program{ShortName, TitledName()}

// Returns a Titled string for ShortName
func TitledName() string {
	return strings.Title(ShortName)
}

// Returns the template INPUT as a strings
func ParseProgram(input string) (out string) {
	var (
		buf bytes.Buffer
	)
	defer func() {
		if recover() != nil {
			out = "Error parsing input"
		}
	}()
	t, err := template.New("program").Parse(input)
	if err != nil {
		panic(err)
	}
	err = t.Execute(&buf, DefaultProgram)
	if err != nil {
		panic(err)
	}
	out = buf.String()
	return out
}
