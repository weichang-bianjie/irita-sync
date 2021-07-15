package config

import (
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/bianjieai/irita-sync/utils/constant"
	"github.com/spf13/viper"
	"os"
)

var (
	ConfigFilePath string
)

type (
	Config struct {
		DataBase DataBaseConf `mapstructure:"database"`
		Server   ServerConf   `mapstructure:"server"`
	}
	DataBaseConf struct {
		Addrs    string `mapstructure:"addrs"`
		User     string `mapstructure:"user"`
		Passwd   string `mapstructure:"passwd" json:"-"`
		Database string `mapstructure:"database"`
	}

	ServerConf struct {
		ChainId               string `mapstructure:"chain_id"`
		WriteDir              string `mapstructure:"write_dir"`
		FilePrefix            string `mapstructure:"file_prefix"`
		PromethousPort        string `mapstructure:"promethous_port"`
		OnlySupportModule     string `mapstructure:"only_support_module"`
		ChainBlockResetHeight int64  `mapstructure:"chain_block_reset_height"`
	}
)

func init() {
	websit, found := os.LookupEnv(constant.EnvNameConfigFilePath)
	if found {
		ConfigFilePath = websit
	} else {
		panic("not found CONFIG_FILE_PATH")
	}
}

func ReadConfig() (*Config, error) {

	rootViper := viper.New()
	// Find home directory.
	rootViper.SetConfigFile(ConfigFilePath)

	// Find and read the config file
	if err := rootViper.ReadInConfig(); err != nil { // Handle errors reading the config file
		return nil, err
	}

	var cfg Config
	if err := rootViper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	logger.Debug("config: " + utils.MarshalJsonIgnoreErr(cfg))

	return &cfg, nil
}
