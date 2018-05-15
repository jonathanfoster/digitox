#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "device name not provided" >&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "device password not provided" >&2
    exit 1
fi

curl -s -X PUT -d \
"{
  \"name\": \"${1}\",
  \"password\": \"${2}\"
}" \
"http://localhost:8080/devices/${1}?access_token=$DIGITOX_ACCESS_TOKEN"
