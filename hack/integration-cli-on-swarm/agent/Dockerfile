# this Dockerfile is solely used for the master image.
# Please refer to the top-level Makefile for the worker image.
FROM golang:1.7
ADD . /go/src/github.com/ellcrys/docker/hack/integration-cli-on-swarm/agent
RUN go build -buildmode=pie -o /master github.com/ellcrys/docker/hack/integration-cli-on-swarm/agent/master
ENTRYPOINT ["/master"]
