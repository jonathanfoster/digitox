#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "blocklist name not provided" >&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "blocklist domain not provided" >&2
    exit 1
fi

curl -s -X POST -d \
"{
  \"name\": \"${1}\",
  \"domains\": [\"${2}\"]
}" \
"http://localhost:8080/blocklists/?access_token=$DIGITOX_ACCESS_TOKEN"
