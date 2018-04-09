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
* [ ] CI/CD: Run unit tests during CI build

## API Server

* [X] API: Create API server
* [X] API: Create session routes
* [X] API: Create block list routes
* [X] API: Create device routes

## Session

* [X] Session: Handle list sessions
* [X] Session: Handle find session
* [X] Session: Handle remove session
* [X] Session: Handle save session
* [X] Session: Validate model before persisting
* [ ] Session: Validate blocklists exist before persisting

## Blocklist

* [X] Blocklist: Handle list blocklists
* [X] Blocklist: Handle find blocklist
* [X] Blocklist: Handle remove blocklist
* [X] Blocklist: Handle save blocklist
* [X] Blocklist: Validate model before persisting

## Proxy

* [X] Proxy: Start session by copying active session blocklists to /etc/squid/blocklist
* [ ] Proxy: End session by removing /etc/squid/blocklist - **IN PROGRESS**
* [ ] Proxy: Restrict proxy access to devices - use OpenVPN instead?

## Devices

* [ ] Device: Handle list devices
* [ ] Device: Handle find device
* [ ] Device: Handle remove device
* [ ] Device: Handle save device
* [ ] Device: Store devices in database
* [ ] Device: Validate model before persisting

## Authorization

* [ ] Auth: Restrict access using auth token

## Misc

* [ ] Misc: Rename project to Digitox or Digiclutter
* [ ] Misc: Create freedomctl CLI
