/*
Copyright © 2022 Julien Lafaye lafaye.julien92@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"io/ioutil"
)

var Dir string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Dislays a list of a directory elements",
	Long: `Reads a directory content, and displays its elements in a list. The selection of a specific directory can be set by the flag --dir, if not set command will then read current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			WARNING:
			Usage of path starting by "/" will not target the root of quivi-cli directory but the Golang's config root path ($HOME)
				"/*" = C:/Users/PARTUM~1/AppData/Local/Temp/*
		*/
		if Dir == "" {
			Dir = "./"
		}

		files, err := readDir(Dir)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			fmt.Println(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&Dir, "dir", "d", "", "Dir (not required)")
}

func readDir(root string) ([]string, error) {
	var files []string
	
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
			return files, err
	}

	for _, file := range fileInfo {
			files = append(files, file.Name())
	}

	return files, nil
}
