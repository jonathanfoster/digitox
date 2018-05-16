# Digitox

[![Go Report Card](https://goreportcard.com/badge/github.com/jonathanfoster/digitox)](https://goreportcard.com/report/github.com/jonathanfoster/digitox)
[![Coverage](https://codecov.io/gh/jonathanfoster/digitox/branch/master/graph/badge.svg)](https://codecov.io/gh/jonathanfoster/digitox)
[![Build Status](https://img.shields.io/travis/jonathanfoster/digitox.svg?style=flat-square&&branch=master)](https://travis-ci.org/jonathanfoster/digitox)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/jonathanfoster/digitox)
[![Release](https://img.shields.io/github/release/jonathanfoster/digitox.svg?style=flat-square)](https://github.com/jonathanfoster/digitox/releases/latest)

Digitox is a self-hosted website blocker that allows you to block sites on all your devices without relying upon third-party services like [Freedom](https://freedom.to/).

## Features

* Block websites
* Schedule sessions
* Customize blocklists

## Getting Started

1. Start Digitox container

    ```bash
    docker run -d -p 3128:3128 -p 8080:8080 --name digitox jonathanfoster/digitox
    ```

2. Get access token

    ```bash
    # Save access token in $DIGITOX_ACCESS_TOKEN
    curl "http://localhost:8080/oauth/token?grant_type=client_credentials&client_id=59f92849-b883-402c-b429-15a67663d4f3&client_secret=450a31ea-0c18-4925-97db-b9f981ca4a62&redirect_uri=http://localhost"
    ```

3. Create device

    ```bash
    curl -s -X POST -d \
    '{
      "name": "digitox",
      "password": "Digitox123"
    }' \
    "http://localhost:8080/devices/?access_token=$DIGITOX_ACCESS_TOKEN"
    ```

4. Create blocklist

    ```bash
    # Save blocklist ID in DIGITOX_BLOCKLIST_ID
    curl -s -X POST -d \
    '{
      "name\": "HackerNews",
      "domains\": ["news.ycombinator.com"]
    }' \
    "http://localhost:8080/blocklists/?access_token=$DIGITOX_ACCESS_TOKEN"
    ```

5. Create session

    ```bash
    TODAY=$(date +%Y-%m-%d)

    curl -s -X POST -d \
    "{
      \"name\": \"No HackerNews after 8PM",
      \"starts\": \"${TODAY}T20:00:00Z\",
      \"ends\": \"${TODAY}T23:59:59Z\",
      \"blocklists\": [{\"id\":\"$DIGITOX_BLOCKLIST_ID\"}]
    }" \
    "http://localhost:8080/sessions/?access_token=$DIGITOX_ACCESS_TOKEN"
    ```

6. Enjoy digital freedom

    ```bash
    # Configure browser proxy to use locahost:3128 and digitox:Digitox123
    curl -x digitox:Digitox123@localhost:3128 https://news.ycombinator.com
    ```

## Developing Locally

1. Clone this repo

    ```bash
    git clone git@github.com:jonathanfoster/digitox.git $GOPATH/src/github.com/jonathanfoster/digitox
    cd $GOPATH/src/github.com/jonathanfoster/digitox
    ```

2. Code

3. Test

    ```bash
    make lint
    make test
    make test-e2e
    ```

2. Run

    ```bash
    docker-compose up -d
    ```
