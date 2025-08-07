package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	ApiConfig *ApiConfig
	DBConfig  *DBConfig
}

type ApiConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SslMode  string
	Url      string
}

var Cfg *config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func LoadConfig() *config {
	if Cfg != nil {
		fmt.Println("Error loading config: Config already loaded")
		return Cfg
	}

	setDefaultValues()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, using default values:", err)
	}

	dbConfig := &DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     viper.GetString("database.name"),
		SslMode:  viper.GetString("database.sslmode"),
	}

	dbConfig.Url = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.SslMode,
	)

	Cfg = &config{
		ApiConfig: &ApiConfig{
			Port: viper.GetString("api.port"),
		},
		DBConfig: dbConfig,
	}

	return Cfg
}

func setDefaultValues() {
	// API Config
	viper.SetDefault("api.port", 3333)

	// DB Config
	viper.SetDefault("database.host", "postgres")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.name", "youtofy")
	viper.SetDefault("database.sslmode", "disable")
}
