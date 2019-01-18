package config

import "os"

const (
	PORT = 8080
	IPADDR = "127.0.0.1"
	ERRPATH =  "{\"error\":\"Unsupported Path\"}"
)

var DefaultPath,_ = os.Getwd()

func init() {
	DefaultPath +="/storage"
}

