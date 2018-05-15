#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "session name not provided" >&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "blocklist not provided" >&2
    exit 1
fi

TODAY=$(date +%Y-%m-%d)

curl -s -X POST -d \
"{
  \"name\": \"${1}\",
  \"starts\": \"${TODAY}T00:00:00Z\",
  \"ends\": \"${TODAY}T23:59:59Z\",
  \"blocklists\": [{\"id\":\"${2}\"}]
}" \
"http://localhost:8080/sessions/?access_token=$DIGITOX_ACCESS_TOKEN"
