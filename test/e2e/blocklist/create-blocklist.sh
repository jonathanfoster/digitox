#!/bin/bash
set -e

curl -s -X POST -d \
"{
  \"name\": \"$(uuidgen)\",
  \"domains\": [\"www.reddit.com\"]
}" \
"http://localhost:8080/blocklists/?access_token=$DIGITOX_ACCESS_TOKEN"
