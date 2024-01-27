package global

import (
	"path/filepath"

	"github.com/limitcool/palworld-admin/config"
	"github.com/limitcool/palworld-admin/util"
)

var Config config.Config
var AppDir = util.AppDataDir("palworld-admin", false)

// var ConfigPath = filepath.Join(AppDir, "config")
var ConfigFile = filepath.Join(AppDir, "config.yaml")

const VERSION = "0.1.3"
