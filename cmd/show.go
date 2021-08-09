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

var showFlags ui.Flags

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Args:  cobra.MinimumNArgs(1),
	Short: "対象のクレデンシャルを1件取得する",
	Long: `対象のクレデンシャルを1件取得する.
	デフォルトではパスワードは表示しない`,
	Run: func(cmd *cobra.Command, args []string) {
		c := db.NewCredentialRepository().GetSingleRowByItemName(args[0])
		showFlags.DisplayRow(c)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().BoolVarP(&showFlags.NoPass, "all", "a", true, "show all column without password")
	showCmd.Flags().BoolVarP(&showFlags.All, "all-with-password", "", false, "show all column with password")
	showCmd.Flags().BoolVarP(&showFlags.ItemName, "item-name", "i", false, "show item name")
	showCmd.Flags().BoolVarP(&showFlags.UserName, "user-name", "u", false, "show user name")
	showCmd.Flags().BoolVarP(&showFlags.Password, "password", "p", false, "show password")
	showCmd.Flags().BoolVarP(&showFlags.Tag, "tag", "t", false, "show tag")
	showCmd.Flags().BoolVarP(&showFlags.FormatCSV, "csv", "c", false, "output fomat csv")
	showCmd.Flags().BoolVarP(&showFlags.FormatTable, "table", "", true, "output fomat ascii table")
}
