package config

import "github.com/spf13/viper"

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	StringConn string
}

var cfg *config

func Load() error {
	viper.SetConfigName("cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../config/")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		StringConn: viper.GetString("db.stringConn"),
	}

	return nil
}

func DB() DBConfig {
	return cfg.DB
}

func ServerPort() string {
	return cfg.API.Port
}
