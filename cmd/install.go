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
package cmd

import (
	"runtime"

	. "github.com/mrosset/zeus/pkg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs various bitcoin related daemons including Bitcoin Core",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) { },

}

func init() {
	var (
		bitcoinCmd = &cobra.Command{
			Use:   "bitcoin",
			Short: "Installs Bitcoin Core",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Run: installBitcoin,
		}
		lndCmd = &cobra.Command{
			Use:   "lnd",
			Short: "Installs The Lighting Network Daemon",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Run: installLnd,
		}
		allCmd = &cobra.Command{
			Use:   "all",
			Short: "Installs both Bitcoin Core and The Lighting Network Daemon",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
			Run: installAll,
		}
	)

	rootCmd.AddCommand(installCmd)
	installCmd.AddCommand(bitcoinCmd)
	installCmd.AddCommand(lndCmd)
	installCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	installCmd.PersistentFlags().BoolP("debug", "d", false, "If used, uses debug URI for downloads")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func mirror(cmd *cobra.Command) MirrorType {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		pterm.Fatal.Println(err)
	}
	switch debug {
	case true:
		return LAN
	default:
		return WEB
	}
}

func lighting(cmd *cobra.Command) bool {
	flag, err := cmd.Flags().GetBool("lighting")
	if err != nil {
		pterm.Fatal.Println(err)
	}
	return flag
}

func doInstall(i Installer) {
	printLogo()
	pterm.Info.Println("Installing:", i.Description)
	if err := i.Install(); err != nil {
		pterm.Fatal.Println(err)
	}
}

func installBitcoin(cmd *cobra.Command, args []string) {
	var (
		prefix   = prefixFlag(cmd)
		bitcoind = NewBitcoinInstaller(runtime.GOARCH, runtime.GOOS, prefix, mirror(cmd))
	)
	doInstall(Installer(bitcoind))
	if err := bitcoind.PostInstall(); err != nil {
		pterm.Fatal.Println(err)
	}
	pterm.Info.Println("Wrote: default config file:", bitcoind.Config())
}

func installLnd(cmd *cobra.Command, args []string) {
	var (
		prefix = prefixFlag(cmd)
		lnd    = NewLNDInstaller(runtime.GOARCH, runtime.GOOS, prefix, mirror(cmd))
	)
	doInstall(Installer(lnd))
}

func installAll(cmd *cobra.Command, args []string) {
	installBitcoin(cmd, args)
	installLnd(cmd, args)
}
