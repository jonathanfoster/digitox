#!/bin/bash
set -e

curl -i -X POST -d \
'{
  "name": "test-create",
  "domains": ["www.reddit.com"]
}' \
http://localhost:8080/blocklists/
