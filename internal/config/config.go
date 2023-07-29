package config

type AppEnvironment string

type Database struct {
	Hostname string
	Port     string
	Database string
	Username string
	Password string
	DebugLog bool
}

type HTTPServer struct {
	Debug    bool
	Hostname string
	Port     string
}

type Config struct {
	Database    Database
	HTTPServer  HTTPServer
	IsLocal     bool
	Environment string
}
