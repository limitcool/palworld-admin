package save

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/limitcool/palworld-admin/config"
	"github.com/mholt/archiver/v4"
)

const SaveGamesPath = "SaveGames"

func BackupRoutine(config config.Config) {
	log.Debug("Start Backup")
	ticker := time.NewTicker(time.Duration(config.SaveConfig.BackupInterval) * time.Second)
	defer ticker.Stop()
	for range ticker.C {

		err := performBackup(config)
		if err != nil {
			log.Errorf("Backup failed: %v\n", err)
		} else {
			log.Info("Backup successful!")
		}

	}
}

func performBackup(config config.Config) error {
	// Create a filename based on the current timestamp
	backupFileName := time.Now().Format("2006-01-02-150405.zip")
	backupFilePath := filepath.Join(config.SaveConfig.BackupDirectory, backupFileName)

	return CompressZip(filepath.Join(config.PalSavedPath, SaveGamesPath), backupFilePath)
}
func cleanupRoutine(config config.Config) {
	for {
		err := performCleanup(config)
		if err != nil {
			fmt.Printf("Cleanup failed: %v\n", err)
		}
		time.Sleep(24 * time.Hour) // Sleep for a day before each cleanup cycle
	}
}

func performCleanup(config config.Config) error {
	cutoffTime := time.Now().Add(-time.Duration(config.SaveConfig.MaxRetentionDays) * 24 * time.Hour)

	err := filepath.Walk("/path/to/backup/directory", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file is a directory or if it's modified before the cutoff time
		if !info.IsDir() && info.ModTime().Before(cutoffTime) {
			fmt.Printf("Deleting old backup: %s\n", path)
			return os.Remove(path)
		}
		return nil
	})

	return err
}

func CompressZip(src string, dstPath string) error {
	files, err := archiver.FilesFromDisk(nil, map[string]string{
		// src: "", // contents added recursively
		src: "",
	})
	if err != nil {
		return err
	}
	// create the output file we'll write to
	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()
	format := archiver.CompressedArchive{
		Archival: archiver.Zip{},
	}
	err = format.Archive(context.Background(), out, files)
	if err != nil {
		return err
	}
	return nil
}
