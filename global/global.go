package global

import (
	"path/filepath"

	"github.com/decred/dcrd/dcrutil/v2"
	"github.com/limitcool/palworld-admin/config"
)

var Config config.Config
var AppDir = dcrutil.AppDataDir("palworld-Admin", false)
var ConfigPath = filepath.Join(AppDir, "config")
var ConfigFile = filepath.Join(ConfigPath, "config.yaml")

const VERSION = "0.1.1"
