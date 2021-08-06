package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fuzziekus/pimento/crypto"
	"github.com/fuzziekus/pimento/ui/util"
	"github.com/go-yaml/yaml"
	"github.com/spf13/viper"
)

type config struct {
	Db struct {
		Path string
	}
	Secret_key string
}

var cfg config
var CfgFile string
var DbCryptor *crypto.Cryptor
var RowCryptor *crypto.Cryptor

// var SecretKey string

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
		config_file_path := getConfigPath()
		createDir(config_file_path)

		// Search config.yaml
		viper.AddConfigPath(config_file_path)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	viper.BindEnv("secret_key", "PIMENTO_SECRET_KEY")
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

	var db_key string
	for len(db_key) != 16 && len(db_key) != 32 && len(db_key) != 64 {
		db_key = util.InputSecretString("DB Key (16 or 32 or 64 byte)")
	}

	cfg.Secret_key = getEnv(viper.GetString("pimento_secret_key"), viper.GetString("secret_key"))
	if cfg.Secret_key == "" {
		for len(cfg.Secret_key) != 16 && len(cfg.Secret_key) != 32 && len(cfg.Secret_key) != 64 {
			cfg.Secret_key = util.InputSecretString("Secret Key (16 or 32 or 64 byte)")
		}
	}

	RowCryptor = crypto.NewCryptor(cfg.Secret_key)
	DbCryptor = crypto.NewCryptor(db_key)
	if existFile(cfg.Db.Path) {
		if err := DbCryptor.DecryptFile(cfg.Db.Path, cfg.Db.Path+".temp"); err != nil {
			log.Fatal(err)
		}
	}
}

func getConfigPath() string {
	xdg_config_home := getEnv("XDG_CONFIG_HOME", os.Getenv("HOME"))
	return filepath.Join(xdg_config_home, "pimento")
}

func existFile(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func WriteConfig(exeistCheck bool) error {
	configFileName := "config.yaml"
	buf, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}
	configFilePath := filepath.Join(getConfigPath(), configFileName)
	if existFile(configFilePath) {
		if exeistCheck && util.ChoiceYN("既存のコンフィグファイルを上書きしますか?") {
			return os.WriteFile(configFilePath, buf, 0644)
		}
		return nil
	}

	return os.WriteFile(configFilePath, buf, 0644)
}
