/*
Copyright © 2022 Julien Lafaye lafaye.julien92@gmail.com

*/
package cmd

import (
	"fmt"

	"encoding/json"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Path string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Dislays the quiviv-cli config file",
	Long: `Reads the Yaml format quiviv-cli config file and displays it in output in JSON format. Full path of file required with flag --path.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			WARNING:
			Usage of path starting by "/" will not target the root of quivi-cli directory but the Golang's config root path ($HOME)
				"/*" = C:/Users/PARTUM~1/AppData/Local/Temp/*
		*/
		file, err := ioutil.ReadFile(Path)
		if err != nil {
			panic(err)
		}

		var body interface{}
    if err := yaml.Unmarshal([]byte(file), &body); err != nil {
			panic(err)
    }

		/*
		body is now unmarshaled without any structure
			body = map[employed:true foods:[Pomme Orange Fraise Mangue] job:Développeur languages:map[pascal:Bases perl:Élite python:Élite] name:Martin Développeur skill:Élite]
		*/

    body = convert(body)

    if b, err := json.MarshalIndent(body, "", "    "); err != nil {
			panic(err)
    } else {
			fmt.Printf("%s\n", b)
    }
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&Path, "path", "p", "", "Path (required)")

	configCmd.MarkFlagRequired("path")
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
		case map[interface{}]interface{}:
				m2 := map[string]interface{}{}
				for k, v := range x {
					m2[k.(string)] = convert(v)
				}
			return m2

		case []interface{}:
				for i, v := range x {
					x[i] = convert(v)
				}
	}

	return i
}