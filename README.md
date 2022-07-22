# GOingToREST

GOingToREST is a simple 3-app system built for me to learn about communicating between different services.

* [`poster/poster.go`](poster/poster.go) 
  * Listens to POST requests
  * Redirects data from POSTs to Rabbitmq
* [`saver/saver.go`](saver/saver.go)
  * Consumes Rabbitmq queue
  * Saves incoming data to Redis
* [`explorer/explorer.go`](explorer/explorer.go)
  * Listens to GET requests @ `/size`
  * Returns the count of all keys in Redis
* [`user/user.go`](user/user.go)
  * Struct User with some sample fields like id, name, surname, and age.
  
For debugging/testing purposes, in [`scripts/`](scripts/) there is [`post_curl.sh`](scripts/post_curl.sh) which send random POSTs of correct form
(remark: [jq](https://stedolan.github.io/jq/) required)

## Requirements
* Docker bridge
```
$ docker network create <network-name>
```

* Redis (@ port 6379 & name=redis by default, can be changed in [`config.yaml`](config/config.yaml))
* Rabbitmq (@ port 5672 & name=rabbitmq by default, can be changed in [`config.yaml`](config/config.yaml))

# Usage

In order to build docker images for apps in [`poster/`](poster/), [`saver/`](saver/), and [`explorer/`](explorer/) run

```sh
$ # Let <name> = poster, etc.
$ docker build -t <name> .
```

To run:

```sh
$ docker run --rm -it --name <name> --network <network-name> -v <path to config file>:<$CONFIG_PATH><$CONFIG_NAME>.<$CONFIG_TYPE> -p <ports> <name>
```
