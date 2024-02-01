package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/limitcool/palworld-admin/bootstrap"
	"github.com/limitcool/palworld-admin/global"
	"github.com/limitcool/palworld-admin/routers"
	"github.com/limitcool/palworld-admin/save"
	"github.com/limitcool/palworld-admin/util"

	"github.com/charmbracelet/log"
)

func main() {
	// 打印配置
	log.Infof("PalSavedPath: %v\n", global.Config.PalSavedPath)
	log.Infof("AdminPassword: %s\n", global.Config.AdminPassword)
	log.Infof("Port: %d\n", global.Config.Port)
	util.EnsureDirectoryExists(global.Config.SaveConfig.BackupDirectory)
	save.RunBackupAndCleanup(global.Config)
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
