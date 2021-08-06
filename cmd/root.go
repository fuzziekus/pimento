/*
Copyright © 2021 fuzziekus

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/fuzziekus/pimento/config"
	"github.com/fuzziekus/pimento/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pimento",
	Short: "TUI PWマネージャ",
	Long:  `pimento TUI PWマネージャ管理ツール`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())

	// 暗号化前にDBのコネクションをクローズ
	db.Mgr().Close()
	// DB を外部から使用できないよう暗号化する
	cobra.CheckErr(config.DbCryptor.EncryptFile(config.Mgr().Db.Path+".temp", config.Mgr().Db.Path))

	// 一時的に復号していたDBファイルを削除
	if err := os.Remove(config.Mgr().Db.Path + ".temp"); err != nil {
		cobra.CheckErr(err)
	}
	// 初回起動の場合、コンフィグを設定
	config.WriteConfig(false)
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/pimento/config.yaml)")
	rootCmd.PersistentFlags().String("secret_key", "", "pimento secret key")
	viper.BindPFlag("pimento_secret_key", rootCmd.PersistentFlags().Lookup("secret_key"))

}
