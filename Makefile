.PHONY: build run clean

# Docker image name
IMAGE_NAME := volumecli

# Docker container name
CONTAINER_NAME := volumecli-container

# Local directory for the binary
BIN_DIR := ./bin

# Binary name
BINARY_NAME := volumecli

CONTAINER_WORKDIR := /go/src

# Binary path in the container
CONTAINER_BIN_PATH := ${CONTAINER_WORKDIR}/${BIN_DIR}/$(BINARY_NAME)

build:
	@docker build -t $(IMAGE_NAME) .

run: build
	@mkdir -p $(BIN_DIR)
	docker run --name $(CONTAINER_NAME) -v $(PWD):${CONTAINER_WORKDIR} $(IMAGE_NAME) go build -o $(CONTAINER_BIN_PATH) main.go
	@docker rm $(CONTAINER_NAME)

clean:
	@docker rmi -f $(IMAGE_NAME)
	@rm -f $(LOCAL_BIN_PATH)

all: run
