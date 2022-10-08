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
	"log"
	"os"
	"runtime"

	. "github.com/mrosset/raijin/pkg"
	"github.com/pterm/pterm"
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
	installCmd.Flags().BoolP("debug", "d", true, "If false use debug URI for downloads")
}

func install(cmd *cobra.Command, args []string) {
	var (
		confFile   = configFile(cmd)
		data       = dataDir(cmd)
		prefix     = prefixFlag(cmd)
		installers = []*Installer{
			NewBitcoinInstaller(runtime.GOARCH, runtime.GOOS, prefix, WEB),
			NewLNDInstaller(runtime.GOARCH, runtime.GOOS, prefix, WEB),
		}
	)
	for _, i := range installers {
		pterm.Info.Println("Installing:", i.Description)
		if err := i.Install(); err != nil {
			log.Fatal(err)
		}
	}
	// TODO: Prompt before overwriting bitcoind.config
	if err := NewDefaultConfig(prefix).Write(confFile); err != nil {
		log.Fatal(err)
	}
	pterm.Info.Println("Wrote: default config file:", confFile)
	if !Exists(data) {
		os.Mkdir(data, 0755)
	}
}
