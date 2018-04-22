#!/bin/bash
set -e

curl -i -X POST -d \
'{
  "name": "jonathan-laptop",
  "password": "Digitox123"
}' \
http://localhost:8080/devices/
