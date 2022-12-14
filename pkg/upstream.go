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
	"context"
	"fmt"
	"os"

	"github.com/codeclysm/extract"
	"github.com/mrosset/gurl"
)

const (
	BITCOIN_VERSION = "23.0"
	LND_VERSION     = "0.15.3-beta"

	BITCOIN_URI = "https://bitcoincore.org/bin/bitcoin-core-%s"
	LND_URI     = "https://github.com/lightningnetwork/lnd/releases/download/v%s"
	LAN_URI     = "http://devel"
)

type MirrorType int

const (
	LAN MirrorType = iota
	WEB
)

type Tarball struct {
	Hash string
	File string
}

type UpstreamFiles map[string]map[string]Tarball

func (u UpstreamFiles) Entry(arch, os string) (Tarball, error) {
	var (
		entry, ok = u[arch][os]
	)

	if !ok {
		return entry, fmt.Errorf("Could not find Tarball entry for arch: %s os: %s", arch, os)
	}
	return entry, nil
}

var bitcoinUpstream = UpstreamFiles{
	"amd64": {"linux": Tarball{
		Hash: "2CCA490C1F2842884A3C5B0606F179F9F937177DA4EADD628E3F7FD7E25D26D0",
		File: fmt.Sprintf("bitcoin-%s-x86_64-linux-gnu.tar.gz", BITCOIN_VERSION)}},
	"ppc64le": {"linux": Tarball{
		Hash: "217DD0469D0F4962D22818C368358575F6A0ABCBA8804807BB75325EB2F28B19",
		File: fmt.Sprintf("bitcoin-%s-powerpc64le-linux-gnu.tar.gz", BITCOIN_VERSION)}},
	"arm64": {"linux": Tarball{
		Hash: "06F4C78271A77752BA5990D60D81B1751507F77EFDA1E5981B4E92FD4D9969FB",
		File: fmt.Sprintf("bitcoin-%s-aarch64-linux-gnu.tar.gz", BITCOIN_VERSION)}},
}

var lndUpstream = UpstreamFiles{
	"amd64": {"linux": Tarball{
		Hash: "1F7903C8F700860502D0E7D369130F86DC43E80B0887CC04D7DBEEC3122DBF50",
		File: fmt.Sprintf("lnd-linux-amd64-v%s.tar.gz", LND_VERSION)}},
	"ppc64le": {"linux": Tarball{
		Hash: "0673768E657AC004367D07C20395D544A3D1DF926BE1A1990A17E23A8A91D4FB",
		File: fmt.Sprintf("lnd-linux-amd64-v%s.tar.gz", LND_VERSION)}},
	"arm64": {"linux": Tarball{
		Hash: "8A1EE8B000173C5F938BB6B47F1E6BF95A02523128304E2CB420C98182094296",
		File: fmt.Sprintf("lnd-linux-arm64-v%s.tar.gz", LND_VERSION)}},
}

// Download URI to DIR path. Returns downloaded file path
func Fetch(dir, uri string) error {
	return gurl.DownloadHideAfter(dir, uri)
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
