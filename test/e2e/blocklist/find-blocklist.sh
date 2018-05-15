#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "blocklist not provided" >&2
    exit 1
fi

curl -s "http://localhost:8080/blocklists/${1}?access_token=$DIGITOX_ACCESS_TOKEN"
