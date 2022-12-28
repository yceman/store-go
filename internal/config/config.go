package config

import (
	"os"
	"strconv"
)

const (
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologation"
	PRODUCTION   = "production"
)

type Config struct {
	SRV_PORT    string `json:"srv_port"`
	WEB_UI      bool   `json:"web_ui"`
	OpenBrowser bool   `json:"open_browser"`
	Mode        string `json:"mode"`

	DBConfig `json:"db_config"`
}

type DBConfig struct {
	DB_DRIVE string `json:"db_drive"`
	DB_HOST  string `json:"db_host"`
	DB_PORT  string `json:"db_port"`
	DB_USER  string `json:"db_user"`
	DB_PASS  string `json:"db_pass"`
	DB_NAME  string `json:"db_name"`
	DB_DSN   string `json:"-"`
}

func NewConfig(config *Config) *Config {
	var conf *Config
	if config == nil || config.SRV_PORT == "" {
		conf = defaultConfig()
	} else {
		conf = config
	}

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}
	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	SRV_OPEN_BROWSER := os.Getenv("SRV_OPEN_BROWSER")
	if SRV_OPEN_BROWSER != "" {
		conf.OpenBrowser, _ = strconv.ParseBool(SRV_OPEN_BROWSER)
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	// Config DB

	SRV_DB_DRIVE := os.Getenv("SRV_DB_DRIVE")
	if SRV_DB_DRIVE != "" {
		conf.DB_DRIVE = SRV_DB_DRIVE
	}

	SRV_DB_HOST := os.Getenv("SRV_DB_HOST")
	if SRV_DB_HOST != "" {
		conf.DB_HOST = SRV_DB_HOST
	}

	SRV_DB_PORT := os.Getenv("SRV_DB_PORT")
	if SRV_DB_PORT != "" {
		conf.DB_PORT = SRV_DB_PORT
	}

	SRV_DB_USER := os.Getenv("SRV_DB_USER")
	if SRV_DB_USER != "" {
		conf.DB_USER = SRV_DB_USER
	}

	SRV_DB_PASS := os.Getenv("SRV_DB_PASS")
	if SRV_DB_PASS != "" {
		conf.DB_PASS = SRV_DB_PASS
	}

	SRV_DB_NAME := os.Getenv("SRV_DB_NAME")
	if SRV_DB_NAME != "" {
		conf.DB_NAME = SRV_DB_NAME
	}

	return conf
}

func defaultConfig() *Config {
	default_config := Config{
		SRV_PORT:    "8080",
		WEB_UI:      true,
		OpenBrowser: true,
		Mode:        PRODUCTION,
		DBConfig: DBConfig{
			DB_DRIVE: "sqlite3",
			DB_NAME:  "db.sqlite3",
		},
	}
	return &default_config
}
