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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "クレデンシャル情報を追加で保存します",
	Long:  `クレデンシャル情報を追加で保存します`,
	Run: func(cmd *cobra.Command, args []string) {
		editflags := ui.Flags{}
		editflags.All = true
		credential := editflags.GenerateCredentialFromInputInterfaces()
		db.NewCredentialRepository().Create(credential)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
