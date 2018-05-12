#!/bin/bash
set -e

curl -s "http://localhost:8080/devices/?access_token=$DIGITOX_ACCESS_TOKEN"
