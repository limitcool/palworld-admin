package bootstrap

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/limitcool/lib"
	"github.com/limitcool/palworld-admin/config"
	"github.com/limitcool/palworld-admin/global"
	"github.com/spf13/viper"
)

func init() {
	lib.SetDebugMode(func() {
		log.Info("Debug Mode")
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	})

	log.Infof("https://github.com/limitcool/palworld-admin version: %s ", global.VERSION)
	log.Infof("AppDir: %v\n", global.AppDir)
	viper.SetConfigFile(global.ConfigFile)
	err := viper.ReadInConfig()
	// 如果找不到配置文件
	if os.IsNotExist(err) {
		log.Info("Config file not found. Initializing with default values...")
		// 初始化并生成默认配置
		config.InitDefaultConfig(global.AppDir, global.ConfigFile)
		// 重新尝试读取配置文件
		err = viper.ReadInConfig()
		if err != nil {
			// 处理错误
			log.Error("Error reading config file:", err)
		}
	}
	// 解析配置到结构体
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		// 处理错误
		log.Error("Error unmarshalling config:", err)
	}
}
