package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

// User ...
type User struct {
	ID      string `json:"Id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
}

// Users
var Users []User

func add(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	err := json.Unmarshal(reqBody, &user)

	if err != nil {
		panic("400: Bad Request")
	}

	Users = append(Users, user)

	json.NewEncoder(w).Encode(user)

	fmt.Println("User: ", user)

	url := os.Getenv("AMQP_URL")
	//If it doesnt exist, use the default connection string
	if url == "" {
		url = "amqp://guest:guest@localhost:5672"
	}
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        b,
	}

	err = channel.Publish("events", "random-key", false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}

	_, err = channel.QueueDeclare("addingUser", true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("addingUser", "#", "events", false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/add", add).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	fmt.Println("Listening... d-_-b")
	handleRequests()
}
