#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "session not provided" 1>&2
    exit 1
fi

curl -i http://localhost:8080/sessions/${1}
