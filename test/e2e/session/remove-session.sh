#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "session not provided" >&2
    exit 1
fi

curl -i -X DELETE "http://localhost:8080/sessions/${1}?access_token=$DIGITOX_ACCESS_TOKEN"
