package config

import (
	"os"

	"github.com/charmbracelet/log"

	"github.com/spf13/viper"
)

type Config struct {
	PalWorldConfigFilePath string
	AdminPassword          string
	Port                   int
}

// 初始化并生成默认配置
func InitDefaultConfig(configPath string, configFile string) {
	defaultConfig := Config{
		PalWorldConfigFilePath: "/root/palworld/data/Config/LinuxServer/PalWorldSettings.ini",
		AdminPassword:          "initcool-https://blog.nmslwsnd.com",
		Port:                   8080,
	}
	if err := os.MkdirAll(configPath, 0755); err != nil {
		log.Error(err)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	// 将默认配置写入配置文件
	viper.Set("PalWorldConfigFilePath", defaultConfig.PalWorldConfigFilePath)
	viper.Set("AdminPassword", defaultConfig.AdminPassword)
	viper.Set("Port", 8080)
	err := viper.WriteConfig()
	if err != nil {
		// 处理错误
		log.Error("Error writing default config:", err)
	}
}
