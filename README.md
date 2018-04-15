# Digitox

[![Go Report Card](https://goreportcard.com/badge/github.com/jonathanfoster/digitox?style=flat-square)](https://goreportcard.com/report/github.com/jonathanfoster/digitox)
[![Coverage](https://codecov.io/gh/jonathanfoster/digitox/branch/master/graph/badge.svg)](https://codecov.io/gh/jonathanfoster/digitox)
[![Build Status](https://img.shields.io/travis/jonathanfoster/digitox.svg?style=flat-square&&branch=master)](https://travis-ci.org/jonathanfoster/digitox)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/jonathanfoster/digitox)
[![Release](https://img.shields.io/github/release/jonathanfoster/digitox.svg?style=flat-square)](https://github.com/jonathanfoster/digitox/releases/latest)

Digitox is a self-hosted website blocker that allows you to block websites on your laptop and mobile devices without relying upon third-party services like [Freedom](https://freedom.refersion.com/c/ddb297).

## Features

* Block websites
* Schedule sessions
* Customize blocklists
* Manage devices
* Secure connection

## Getting Started

1. Clone this repo

    ```bash
    git clone git@github.com:jonathanfoster/digitox.git $GOPATH/src/github.com/jonathanfoster/digitox
    cd $GOPATH/src/github.com/jonathanfoster/digitox
    ```

2. Build

    ```bash
    make
    ```

3. Run

    ```bash
    docker-compose up -d
    ```

## Testing

You can use a series of shell scripts to perform end-to-end testing locally:

```bash
# Make sure all scripts are executable
chmod +x ./test/e2e
cd test/e2e

# Start containers (drop `-d` if prefer to see real-time logs and open another terminal window to run the test scripts)
docker-compose up -d

# Confirm proxy allows all requests when no session is active
./get-proxy-allow.sh
./get-proxy-deny.sh

# Start session
# You'll need to copy the blocklist ID return for use when creating the session
./post-blocklist.sh
./post-session.sh $BLOCKLIST_ID

# Confirm proxy allows non-blocked requests and denies blocked requests
./get-proxy-allow.sh
./get-proxy-deny.sh

# End session
./delete-session.sh $BLOCKLIST_ID

# Confirm proxy allows all requests now that sesson has ended
./get-proxy-allow.sh
./get-proxy-deny.sh
```