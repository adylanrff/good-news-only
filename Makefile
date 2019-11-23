# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=good-news-only.o

all:
	run clear

run:
	${GOBUILD} -o ${BINARY_NAME} -v ./...
	./${BINARY_NAME}

clear:
	rm -rf *.o
