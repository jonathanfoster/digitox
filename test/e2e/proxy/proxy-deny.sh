#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "device name not provided" 1>&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "device password not provided" 1>&2
    exit 1
fi

curl -I -x ${1}:${2}@localhost:3128 https://www.reddit.com
