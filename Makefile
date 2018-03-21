SHELL=/bin/bash
SRC=$(shell go list -f '{{ .Dir }}' ./...)
VERSION=0.1.0.1

all: clean build

build: build-apiserver build-proxy

build-apiserver:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/freedom-apiserver ./cmd/apiserver

build-proxy:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/freedom-proxy ./cmd/proxy

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

deploy: deploy-apiserver deploy-proxy

deploy-apiserver:
	kubectl apply -f ./k8s/apiserver/

deploy-proxy:
	kubectl apply -f ./k8s/proxy/

docker-build: docker-build-apiserver docker-build-proxy

docker-push: docker-push-apiserver docker-push-proxy

docker-build-apiserver:
	docker build -t freedom/freedom-apiserver -f ./build/apiserver/Dockerfile .

docker-build-proxy:
	docker build -t freedom/freedom-proxy -f ./build/proxy/Dockerfile .

docker-push-apiserver: docker-build-apiserver
	$(shell aws ecr get-login --region us-east-1)
	docker tag freedom/freedom-apiserver:latest 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:latest
	docker tag freedom/freedom-apiserver:latest freedom/freedom-apiserver:$(VERSION)	
	docker tag freedom/freedom-apiserver:$(VERSION) 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-apiserver:latest

docker-push-proxy: docker-build-proxy
	$(shell aws ecr get-login --region us-east-1)
	docker tag freedom/freedom-proxy:latest 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-proxy:latest
	docker tag freedom/freedom-proxy:latest freedom/freedom-proxy:$(VERSION)
	docker tag freedom/freedom-proxy:$(VERSION) 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-proxy:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-proxy:$(VERSION)
	docker push 672132384976.dkr.ecr.us-east-1.amazonaws.com/freedom/freedom-proxy:latest

fmt:
	echo "[fmt] Formatting code"
	@gofmt -s -w $(SRC)

fmt-check:
	@gofmt -l -s $(SRC) | read && echo "[fmt-check] Format check failed" 1>&2 && exit 1 || exit 0

lint:
	gometalinter --vendor ./...

precommit: fmt-check lint test

release: docker-push deploy

run-apiserver:
	go run ./cmd/apiserver/main.go

run-proxy:
	go run ./cmd/proxy/main.go

test:
	go test -v ./...
