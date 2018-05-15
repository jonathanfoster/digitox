#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "blocklist id not provided" >&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "blocklist name not provided" >&2
    exit 1
fi

if [ -z "$3" ]
then
    echo "blocklist domain not provided" >&2
    exit 1
fi

curl -s -X PUT -d \
"{
  \"name\": \"${2}\",
  \"domains\": [\"${3}\"]
}" \
"http://localhost:8080/blocklists/${1}?access_token=$DIGITOX_ACCESS_TOKEN"

