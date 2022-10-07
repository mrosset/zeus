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
	"fmt"
	"os"
	"path/filepath"
)

const (
	BITCOIN InstallType = iota
	LIGHTING
)

type InstallType int

type Installer struct {
	hash     string
	commands []string
	arch     string
	os       string
	prefix   string
	uri      string
}

func NewBitcoinInstaller(arch, os, prefix string, release MirrorType) *Installer {
	return &Installer{
		arch:   arch,
		os:     os,
		prefix: prefix,
		commands: []string{
			"test_bitcoin",
			"bitcoind",
			"bitcoin-wallet",
			"bitcoin-qt",
			"bitcoin-tx",
			"bitcoin-util",
			"bitcoin-cli"},
		hash: BitcoinHash(arch, os),
		uri:  BitcoinUri(arch, os, release)}
}

func (i *Installer) GzDir() string {
	return filepath.Join(i.prefix, "gz")
}

func (i *Installer) GzPath() string {
	return filepath.Join(i.GzDir(), filepath.Base(i.uri))
}

func (i *Installer) Fetch() error {
	var (
		file = i.GzPath()
	)
	if !Exists(i.GzDir()) {
		os.MkdirAll(i.GzDir(), 0700)
	}
	if !Exists(file) {
		err := Fetch(i.GzDir(), i.uri)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Installer) Verify() error {
	var (
		tarBall = i.GzPath()
	)
	if !Verify(tarBall, BitcoinHash(i.arch, i.os)) {
		os.Remove(tarBall)
		return fmt.Errorf("Could not verify sha256 sum for %s", tarBall)
	}
	return nil
}

func (i *Installer) Extract() error {
	var (
		tarBall = i.GzPath()
		binDir  = filepath.Join(i.prefix, "bin")
	)
	index, err := TarDir(tarBall)
	if err != nil {
		return err
	}
	tarDir := filepath.Join(i.prefix, index)
	if !Exists(tarDir) {
		if err := Extract(i.prefix, tarBall); err != nil {
			return err
		}
	}
	defer os.RemoveAll(tarDir)
	if !Exists(binDir) {
		os.Mkdir(binDir, 0755)
	}
	for _, e := range i.commands {
		if err := os.Rename(filepath.Join(tarDir, "bin", e), filepath.Join(binDir, e)); err != nil {
			return err
		}
	}
	return nil
}

func (i *Installer) Install() error {
	if err := i.Fetch(); err != nil {
		return err
	}
	if err := i.Verify(); err != nil {
		return err
	}
	if err := i.Extract(); err != nil {
		return err
	}
	return nil
}

func (i *Installer) UnInstall() error {
	for _, e := range i.commands {
		os.Remove(filepath.Join(i.prefix, "bin", e))
	}
	return nil
}
