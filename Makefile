# Copyright © 2022 Michael Rosset <mike.rosset@gmail.com>

# This file is part of Raijin

# Raijin is free software: you can redistribute it and/or modify it
# under the terms of the GNU General Public License as published by the
# Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# Raijin is distributed in the hope that it will be useful, but
# WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
# See the GNU General Public License for more details.

# You should have received a copy of the GNU General Public License along
# with this program.  If not, see <http://www.gnu.org/licenses/>.

OUT = zeus
CMD = ./$(OUT) --prefix=$(PWD)/temp

.NOTPARALLEL:
.PHONY: $(OUT)

$(OUT):
	go build -v -o $(OUT)
	strip ./$@

start: $(OUT)
	$(CMD) $@

test-uinstall: $(OUT)
	$(CMD) uninstall

test-install: check $(OUT) test-uinstall
	$(CMD)
	$(CMD) install -d -l

check:
	$(MAKE) -C ./pkg

install:
	go install

clean:
	-$(CMD) uninstall
	-rm $(OUT)
	-rm temp/bitcoin.conf
