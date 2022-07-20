package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/korsakjakub/GOingToREST/config"
	usr "github.com/korsakjakub/GOingToREST/user"
	"github.com/streadway/amqp"
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
	connection, channel, err := connectRabbitMQ()
	msgs, err := channel.Consume(conf.RabbitMQ.QueueName, "", false, false, false, false, nil)
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

func connectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
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
	_, err = channel.QueueDeclare(conf.RabbitMQ.QueueName, true, false, false, false, nil)
	if err != nil {
		panic("error declaring the queue: " + err.Error())
		return nil, nil, err
	}
	err = channel.QueueBind(conf.RabbitMQ.QueueName, "#", conf.RabbitMQ.ExchangeName, false, nil)
	if err != nil {
		panic("error binding to the queue: " + err.Error())
		return nil, nil, err
	}
	return connection, channel, err
}
