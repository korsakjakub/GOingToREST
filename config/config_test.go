package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	want := Config{
		Redis: RedisConfig{
			Address:  "test",
			Port:     "0000",
			Password: "psss",
			DB:       "dbb",
		},
		RabbitMQ: RabbitMQConfig{
			Login:        "login",
			Password:     "pssdasl",
			Address:      "alskjd",
			Port:         "ilaksdj",
			ExchangeName: "alsdkj",
			ExchangeType: "djskal",
			QueueName:    "djskal",
		},
		Poster: PosterConfig{
			Port:     "port",
			Function: "asldkj",
		},
		Explorer: ExplorerConfig{
			Port:     "sdkj",
			Function: "fff",
		},
	}

	configJson, err := json.Marshal(want)
	if err != nil {
		t.Error("Could not marshal User struct")
	}
	_ = ioutil.WriteFile("test.json", configJson, 0644)
	got := LoadConfig([]string{"."}, "test", "json")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %s, got %s", want, got)
	}
	err = os.Remove("test.json")
	if err != nil {
		return
	}
}
