#!/bin/bash
set -e

curl -i http://localhost:8080/sessions/?access_token=$DIGITOX_ACCESS_TOKEN
