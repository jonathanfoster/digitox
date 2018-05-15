#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "session not provided" >&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "session name not provided" >&2
    exit 1
fi

if [ -z "$3" ]
then
    echo "blocklist not provided" >&2
    exit 1
fi

TODAY=$(date +%Y-%m-%d)

curl -s -X PUT -d \
"{
  \"name\": \"${2}\",
  \"starts\": \"${TODAY}T00:00:00Z\",
  \"ends\": \"${TODAY}T23:59:59Z\",
  \"blocklists\": [{\"id\":\"${3}\"}]
}" \
"http://localhost:8080/sessions/${1}?access_token=$DIGITOX_ACCESS_TOKEN"
