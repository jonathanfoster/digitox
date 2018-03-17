SHELL=/bin/bash
SRC=$(shell go list -f '{{ .Dir }}' ./...)
VERSION=0.1.0

all: clean build

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/freedom-apiserver ./cmd/apiserver

build-docker:
	docker build -t jonathanfoster/freedom-apiserver .

clean:
	rm -rf bin/

dep: dep-build dep-dev

dep-build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

dep-deploy:
	./install-kops.sh
	./install-kubectl.sh

dep-dev:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

deploy:
	kubectl apply -f ./deployment.yml

fmt:
	echo "[fmt] Formatting code"
	@gofmt -s -w $(SRC)

fmt-check:
	@gofmt -l -s $(SRC) | read && echo "[fmt-check] Format check failed" 1>&2 && exit 1 || exit 0

lint:
	gometalinter --vendor ./...

precommit: fmt-check lint test

run:
	go run ./cmd/apiserver/main.go

test:
	go test -v ./...
