#!/bin/bash
id="$(($RANDOM % 666))"
age=$"$((10 + $RANDOM % 50))"
name="$(echo $RANDOM | md5sum | head -c 10; echo;)"
surname="$(echo $RANDOM | md5sum | head -c 10; echo;)"
jsonfmt='{"id": "%s", "name":"%s", "surname":"%s","age":"%s"}'
#json=$(printf "$jsonfmt" "$id" "$name" "$surname" "$age")
json=$(jq -n --arg ii "$id" --arg nm "$name" --arg sr "$surname" --arg ag "$age" '{id: $ii, name: $nm, surname: $sr, age: $ag}')
curl -X POST -H "Content-Type: application-json" -d "$json" 20.105.72.86/add
#curl -X POST -H "Content-Type: application-json" -d "$json" localhost:6666/add
