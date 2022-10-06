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
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	. "github.com/mrosset/raijin/pkg"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs bitcoin core",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: install,
}

const (
	prodUri = "https://bitcoincore.org/bin/bitcoin-core-23.0/bitcoin-23.0-x86_64-linux-gnu.tar.gz"
	devUri  = "http://10.119.176.16/bitcoin-23.0-x86_64-linux-gnu.tar.gz"
)

var bitcoinUri = devUri

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var tarEntries = []string{"include", "lib", "bin", "share", "README.md"}

func install(cmd *cobra.Command, args []string) {
	var (
		prefix   = prefixFlag(cmd)
		gzDir    = filepath.Join(prefix, "gz")
		tarBall  = filepath.Join(gzDir, filepath.Base(bitcoinUri))
		confFile = configFile(cmd)
		data     = dataDir(cmd)
	)
	if Exists(bitcoindCmd(cmd)) {
		log.Fatalf("Bitcoin already installed in %s", prefix)
	}
	fmt.Printf("Installing Bitcoin Core to %s\n", prefix)
	if !Exists(gzDir) {
		os.MkdirAll(gzDir, 0775)
	}
	if !Exists(tarBall) {
		_, err := Fetch(gzDir, bitcoinUri)
		if err != nil {
			log.Fatal(err)
		}

	}
	if !Verify(tarBall) {
		os.Remove(tarBall)
		log.Fatalf("Could not verify sha256 sum for %s", tarBall)
	}
	index, err := TarDir(tarBall)
	if err != nil {
		log.Fatal(err)
	}
	tarDir := filepath.Join(prefix, index)
	if !Exists(tarDir) {
		if err := Extract(prefix, tarBall); err != nil {
			log.Fatal(err)
		}
	}
	defer os.RemoveAll(tarDir)
	for _, e := range tarEntries {
		if err := os.Rename(filepath.Join(tarDir, e), filepath.Join(prefix, e)); err != nil {
			log.Fatal(err)
		}
	}
	// TODO: Prompt before overwriting bitcoind.config
	fmt.Println("Writing default config file.")

	if err := NewDefaultConfig(prefix).Write(confFile); err != nil {
		log.Fatal(err)
	}
	if !Exists(data) {
		os.Mkdir(data, 0755)
	}
}
