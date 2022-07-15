# GOingToREST

GOintToREST is a simple 3-app system built for me to learn about communicating between different services.

* [`poster/post_producer.go`](poster/post_producer.go) 
  * Listens to POST requests @ port 6666
  * Redirects data from POSTs to Rabbitmq
* [`saver/saver.go`](saver/saver.go)
  * Consumes Rabbitmq queue
  * Saves incoming data to Redis
* [`explorer/explorer.go`](explorer/explorer.go)
  * Listens to GET requests @ localhost:8000/size
  * Returns the count of all keys in Redis
* [`user/user.go`](user/user.go)
  * Struct User with some sample fields like id, name, surname, and age.
  
For debugging/testing purposes, in [`scripts/`](scripts/) there is [`post_curl.sh`](scripts/post_curl.sh) which send random POSTs of correct form

## Requirements
* Redis (@ port 6379)
* Rabbitmq (@ port 5672)
