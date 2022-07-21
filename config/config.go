package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
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

func LoadConfig(additionalPath []string, args ...string) Config {
	vp = viper.New()
	var config Config
	log.Println(os.Getenv("CONFIG_NAME"))
	log.Println(os.Getenv("CONFIG_TYPE"))
	if len(os.Getenv("CONFIG_NAME")) > 0 && len(os.Getenv("CONFIG_TYPE")) > 0 {
		vp.SetConfigName(os.Getenv("CONFIG_NAME"))
		vp.SetConfigType(os.Getenv("CONFIG_TYPE"))
	} else if len(args) > 0 {
		vp.SetConfigName(args[0])
		vp.SetConfigType(args[1])
	} else {
		vp.SetConfigName("config")
		vp.SetConfigType("yaml")
	}
	vp.AddConfigPath("/")
	vp.AddConfigPath("./config")
	vp.AddConfigPath(".")
	for _, path := range additionalPath {
		vp.AddConfigPath(path)
	}

	err := vp.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read the config file: ", err.Error())
		return Config{}
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Fatal("Cannot unmarshal the config file: ", err.Error())
		return Config{}
	}

	return config
}
