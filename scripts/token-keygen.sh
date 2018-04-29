#!/bin/bash
set -e

if [ -z "$1" ]
then
    echo "output private key file not provided" 1>&2
    exit 1
fi

if [ -z "$2" ]
then
    echo "output public key file not provided" 1>&2
    exit 1
fi

openssl genrsa -out $1 2048
openssl rsa -in $1 -outform PEM -pubout -out $2
