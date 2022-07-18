package config

import (
    "testing"
	"log"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig([]string{".."})
	if err != nil {
		panic(err.Error())
	}
	log.Println(config)
}