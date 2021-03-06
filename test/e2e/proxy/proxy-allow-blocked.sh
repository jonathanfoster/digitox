#!/bin/bash

if [ -z "$1" ]
then
    echo "device name not provided" >&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "device password not provided" >&2
    exit 1
fi

HTTP_CODE=$(curl -s -o /dev/null -w '%{http_code}' -x ${1}:${2}@localhost:3128 https://news.ycombinator.com)
RETURN_CODE=$?
if [ $RETURN_CODE -ne 0 ]; then
   echo "error: curl exited with return code $RETURN_CODE" >&2
   exit $RETURN_CODE
fi

if [ $HTTP_CODE -eq 200 ]; then
    echo "pass: received http code $HTTP_CODE"
else
    echo "fail: received http code $HTTP_CODE" >&2
    exit 1
fi
