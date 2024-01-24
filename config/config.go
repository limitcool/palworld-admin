package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PalWorldConfigFilePath string
	AdminPassword          string
}

// 初始化并生成默认配置
func InitDefaultConfig() {

	defaultConfig := Config{
		PalWorldConfigFilePath: "/root/palworld/data/Config/LinuxServer/PalWorldSettings.ini",
		AdminPassword:          "initcool-https://blog.nmslwsnd.com",
	}
	viper.SetConfigType("yaml")
	// 将默认配置写入配置文件
	viper.Set("PalWorldConfigFilePath", defaultConfig.PalWorldConfigFilePath)
	viper.Set("AdminPassword", defaultConfig.AdminPassword)
	err := viper.WriteConfig()
	if err != nil {
		// 处理错误
		fmt.Println("Error writing default config:", err)
	}
}
