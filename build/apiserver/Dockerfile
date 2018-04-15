FROM golang:latest as build
WORKDIR /go/src/github.com/jonathanfoster/digitox/
COPY . .
RUN make dep-build && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

FROM alpine:latest
EXPOSE 8080
COPY --from=build /go/src/github.com/jonathanfoster/digitox/bin/digitox-apiserver /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/digitox-apiserver"]
