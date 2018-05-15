#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "device name not provided" >&2
    exit 1
fi

curl -s -i -X DELETE "http://localhost:8080/devices/${1}?access_token=$DIGITOX_ACCESS_TOKEN"
