FROM golang:latest as build
WORKDIR /go/src/github.com/jonathanfoster/freedom/
COPY . .
RUN make dep-build && CGO_ENABLED=0 GOOS=linux GOARCH=amd64  make build

FROM alpine:latest
EXPOSE 8080
COPY --from=build /go/src/github.com/jonathanfoster/freedom/bin/freedom-proxy /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/freedom-proxy"]
