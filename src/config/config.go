//Package config contain all the parameters and configurations of the project ET

package config

import (
	"os"
	"path/filepath"
)

const (
	PORT = 5000
	IPADDR = "0.0.0.0"
	ERRPATH =  "{\"error\":\"Unsupported Path\"}"
)

var DefaultPath,_ = os.Getwd()
var DefaultPathStorage = filepath.Join(DefaultPath,"/storage")
var DefaultPathPythonScript =  filepath.Join(DefaultPath,"/pythonScriptStatsFiles/script.py")

