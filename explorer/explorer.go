package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	redis "github.com/go-redis/redis/v9"
	mux "github.com/gorilla/mux"
	config "github.com/korsakjakub/GOingToREST/config"
)

var ctx = context.Background()
var rdb *redis.Client
var conf config.Config

func getSize(w http.ResponseWriter, r *http.Request) {

	val, err := rdb.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, strconv.FormatInt(val, 10))
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

func main() {
	conf, _ = config.LoadConfig([]string{"../config"})
	err := connectRedis()
	if err != nil {
		panic("error connecting to Redis:" + err.Error())
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/"+conf.Explorer.Function, getSize).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+conf.Explorer.Port, router))
}
