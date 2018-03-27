# Roadmap

This document provides a high-level roadmap for Freedom development.

Prioritization follows the principles found in Ron Jeffries' [The Nature of Software Development](https://www.amazon.com/gp/product/1941222374/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1941222374&linkCode=as2&tag=github-jonathanfoster-freedom-20&linkId=eb959b758bf93091a58f633b92397024) (if you only read one Agile book, read this one).

* Keep it simple
* Make it valuable
* Build it piece by piece

## CI/CD [In Progress]

* [X] CI/CD: Build API server docker image
* [X] CI/CD: Run proxy and API server containers with shared data store
* [X] CI/CD: Create deployment pipeline using Kubernetes

## API Server

* [X] API: Create API server
* [X] API: Create session routes
* [X] API: Create block list routes
* [X] API: Create device routes

## Session

* [X] Session: Handle list sessions
* [ ] Session: Handle find session
* [ ] Session: Handle remove session
* [ ] Session: Handle save session
* [ ] Session: Store sessions in /etc/freedom/session/
* [ ] Session: Start session by copy block lists to /etc/squid/blocklist/block
* [ ] Session: End session by removing /etc/squid/blocklist/block
* [ ] Session: Validate model before persisting

## Block List

* [X] Block List: Handle list block lists
* [X] Block List: Handle find block list
* [X] Block List: Handle remove block list
* [X] Block List: Handle save block list
* [X] Block List: Store block lists in /etc/freedom/blocklist/
* [X] Block List: Validate model before persisting

## Devices

* [ ] Device: Handle list devices
* [ ] Device: Handle find device
* [ ] Device: Handle remove device
* [ ] Device: Handle save device
* [ ] Device: Store devices in /etc/freedom/passwd
* [ ] Device: Validate model before persisting
