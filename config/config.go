package config

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Address string `mapstructure:"address"`
	Port 	string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	db string `mapstructure:"db"`
}

type RabbitMQConfig struct {
	Login string `mapstructure:"login"`
	Password string `mapstructure:"password"`
	Address string `mapstructure:"address"`
	Port string `mapstructure:"port"`
}

type Config struct {
	Redis RedisConfig `mapstructure:"redis"`
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
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