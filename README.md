# Freedom

[![Go Report Card](https://goreportcard.com/badge/github.com/jonathanfoster/freedom?style=flat-square)](https://goreportcard.com/report/github.com/jonathanfoster/freedom)
[![Coverage](https://codecov.io/gh/jonathanfoster/freedom/branch/master/graph/badge.svg)](https://codecov.io/gh/jonathanfoster/freedom)
[![Build Status](https://img.shields.io/travis/jonathanfoster/freedom.svg?style=flat-square&&branch=master)](https://travis-ci.org/jonathanfoster/freedom)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/jonathanfoster/freedom)
[![Release](https://img.shields.io/github/release/jonathanfoster/freedom.svg?style=flat-square)](https://github.com/jonathanfoster/freedom/releases/latest)

Freedom is a self-hosted website blocker that allows you to block websites on your laptop and mobile devices without relying upon third-party services like [Freedom](https://freedom.refersion.com/c/ddb297).

## Features

* Block websites
* Schedule sessions
* Customize block lists
* Manage devices

## Getting Started

1. Clone this repo

    ```bash
    git clone git@github.com:jonathanfoster/freedom.git $GOPATH/src/github.com/jonathanfoster/freedom
    cd $GOPATH/src/github.com/jonathanfoster/freedom
    ```

2. Build

    ```bash
    make
    ```

3. Run

    ```bash
    ./bin/freedom-apiserver
    ```
