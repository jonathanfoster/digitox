SHELL=/bin/bash
SRC=$(shell go list -f '{{ .Dir }}' ./...)
VERSION=0.1.0

all: clean build

.PHONY: build
build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/digitox-apiserver ./cmd/apiserver

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
	docker build -t digitox/digitox .

.PHONY: docker-push
docker-push: docker-build
	$(shell aws ecr get-login --region us-east-1)
	docker tag digitox/digitox:latest digitox/digitox:$(VERSION)
	docker tag digitox/digitox:latest 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox:latest
	docker tag digitox/digitox:$(VERSION) 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox:latest

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

.PHONY: release
release: precommit docker-push

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

.PHONY: cover-html
cover-html:
	go tool cover -html=./bin/coverage.out
