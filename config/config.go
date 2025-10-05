package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// user -> postgres
// password -> redass
// host -> localhost
// port -> 5432
// db name -> gocom
type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

var configurations *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	SecretKey   string
	DB          *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	http_port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse the http port")
		os.Exit(1)
	}
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		fmt.Println("Secret key is required")
		os.Exit(1)
	}

	db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		fmt.Println("Db host is required")
		os.Exit(1)
	}

	db_port := os.Getenv("DB_PORT")
	if db_port == "" {
		fmt.Println("Db port is required")
		os.Exit(1)
	}
	db_prt, err := strconv.ParseInt(db_port, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse the db port")
		os.Exit(1)
	}

	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		fmt.Println("Db name is required")
		os.Exit(1)
	}
	db_user := os.Getenv("DB_USER")
	if db_user == "" {
		fmt.Println("Db user is required")
		os.Exit(1)
	}
	db_password := os.Getenv("DB_PASSWORD")
	if db_password == "" {
		fmt.Println("Db password is required")
		os.Exit(1)
	}
	db_ssl_mode := os.Getenv("ENABLE_SSL_MODE")
	mode, err := strconv.ParseBool(db_ssl_mode)
	if err != nil {
		fmt.Println("Failed to parse the db ssl mode value", err)
		os.Exit(1)
	}

	db_config := &DBConfig{
		Host:          db_host,
		Port:          int(db_prt),
		Name:          db_name,
		User:          db_user,
		Password:      db_password,
		EnableSSLMode: mode,
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(http_port), // type casting -- two types big to sm / sm to big...
		SecretKey:   secret,
		DB:          db_config,
	}
}

func GetConfigs() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
