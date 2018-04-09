#!/bin/bash
set -e

curl -i -X POST -d \
'{
  "name": "test",
  "domains": ["www.reddit.com"]
}' \
http://0.0.0.0:8080/blocklists/
