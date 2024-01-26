package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenOrCreateFile(configPath, configFileName string) (*os.File, error) {
	// 确保目录存在
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return nil, fmt.Errorf("无法创建目录: %v", err)
	}

	// 构建完整的文件路径
	fullPath := filepath.Join(configPath, configFileName)

	// 打开或创建文件
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("无法打开或创建文件: %v", err)
	}

	return file, nil
}
