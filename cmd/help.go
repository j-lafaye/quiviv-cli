/*
Copyright © 2022 Julien Lafaye lafaye.julien92@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Displays a brief description of quiviv-cli and its available commands",
	Long: `Quiviv-cli is a CLI made in Go that empowers applications.
This command returns a list of available commands with their optional flags in quiviv-cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Usage:
	quiviv-cli [command]

Available Commands:
	config      Displays a JSON formatted configuration file from a specified path
	help        Help about any command
	list        List all the contents of a directory

Flags:
	--dir string    directory name for content listing (optionnal)
	--path string   path for config file display`)
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
