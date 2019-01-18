package config

import (
	"os"
	"path/filepath"
)

const (
	PORT = 8080
	IPADDR = "127.0.0.1"
	ERRPATH =  "{\"error\":\"Unsupported Path\"}"
)

var DefaultPath = os.Getenv("GOPATH")
var DefaultPathStorage = filepath.Join(DefaultPath,"/storage")
var DefaultPathPythonScript =  filepath.Join(DefaultPath,"/pythonScriptStatsFiles/script.py")

