# Roadmap

This document provides a high-level roadmap for Digitox development.

## CI/CD

* [X] CI/CD: Build API server docker image
* [X] CI/CD: Run proxy and API server containers with shared data store
* [X] CI/CD: Create deployment pipeline using Kubernetes
* [X] CI/CD: Run unit tests during CI build

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
* [X] Session: Validate blocklists exist before persisting

## Blocklist

* [X] Blocklist: Handle list blocklists
* [X] Blocklist: Handle find blocklist
* [X] Blocklist: Handle remove blocklist
* [X] Blocklist: Handle save blocklist
* [X] Blocklist: Validate model before persisting

## Proxy

* [X] Proxy: Start session by copying active session blocklists to /etc/squid/blocklist
* [X] Proxy: End session by removing /etc/squid/blocklist
* [X] Proxy: Restart proxy after blocklist update
* [X] Proxy: Restart proxy immediately after session or blocklist change
* [X] Proxy: Restrict proxy access to devices

## Devices

* [X] Device: Handle list devices
* [X] Device: Handle find device
* [X] Device: Handle remove device
* [X] Device: Handle save device
* [X] Device: Validate model before persisting

## Authorization

* [X] Auth: Restrict access using auth token

## Misc

* [X] Misc: Rename project to Digitox
* [X] Misc: Combine Squid and API server containers
* [ ] Misc: Create digitoxctl CLI
* [ ] Misc: Switch to Ginkgo for testing
* [ ] Misc: Generate OpenAPI spec for API
* [X] Misc: Use SQLite as data store
* [ ] Misc: Update README with Docker instructions

## Someday Maybe

* [ ] Security: Generate OAuth credentials if not provided (no default)
* [ ] Security: Generate RSA key pair if not provided (no default)
