#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "blocklist not provided" 1>&2
    exit 1
fi

TODAY=$(date +%Y-%m-%d)

curl -i -X POST -d \
"{
  \"name\": \"test\",
  \"starts\": \"${TODAY}T00:00:00Z\",
  \"ends\": \"${TODAY}T23:59:59Z\",
  \"blocklists\": [\"${1}\"]
}" \
http://localhost:8080/sessions/