/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"io/ioutil"
	"log"

	"github.com/fuzziekus/pimento/config"
	"github.com/fuzziekus/pimento/db"
	"github.com/jszwec/csvutil"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import [CSVFILE]",
	Args:  cobra.MinimumNArgs(1),
	Short: "CSVからクレデンシャルを取り込む",
	Long:  `CSVからクレデンシャルを取り込む`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(args[0])
		cobra.CheckErr(err)

		var credentials db.Credentials
		cobra.CheckErr(csvutil.Unmarshal(b, &credentials))
		for _, c := range credentials {
			cipertext, err := config.RowCryptor.Encrypt(c.Password)
			if err != nil {
				log.Fatal(err)
			}
			c.Password = string(cipertext)
			db.NewCredentialRepository().Create(c)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
