package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	config "github.com/korsakjakub/GOingToREST/config"
	usr "github.com/korsakjakub/GOingToREST/user"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

var users []usr.User
var conf config.Config

func main() {
	conf, _ = config.LoadConfig([]string{"../config"})
	fmt.Println("Listening for POSTs... d-_-b")
	handleRequests()
}

func connectToRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", conf.RabbitMQ.Login, conf.RabbitMQ.Password, conf.RabbitMQ.Address, conf.RabbitMQ.Port)
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
		return nil, nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
		return nil, nil, err
	}
	err = channel.ExchangeDeclare(conf.RabbitMQ.ExchangeName, conf.RabbitMQ.ExchangeType, true, false, false, false, nil)
	if err != nil {
		panic(err)
		return nil, nil, err
	}
	return connection, channel, err
}

func add(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user usr.User
	err := json.Unmarshal(reqBody, &user)

	if err != nil || reflect.DeepEqual(user, usr.User{}) {
		w.WriteHeader(400)
		w.Write([]byte("400 Bad Request"))
		return
	}
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
	log.Println("User: ", user)
	user_data_json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        user_data_json,
	}
	connection, channel, err := connectToRabbitMQ()
	err = channel.Publish(conf.RabbitMQ.ExchangeName, "random-key", false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}
	defer connection.Close()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/"+conf.Poster.Function, add).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+conf.Poster.Port, router))
}
