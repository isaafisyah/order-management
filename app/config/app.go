package config

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Env       string
	Name      string
	Host      string
	Port      string
	SecretKey string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}