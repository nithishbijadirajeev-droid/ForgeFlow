package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Host string
		Port int
	}

	App struct {
		Name    string
		Version string
	}
}

var AppConfig Config

func Load() {
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", 8080)

	viper.SetDefault("APP_NAME", "ForgeFlow")
	viper.SetDefault("APP_VERSION", "v0.1.0")

	viper.AutomaticEnv()

	AppConfig.Server.Host = viper.GetString("SERVER_HOST")
	AppConfig.Server.Port = viper.GetInt("SERVER_PORT")

	AppConfig.App.Name = viper.GetString("APP_NAME")
	AppConfig.App.Version = viper.GetString("APP_VERSION")
}