GO = go
# GOFLAGS = -ldflags="-s -w"
MAIN_PATH = ./cmd/main.go
BIN_PATH = ./bin
BIN_NAME = gotty-server

.PHONY: all
all: clean build run

.PHONY: build
build:
	$(GO) build -o $(BIN_PATH)/$(BIN_NAME) $(MAIN_PATH)

.PHONY: run
run: clean build
	$(BIN_PATH)/$(BIN_NAME)

.PHONY: clean
clean:
	rm -rf $(BIN_PATH)

.PHONY: test
test: 
	$(GO) test ./...
