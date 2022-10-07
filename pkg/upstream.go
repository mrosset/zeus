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
	"os"

	"github.com/codeclysm/extract"
	"github.com/mrosset/gurl"
)

const (
	BITCOIN_VERSION = "23.0"
	BITCOIN_URI     = "https://bitcoincore.org/bin/bitcoin-core-%s"
	LND_VERSION     = "0.15.1-beta"
	LND_URI         = "https://github.com/lightningnetwork/lnd/releases/download/%s"
	LAN_URI         = "http://10.119.176.16"
)

type MirrorType int

const (
	LAN MirrorType = iota
	WEB
)

type Tarball struct {
	Hash   string
	File   string
	TarDir string
}

type UpstreamFiles map[string]map[string]Tarball

// https://bitcoincore.org/bin/bitcoin-core-23.0/bitcoin-23.0-x86_64-linux-gnu.tar.gz
var bitcoinUpstream = UpstreamFiles{
	"amd64": {"linux": Tarball{
		Hash:   "2CCA490C1F2842884A3C5B0606F179F9F937177DA4EADD628E3F7FD7E25D26D0",
		TarDir: fmt.Sprintf("bitcoin-%s", BITCOIN_VERSION),
		File:   fmt.Sprintf("bitcoin-%s-x86_64-linux-gnu.tar.gz", BITCOIN_VERSION)}}}

// https://github.com/lightningnetwork/lnd/releases/download/v0.15.1-beta/lnd-linux-amd64-v0.15.1-beta.tar.gz
var lndUpstream = UpstreamFiles{
	"amd64": {"linux": Tarball{
		Hash:   "0673768E657AC004367D07C20395D544A3D1DF926BE1A1990A17E23A8A91D4FB",
		TarDir: fmt.Sprintf("lnd-linux-amd64-v%s", LND_VERSION),
		File:   fmt.Sprintf("lnd-linux-amd64-v%s.tar.gz", LND_VERSION)}}}

// Returns the sha256 hash for ARCH and OS
func BitcoinHash(arch, os string) string {
	return bitcoinUpstream[arch][os].Hash
}

// Download URI to DIR path. Returns downloaded file path
func Fetch(dir, uri string) error {
	if err := gurl.Download(dir, uri); err != nil {
		return err
	}
	return nil
}

// Verify the HASH for PATH. Returns true if verification
// passes. False if it does not pass
func Verify(path, hash string) bool {
	got, err := Sha256sum(path)
	if err != nil {
		return false
	}
	if got == hash {
		return true
	}
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
