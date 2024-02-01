package config

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type SaveConfig struct {
	BackupInterval   int     `yaml:"BackupInterval" default:"60"`        // Interval between backups in minutes
	MaxRetentionDays float64 `yaml:"MaxRetentionDays" default:"7"`       // Maximum retention period for backups in days
	BackupDirectory  string  `yaml:"BackupDirectory" default:"backups/"` // Directory to retain game saves
}
type Config struct {
	PalSavedPath   string     `yaml:"PalSavedPath" default:""`
	AdminPassword  string     `yaml:"AdminPassword" default:"initcool-https://blog.nmslwsnd.com"`
	Port           int        `yaml:"Port" default:"8080"`
	SaveConfig     SaveConfig `yaml:"SaveConfig"`
	RestartCommand string     `yaml:"RestartCommand" default:""` // 新添加的字段，表示用户用于重启的命令
}

// 初始化并生成默认配置
func InitDefaultConfig(configPath string, configFile string) {
	// defaultConfig := Config{
	// 	PalSavedPath:  "",
	// 	AdminPassword: "initcool-https://blog.nmslwsnd.com",
	// 	Port:          8080,
	// 	SaveConfig: SaveConfig{
	// 		BackupInterval:   60,
	// 		MaxRetentionDays: 7,
	// 		BackupDirectory:  "backups/",
	// 	},
	// }
	var defaultConfig Config
	err := defaults.Set(&defaultConfig)
	log.Debug(err)

	if err := os.MkdirAll(configPath, 0755); err != nil {
		log.Error(err)
	}
	data, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		log.Error(err)
	}
	err = os.WriteFile(configFile, data, 0755)
	if err != nil {
		log.Error(err)
	}
}
