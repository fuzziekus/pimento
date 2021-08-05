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
		if !listFlags.NoHeader {
			listFlags.DisplayHeader()
		}

		for _, c := range credentials {
			listFlags.DisplaySpecifyColumn(c)
		}
	},
}

func init() {
	// fmt.Println(config.Mgr().Db.Path)
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&listFlags.NoHeader, "noheader", "n", false, "list with no header")
	listCmd.Flags().BoolVarP(&listFlags.NoPass, "all", "a", true, "list all column")
	listCmd.Flags().BoolVarP(&listFlags.Description, "description", "d", false, "list description")
	listCmd.Flags().BoolVarP(&listFlags.UserId, "user_id", "u", false, "list user_id")
	listCmd.Flags().BoolVarP(&listFlags.Password, "password", "p", false, "list password")
	listCmd.Flags().BoolVarP(&listFlags.Memo, "memo", "m", false, "list memo")
}
