package config

import (
	"os"

	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type SaveConfig struct {
	BackupInterval   int    `yaml:"BackupInterval"`   // Interval between backups in minutes
	MaxRetentionDays int    `yaml:"MaxRetentionDays"` // Maximum retention period for backups in days
	BackupDirectory  string `yaml:"BackupDirectory"`  // Directory to retain game saves
}
type Config struct {
	PalSavedPath  string     `yaml:"PalSavedPath"`
	AdminPassword string     `yaml:"AdminPassword"`
	Port          int        `yaml:"Port"`
	SaveConfig    SaveConfig `yaml:"SaveConfig"`
}

// 初始化并生成默认配置
func InitDefaultConfig(configPath string, configFile string) {
	defaultConfig := Config{
		PalSavedPath:  "",
		AdminPassword: "initcool-https://blog.nmslwsnd.com",
		Port:          8080,
		SaveConfig: SaveConfig{
			BackupInterval:   60,
			MaxRetentionDays: 7,
			BackupDirectory:  "backups/",
		},
	}

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
