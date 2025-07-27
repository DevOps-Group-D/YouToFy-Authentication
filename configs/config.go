package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	ApiConfig *ApiConfig
}

type ApiConfig struct {
	Port string
}

func NewConfig() *config {
	viper.SetDefault("api.port", 3333)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, using default values:", err)
	}

	cfg := &config{
		ApiConfig: &ApiConfig{
			Port: viper.GetString("api.port"),
		},
	}

	return cfg
}
