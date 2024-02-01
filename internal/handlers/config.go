package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/limitcool/palworld-admin/config"
	"github.com/limitcool/palworld-admin/global"
	sp "github.com/limitcool/palworld-admin/settings-parse"
	"github.com/limitcool/palworld-admin/util"
	"gopkg.in/ini.v1"

	"github.com/gin-gonic/gin"
	"github.com/limitcool/starter/pkg/code"
)

const PalGameWorldSettingsName = "PalWorldSettings.ini"

func GetConfig(c *gin.Context) {
	path := filepath.Join(global.Config.PalSavedPath, "Config", util.GetPath(), PalGameWorldSettingsName)
	cfg, err := ini.Load(path)
	// fmt.Printf("cfg.SectionStrings(): %v\n", cfg.SectionStrings())
	// fmt.Println("App Mode:", cfg.Section("/Script/Pal.PalGameWorldSettings").Key("OptionSettings").String())
	if err != nil {
		code.AutoResponse(c, nil, code.NewErrCodeMsg(10001, "读取配置文件失败:"+err.Error()))
		c.Abort()
		return
	}

	d := sp.String2Struct(cfg.Section("/Script/Pal.PalGameWorldSettings").Key("OptionSettings").String())
	fmt.Printf("s: %v\n", d)
	code.AutoResponse(c, d, nil)
}

func UpdateConfig(c *gin.Context) {
	var req sp.ServerConfig
	err := c.ShouldBindJSON(&req)
	if err != nil {
		code.AutoResponse(c, nil, err)
		c.Abort()
		return
	}
	var pws PalGameWorldSettings
	pws.Settings.OptionSettings = sp.Struct2String(&req)
	cfg := ini.Empty()               //初始化一个空配置文件
	err = ini.ReflectFrom(cfg, &pws) //核心代码
	if err != nil {
		code.AutoResponse(c, nil, code.NewErrCodeMsg(10002, "反射配置文件失败:"+err.Error()))
		c.Abort()
		return
	}
	path := filepath.Join(global.Config.PalSavedPath, "Config", util.GetPath(), PalGameWorldSettingsName)
	cfg.SaveTo(path)
	if global.Config.RestartCommand != "" {
		err = restartServer(global.Config)
	}
	// fmt.Printf("cfg.Sections(): %v\n", cfg.SectionStrings())
	code.AutoResponse(c, nil, err)
}

type PalGameWorldSettings struct {
	Settings struct {
		OptionSettings string `ini:"OptionSettings"`
	} `ini:"/Script/Pal.PalGameWorldSettings"`
}

func restartServer(config config.Config) error {
	restartCmd := config.RestartCommand

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", restartCmd)
	case "linux", "darwin":
		cmd = exec.Command("sh", "-c", restartCmd)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
