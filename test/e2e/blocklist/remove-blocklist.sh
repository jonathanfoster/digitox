#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "blocklist not provided" 1>&2
    exit 1
fi

curl -i -X DELETE http://localhost:8080/blocklists/${1}
