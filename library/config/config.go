package config

import (
	"os"
	"pkt_statistics/library/logger"

	"github.com/spf13/viper"
)

// config 配置
var config *Config

// Config 配置文件结构
type Config struct {
	Server  ServerConfig  `json:"server"`
	Storage StorageConfig `json:"storage"`
	Nats    NatsConfig    `json:"nats"`
}

func init() {
	config = new(Config)

	vpr := viper.New()
	if os.Getenv("GIN_MODE") == "release" {
		vpr.SetConfigName(ConfigProFileName.String())
	} else {
		vpr.SetConfigName(ConfigDevFileName.String())
	}
	vpr.SetConfigType(ConfigFileType.String())
	vpr.AddConfigPath(ConfigFilePath.String())
	if err := vpr.ReadInConfig(); err != nil {
		logger.FatalF("ReadInConfig", err)
	}

	if err := vpr.Unmarshal(&config); err != nil {
		logger.FatalF("Unmarshal Config", err)
	}
}

// GetServerConf 获取服务器配置
func GetServerConf() ServerConfig {
	return config.Server
}

// GetStorageConf 获取存储配置
func GetStorageConf() StorageConfig {
	return config.Storage
}

// GetNatsConf 获取Nats配置
func GetNatsConf() NatsConfig {
	return config.Nats
}
