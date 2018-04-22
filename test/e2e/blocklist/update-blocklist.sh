#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "blocklist not provided" 1>&2
    exit 1
fi

curl -i -X POST -d \
'{
  "name": "test-update",
  "domains": ["www.reddit.com"]
}' \
http://localhost:8080/blocklists/${1}

