# Freedom

Freedom is a self-hosted website blocker that allows you to block websites on your laptop and mobile devices without relying upon
third-party services like [Freedom](https://freedom.to/).

## Features

* Block websites
* Advanced scheduling
* Customizable block lists

## Roadmap

* [X] Server: Create API server
* [X] CI/CD: Build API server docker image
* [X] CI/CD: Run proxy and API server containers with shared data store
* [ ] CI/CD: Create deployment pipeline using Kubernetes
* [ ] Server: Create session routes
* [ ] Server: Create block list routes
* [ ] Server: Create device routes
* [ ] Session: Handle list sessions
* [ ] Session: Handle find session
* [ ] Session: Handle remove session
* [ ] Session: Handle save session
* [ ] Session: Store sessions in /etc/freedom/session/
* [ ] Session: Start session by copy block lists to /etc/freedom/blocklist/block
* [ ] Session: End session by removing /etc/freedom/blocklist/block
* [ ] Block List: Handle list block lists
* [ ] Block List: Handle find block list
* [ ] Block List: Handle remove block list
* [ ] Block List: Handle save block list
* [ ] Block List: Store block lists in /etc/freedom/blocklist/
* [ ] Device: Handle list devices
* [ ] Device: Handle find device
* [ ] Device: Handle remove device
* [ ] Device: Handle save device
* [ ] Device: Store devices in /etc/freedom/passwd
