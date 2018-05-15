#!/bin/bash
set -e

curl -s "http://localhost:8080/sessions/?access_token=$DIGITOX_ACCESS_TOKEN"
