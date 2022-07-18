package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	redis "github.com/go-redis/redis/v9"
	usr "github.com/korsakjakub/GOingToREST/user"
	amqp "github.com/streadway/amqp"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

func sendToRedis(msg []byte) error {
	var user usr.User
	err := json.Unmarshal(msg, &user)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, user.ID, msg, 60*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	url := "amqp://guest:guest@rabbitmq:5672"
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
	_, err = channel.QueueDeclare("test", true, false, false, false, nil)
	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}
	err = channel.QueueBind("test", "#", "events", false, nil)
	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}
	msgs, err := channel.Consume("test", "", false, false, false, false, nil)
	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}
	for msg := range msgs {
		log.Println("Received a message from Rabbitmq: " + string(msg.Body))
		msg.Ack(false)
		err = sendToRedis(msg.Body)
		if err != nil {
			panic(err.Error())
		}
	}
	defer connection.Close()
}
