SHELL=/bin/bash
SRC=$(shell go list -f '{{ .Dir }}' ./...)
VERSION=0.1.0

all: clean build

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/digitox-apiserver ./cmd/apiserver

clean:
	rm -rf bin/

dep: dep-build dep-dev

dep-build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

dep-deploy:
	./scripts/install-kops.sh
	./scripts/install-kubectl.sh

dep-dev:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

deploy: 
	kubectl apply -f ./k8s/

docker-build:
	docker build -t digitox/digitox-apiserver .

docker-push: docker-build
	$(shell aws ecr get-login --region us-east-1)
	docker tag digitox/digitox-apiserver:latest 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox-apiserver:latest
	docker tag digitox/digitox-apiserver:latest digitox/digitox-apiserver:$(VERSION)
	docker tag digitox/digitox-apiserver:$(VERSION) 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox-apiserver:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox-apiserver:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/digitox/digitox-apiserver:latest

fmt:
	@echo "[fmt] Formatting code"
	@gofmt -s -w $(SRC)

fmt-check:
	@gofmt -l -s $(SRC) | read && echo "[fmt-check] Format check failed" 1>&2 && exit 1 || exit 0

imports:
	@echo "[imports] Formatting imported packages and code"
	@goimports -w $(SRC)

imports-check:
	@goimports -l $(SRC) | read && echo "[imports-check] Imports check failed" 1>&2 && exit 1 || exit 0

lint:
	gometalinter --vendor ./...

precommit: fmt-check imports-check lint test

release: precommit docker-push deploy

run: build
	mkdir -p bin/test/
	./bin/digitox-apiserver -v --sessions bin/test/sessions --blocklists bin/test/blocklists --proxylist bin/test/blocklist --tick 10s

.PHONY: test
test: clean
	mkdir -p bin/
	go test -coverprofile=./bin/coverage.out -v ./...
	go tool cover -func=./bin/coverage.out

cover-html:
	go tool cover -html=./bin/coverage.out
