# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BIN_DIRECTORY=bin
BINARY_NAME=${BIN_DIRECTORY}/good-news-only

all:
	build run clear

build: 
	mkdir -p ${BIN_DIRECTORY}
	${GOBUILD} -o ${BIN_DIRECTORY} -v ./...

run:
	./${BINARY_NAME}

clear:
	rm -rf *.o
