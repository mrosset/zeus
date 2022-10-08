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
	Description string
	hash        string
	commands    []string
	arch        string
	os          string
	prefix      string
	uri         string
}

type BitcoindInstaller Installer

// Returns the full path of the bitcoind config file.
func (i BitcoindInstaller) Config() string {
	return filepath.Join(i.prefix, "bitcoin.conf")
}

func (i BitcoindInstaller) PostInstall() error {
	// FIXME: Prompt before overwriting bitcoind.config
	if err := NewDefaultConfig(i.prefix).Write(i.Config()); err != nil {
		return err
	}
	// if !Exists(data) {
	//	os.Mkdir(data, 0755)
	// }
	return nil
}

func NewBitcoinInstaller(arch, os, prefix string, release MirrorType) BitcoindInstaller {
	var (
		uri = LAN_URI
	)
	if release == WEB {
		uri = fmt.Sprintf(BITCOIN_URI, BITCOIN_VERSION)

	}
	entry, err := bitcoinUpstream.Entry(arch, os)
	if err != nil {
		panic(err)
	}
	return BitcoindInstaller{
		Description: "Bitcoin Core",
		arch:        arch,
		os:          os,
		prefix:      prefix,
		commands: []string{
			"bin/test_bitcoin",
			"bin/bitcoind",
			"bin/bitcoin-wallet",
			"bin/bitcoin-qt",
			"bin/bitcoin-tx",
			"bin/bitcoin-util",
			"bin/bitcoin-cli"},
		hash: entry.Hash,
		uri:  fmt.Sprintf("%s/%s", uri, entry.File)}
}

func NewLNDInstaller(arch, os, prefix string, release MirrorType) Installer {
	var (
		uri = LAN_URI
	)
	if release == WEB {
		uri = fmt.Sprintf(LND_URI, LND_VERSION)
	}
	entry, err := lndUpstream.Entry(arch, os)
	if err != nil {
		panic(err)
	}
	return Installer{
		Description: "Lightning Network Daemon",
		arch:        arch,
		os:          os,
		prefix:      prefix,
		commands: []string{
			"lnd",
			"lncli"},
		hash: entry.Hash,
		uri:  fmt.Sprintf("%s/%s", uri, entry.File)}
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
		return Fetch(i.GzDir(), i.uri)
	}
	return nil
}

func (i *Installer) Verify() error {
	var (
		tarBall = i.GzPath()
	)
	if !Verify(tarBall, i.hash) {
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
	index, err := TarIndex(tarBall)
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
		src := filepath.Join(tarDir, e)
		dest := filepath.Join(i.prefix, "bin", filepath.Base(e))
		if err := os.Rename(src, dest); err != nil {
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
	for _, c := range i.commands {
		file := filepath.Join(i.prefix, "bin", filepath.Base(c))
		os.Remove(file)
	}
	return nil
}
