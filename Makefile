SHELL=/bin/bash
SRC=$(shell go list -f '{{ .Dir }}' ./...)
VERSION=0.1.0.6

all: clean build

build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/freedom-apiserver ./cmd/apiserver

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
	kubectl apply -f ./k8s/

docker-build:
	docker build -t freedom/freedom-apiserver .

docker-push: docker-build
	$(shell aws ecr get-login --region us-east-1)
	docker tag freedom/freedom-apiserver:latest 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:latest
	docker tag freedom/freedom-apiserver:latest freedom/freedom-apiserver:$(VERSION)
	docker tag freedom/freedom-apiserver:$(VERSION) 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:latest

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
	./bin/freedom-apiserver

test:
	go test -coverprofile=./bin/coverage.out -v ./...
	go tool cover -func=./bin/coverage.out
	go tool cover -html=./bin/coverage.out
