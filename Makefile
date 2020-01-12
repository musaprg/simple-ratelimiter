TARGET=simple-ratelimiter

all: build

.PHONY: build
build:
	go build -o $(TARGET)

.PHONY: run
run: build
	./$(TARGET)

.PHONY: load-test
load-test:
	scripts/run_test.sh

.PHONY: clean
clean:
	go clean

