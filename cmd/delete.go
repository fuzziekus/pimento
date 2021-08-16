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
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [DESCRIPTION]",
	Args:  cobra.MinimumNArgs(1),
	Short: "保存されたクレデンシャル情報を削除します",
	Long:  `保存されたクレデンシャル情報を削除します`,
	Run: func(cmd *cobra.Command, args []string) {
		r := db.NewCredentialRepository()
		c := r.GetSingleRowByItemName(args[0])
		r.DeleteRow(c.ID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
