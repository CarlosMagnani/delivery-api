package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", 4000)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "process")
    viper.SetDefault("database.password", "processdb")
    viper.SetDefault("database.database", "delivery")
	err := LoadEnvVars()
	if err != nil {
		fmt.Println(err)
		return
	}
	
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	
	err := LoadEnvVars()
	if err != nil {
		return err
	}

	err = viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(err)
			return err
		}
	}

	

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB =  DBConfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}
	return nil
}


func GetDatabase() DBConfig{
	return cfg.DB
}

func GetServer() APIConfig{
	return cfg.API
}