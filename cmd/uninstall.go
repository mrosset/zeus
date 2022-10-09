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
package cmd

import (
	. "github.com/mrosset/raijin/pkg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"runtime"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls Bitcoin from --prefix",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: uninstall,
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func uninstall(cmd *cobra.Command, args []string) {
	var (
		prefix     = prefixFlag(cmd)
		installers = []Installer{
			Installer(NewBitcoinInstaller(runtime.GOARCH, runtime.GOOS, prefix, LAN)),
			NewLNDInstaller(runtime.GOARCH, runtime.GOOS, prefix, LAN),
		}
	)
	for _, i := range installers {
		pterm.Info.Println("Uninstalling:", i.Description)
		if err := i.UnInstall(); err != nil {
			pterm.Fatal.Println(err)
		}
	}
}
