SHELL=/bin/bash
SRC=$(shell go list -f '{{ .Dir }}' ./...)
VERSION=0.1.0

all: clean build

.PHONY: build
build:
	go build -ldflags "-X main.version=${VERSION}" -o bin/digitox-apiserver ./cmd/apiserver

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: dep
dep: dep-build dep-dev

.PHONY: dep-build
dep-build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: dep-dev
dep-dev:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: docker-build
docker-build:
	docker build -t jonathanfoster/digitox .

.PHONY: docker-hub-build
docker-hub-build:
	curl -H "Content-Type: application/json" \
		--data '{"source_type": "Branch", "source_name": "master"}" \
		-X POST \
		https://registry.hub.docker.com/u/jonathanfoster/digitox/trigger/${DOCKER_HUB_TOKEN}/

.PHONY: fmt
fmt:
	@echo "[fmt] Formatting code"
	@gofmt -s -w $(SRC)

.PHONY: fmt-check
fmt-check:
	@gofmt -l -s $(SRC) | read && echo "[fmt-check] Format check failed" 1>&2 && exit 1 || exit 0

.PHONY: imports
imports:
	@echo "[imports] Formatting imported packages and code"
	@goimports -w $(SRC)

.PHONY: imports-check
imports-check:
	@goimports -l $(SRC) | read && echo "[imports-check] Imports check failed" 1>&2 && exit 1 || exit 0

.PHONY: lint
lint:
	gometalinter --vendor ./...

.PHONY: precommit
precommit: fmt-check imports-check lint test

.PHONY: run
run: build
	mkdir -p bin/test/
	./bin/digitox-apiserver -v \
	    --sessions bin/test/sessions \
	    --blocklists bin/test/blocklists \
	    --active bin/test/active \
	    --devices bin/test/passwd \
	    --ticker-duration 30s

.PHONY: test
test: clean
	mkdir -p bin/
	go test -coverprofile=./bin/coverage.out -v ./...
	go tool cover -func=./bin/coverage.out

.PHONY: test-codecov
test-codecov: test
    mv bin/coverage.out coverage.txt

.PHONY: test-cover-html
test-cover-html:
	go tool cover -html=./bin/coverage.out
