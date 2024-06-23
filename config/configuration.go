package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	App         appConfiguration
	DB          databaseConfiguration
	ExternalApi externalApiConfiguration
}

type appConfiguration struct {
	ENV     string
	DbDebug bool
	Port    string
	Address string
}

type databaseConfiguration struct {
	Driver   string
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

type externalApiConfiguration struct {
	Bca     BankAPIConfiguration
	Mandiri BankAPIConfiguration
}

type BankAPIConfiguration struct {
	BaseURL string
	ApiKey  string
}

var (
	configuration *Configuration
	once          sync.Once
)

func Load(configPath string) *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configPath)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("unable to read config, %v", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	})

	return configuration
}
