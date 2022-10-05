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
	"log"
	"os"
	"path/filepath"

	. "github.com/mrosset/raijin/pkg"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "raijin",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.raijin.yaml)")
	rootCmd.PersistentFlags().String("prefix", filepath.Join(home, "bitcoin"), "Directory Bitcoin Core is to be installed")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Checks that the prefix flag path exists. If it does not logs a
// fatal error.
func checkPrefix(cmd *cobra.Command) {
	var (
		prefix = prefixFlag(cmd)
	)
	if !Exists(prefix) {
		log.Fatalf("%s prefix does not exists. Have you run `raijin install`?", prefix)
	}
}

// Returns the absolute path for bitcoind using prefix flag
func bitcoindCmd(cmd *cobra.Command) string {
	return filepath.Join(prefixFlag(cmd), "bin", "bitcoind")
}

// Returns full path of bitcoin.conf using prefix flag
func configFile(cmd *cobra.Command) string {
	return filepath.Join(prefixFlag(cmd), "bitcoin.conf")
}

// Returns the full path of data directory using prefix flag
func dataDir(cmd *cobra.Command) string {
	return filepath.Join(prefixFlag(cmd), "data")
}

// Returns the prefix local flag. If an error occurs logs fatal error.
func prefixFlag(cmd *cobra.Command) string {
	prefix, err := cmd.Flags().GetString("prefix")
	if err != nil {
		log.Fatal(err)
	}
	return prefix
}
