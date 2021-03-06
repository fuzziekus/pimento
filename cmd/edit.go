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
	"time"

	"github.com/fuzziekus/pimento/db"
	"github.com/fuzziekus/pimento/ui"
	"github.com/spf13/cobra"
)

var editflags ui.Flags

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [DESCRIPTION]",
	Args:  cobra.MinimumNArgs(1),
	Short: "対象のクレデンシャルを更新する",
	Long:  `対象のクレデンシャルを更新する`,
	Run: func(cmd *cobra.Command, args []string) {
		r := db.NewCredentialRepository()
		c := r.GetSingleRowByItemName(args[0])
		newcredential := editflags.GenerateCredentialFromInputInterfaces()
		newcredential.UpdateAt.Time = time.Now()
		newcredential.UpdateAt.Valid = true
		r.UpdateRow(c.ID, newcredential)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().BoolVarP(&editflags.All, "all", "a", true, "edit all column")
	editCmd.Flags().BoolVarP(&editflags.ItemName, "item-name", "i", false, "edit item name")
	editCmd.Flags().BoolVarP(&editflags.UserName, "user-name", "u", false, "edit user name")
	editCmd.Flags().BoolVarP(&editflags.Password, "password", "p", false, "edit password")
	editCmd.Flags().BoolVarP(&editflags.Tag, "tag", "t", false, "edit tag")
}
