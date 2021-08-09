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
	"github.com/fuzziekus/pimento/db"
	"github.com/fuzziekus/pimento/ui"
	"github.com/spf13/cobra"
)

var listFlags ui.Flags

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "保存しているクレデンシャルの一覧を表示します.",
	Long:  `保存しているクレデンシャルの一覧を表示します. ※PWは表示されません`,
	Run: func(cmd *cobra.Command, args []string) {
		credentials := db.NewCredentialRepository().GetAll()
		listFlags.DisplayRows(credentials)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&listFlags.NoHeader, "noheader", "n", false, "list with no header")
	listCmd.Flags().BoolVarP(&listFlags.NoPass, "all", "a", true, "list all column without password")
	listCmd.Flags().BoolVarP(&listFlags.All, "all-with-password", "", false, "list all column with password")
	listCmd.Flags().BoolVarP(&listFlags.ItemName, "item-name", "i", false, "list item name")
	listCmd.Flags().BoolVarP(&listFlags.UserName, "user-name", "u", false, "list user name")
	listCmd.Flags().BoolVarP(&listFlags.Password, "password", "p", false, "list password")
	listCmd.Flags().BoolVarP(&listFlags.Tag, "tag", "t", false, "list tag")
	listCmd.Flags().BoolVarP(&listFlags.FormatCSV, "csv", "c", false, "output fomat csv")
	listCmd.Flags().BoolVarP(&listFlags.FormatTable, "table", "", true, "output fomat ascii table")
}
