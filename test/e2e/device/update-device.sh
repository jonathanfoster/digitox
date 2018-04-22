#!/bin/bash
set -e

curl -i -X PUT -d \
'{
  "name": "jonathan-laptop",
  "password": "Digitox321"
}' \
http://localhost:8080/devices/jonathan-laptop
