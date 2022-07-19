package config

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type RabbitMQConfig struct {
	Login        string `mapstructure:"login"`
	Password     string `mapstructure:"password"`
	Address      string `mapstructure:"address"`
	Port         string `mapstructure:"port"`
	ExchangeName string `mapstructure:"exchangename"`
	ExchangeType string `mapstructure:"exchangetype"`
	QueueName    string `mapstructure:"queuename"`
}

type PosterConfig struct {
	Port     string `mapstructure:"port"`
	Function string `mapstructure:"function"`
}

type ExplorerConfig struct {
	Port     string `mapstructure:"port"`
	Function string `mapstructure:"function"`
}

type Config struct {
	Redis    RedisConfig    `mapstructure:"redis"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
	Poster   PosterConfig   `mapstructure:"poster"`
	Explorer ExplorerConfig `mapstructure:"explorer"`
}

var vp *viper.Viper

func LoadConfig(additionalPath []string) (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")
	for _, path := range additionalPath {
		vp.AddConfigPath(path)
	}
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
