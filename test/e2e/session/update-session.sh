#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "session not provided" 1>&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "blocklist not provided" 1>&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "device not provided" 1>&2
    exit 1
fi

TODAY=$(date +%Y-%m-%d)

curl -i -X POST -d \
"{
  \"name\": \"test-update\",
  \"starts\": \"${TODAY}T00:00:00Z\",
  \"ends\": \"${TODAY}T23:59:59Z\",
  \"blocklists\": [\"${2}\"],
  \"devices\": [\"${3}\"]
}" \
http://localhost:8080/sessions/${1}
