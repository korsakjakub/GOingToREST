package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	redis "github.com/go-redis/redis/v9"
	mux "github.com/gorilla/mux"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

func getSize(w http.ResponseWriter, r *http.Request) {

	val, err := rdb.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, strconv.FormatInt(val, 10))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/size", getSize).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
