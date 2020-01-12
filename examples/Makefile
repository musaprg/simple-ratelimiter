TARGET=simple-server

all: build

.PHONY: build
build:
	go build -o $(TARGET)

.PHONY: run
run: build
	./$(TARGET)

.PHONY: clean
clean:
	rm $(TARGET)

