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
	"runtime"

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

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	installCmd.Flags().StringP("release", "r", "debug", "Specify to use debug or release URI")
}

var tarEntries = []string{"include", "lib", "bin", "share", "README.md"}

func releaseFlag(cmd *cobra.Command) string {
	release, err := cmd.Flags().GetString("release")
	if err != nil {
		log.Fatal(err)
	}
	return release
}

func install(cmd *cobra.Command, args []string) {
	var (
		uri      = BitcoinUri(runtime.GOARCH, runtime.GOOS, releaseFlag(cmd))
		prefix   = prefixFlag(cmd)
		gzDir    = filepath.Join(prefix, "gz")
		tarBall  = filepath.Join(gzDir, filepath.Base(uri))
		confFile = configFile(cmd)
		data     = dataDir(cmd)
	)
	if Exists(bitcoindCmd(cmd)) {
		log.Fatalf("Bitcoin already installed in %s", prefix)
	}
	fmt.Printf("Installing Bitcoin Core to: %s\n", prefix)
	if !Exists(gzDir) {
		os.MkdirAll(gzDir, 0775)
	}
	if !Exists(tarBall) {
		fmt.Println("Downloading:", uri)
		_, err := Fetch(gzDir, uri)
		if err != nil {
			log.Fatal(err)
		}

	}
	if !Verify(tarBall, BitcoinHash(runtime.GOARCH, runtime.GOOS)) {
		os.Remove(tarBall)
		log.Fatalf("Could not verify sha256 sum for %s", tarBall)
	} else {
		fmt.Println("Verified:", filepath.Base(tarBall), "OK")
	}
	index, err := TarDir(tarBall)
	if err != nil {
		log.Fatal(err)
	}
	tarDir := filepath.Join(prefix, index)
	if !Exists(tarDir) {
		fmt.Println("Extracting:", tarBall, " -> ", prefix)
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
	fmt.Println("Wrote default config file:", confFile)

	if err := NewDefaultConfig(prefix).Write(confFile); err != nil {
		log.Fatal(err)
	}
	if !Exists(data) {
		os.Mkdir(data, 0755)
	}
}
