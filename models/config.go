package models

type ConfigDB struct {
	Host     string
	Port     string
	Name     string
	Username string
	SSLMode  string
	Password string
}

type ConfigApp struct {
	Port string
}

type Config struct {
	DB  ConfigDB
	App ConfigApp
}
