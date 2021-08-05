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
	"github.com/fuzziekus/pimento/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pimento",
	Short: "TUI PWマネージャ",
	Long:  `pimento TUI PWマネージャ管理ツール`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports  flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/pimento/config.yaml)")

	// rootCmd.PersistentFlags().StringVarP(&config.Cfg.Secret_key, "secret_key", "", "", "pimento secret key")
	// rootCmd.PersistentFlags().StringVar(&config.Cfg.Secret_key, "secret_key", "", "pimento secret key")
	rootCmd.PersistentFlags().String("secret_key", "", "pimento secret key")

	viper.BindPFlag("pimento_secret_key", rootCmd.PersistentFlags().Lookup("secret_key"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// fmt.Println(SecretKey)
	// fmt.Println(cMgr().Secret_key)
	// fmt.Println(viper.GetString("pimento_secret_key"))
}
