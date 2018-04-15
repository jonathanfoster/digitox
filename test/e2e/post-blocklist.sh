#!/bin/bash
set -e

curl -i -X POST -d \
'{
  "name": "test",
  "domains": ["www.reddit.com"]
}' \
http://localhost:8080/blocklists/
