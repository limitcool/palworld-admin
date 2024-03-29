package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/limitcool/lib"
	"github.com/limitcool/palworld-admin/config"
	"github.com/limitcool/palworld-admin/global"
	"github.com/limitcool/palworld-admin/routers"
	"github.com/limitcool/palworld-admin/save"
	"github.com/limitcool/palworld-admin/util"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func main() {
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
		fmt.Println("Error unmarshalling config:", err)
	}

	// 打印配置
	log.Infof("PalSavedPath: %v\n", global.Config.PalSavedPath)
	log.Infof("AdminPassword: %s\n", global.Config.AdminPassword)
	log.Infof("Port: %d\n", global.Config.Port)
	util.EnsureDirectoryExists(global.Config.SaveConfig.BackupDirectory)
	go save.BackupRoutine(global.Config)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprint("0.0.0.0:", global.Config.Port),
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	log.Infof("Listen: %s:%d\n", "http://127.0.0.1", global.Config.Port)
	go func() {
		// 服务连接 监听
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("Listen:%s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器,这里需要缓冲
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	//(设置5秒超时时间)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}
