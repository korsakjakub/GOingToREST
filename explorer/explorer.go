package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/korsakjakub/GOingToREST/config"
)

var ctx = context.Background()
var rdb *redis.Client
var conf config.Config

func main() {
	conf = config.LoadConfig([]string{"../config"})
	err := connectRedis()
	if err != nil {
		panic("error connecting to Redis:" + err.Error())
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/"+conf.Explorer.Function, getSize).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+conf.Explorer.Port, router))
}

func getSize(w http.ResponseWriter, _ *http.Request) {
	val, err := rdb.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}
	_, err2 := fmt.Fprintf(w, strconv.FormatInt(val, 10))
	if err2 != nil {
		return
	}
	log.Println(val)
}

func connectRedis() error {
	DBNumber, err := strconv.Atoi(conf.Redis.DB)
	if err != nil {
		panic("Could not load config correctly: " + err.Error())
		return err
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Address, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       DBNumber,
	})
	return nil
}
