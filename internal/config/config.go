package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BotToken string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		BotToken: viper.GetString("bot_token"),
	}

	return config, nil
}
