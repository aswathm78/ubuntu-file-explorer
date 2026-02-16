BINARY_NAME=finder-clone

all: build

build:
	go build -o $(BINARY_NAME) ./cmd/finder

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

deps:
	go mod download
	go mod tidy

.PHONY: all build run clean deps
