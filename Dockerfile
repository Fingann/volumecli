FROM golang:1.21rc2-bullseye

RUN apt-get update && apt-get install -y libasound2-dev
WORKDIR /go/src/

# Golang is is the entrypoint

