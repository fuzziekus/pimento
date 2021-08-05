package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Db struct {
		Path string
	}
}

var cfg config
var CfgFile string

func Mgr() config { return cfg }

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func createDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find XDG_CONFIG_HOME directory if XDG_CONFIG_HOME set.
		// Find HOME directory if XDG_CONFIG_HOME not set
		xdg_config_home := getEnv("XDG_CONFIG_HOME", os.Getenv("HOME"))
		config_file_path := filepath.Join(xdg_config_home, "pimento")
		createDir(config_file_path)

		// Search config.yaml
		viper.AddConfigPath(config_file_path)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()

	// 設定ファイルの内容を構造体にコピーする
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Db.Path == "" {
		xdg_data_home := getEnv("XDG_DATA_HOME", os.Getenv("HOME"))
		db_file_path := filepath.Join(xdg_data_home, "pimento", "credential.db")
		createDir(filepath.Dir(db_file_path))
		cfg.Db.Path = db_file_path
	}
}
