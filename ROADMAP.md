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

## Proxy Server

* [ ] Proxy: Replace Squid with [Oxy](https://github.com/vulcand/oxy) - **IN PROGRESS**

## API Server

* [X] API: Create API server
* [ ] API: Create session routes
* [ ] API: Create block list routes
* [ ] API: Create device routes

## Session

* [ ] Session: Handle list sessions
* [ ] Session: Handle find session
* [ ] Session: Handle remove session
* [ ] Session: Handle save session
* [ ] Session: Store sessions in /etc/freedom/session/
* [ ] Session: Start session by copy block lists to /etc/freedom/blocklist/block
* [ ] Session: End session by removing /etc/freedom/blocklist/block

## Block List

* [ ] Block List: Handle list block lists
* [ ] Block List: Handle find block list
* [ ] Block List: Handle remove block list
* [ ] Block List: Handle save block list
* [ ] Block List: Store block lists in /etc/freedom/blocklist/

## Devices

* [ ] Device: Handle list devices
* [ ] Device: Handle find device
* [ ] Device: Handle remove device
* [ ] Device: Handle save device
* [ ] Device: Store devices in /etc/freedom/passwd
