APP = algohub
APISERVER = apiserver

all: run

run: build
	@echo "Running..."
	./bin/$(APP) apiserver

build: clean
	@echo "Building..."
	go build -o ./bin/$(APP) ./cmd/$(APISERVER)

clean:
	@echo "Cleaning..."
	rm -rf bin/

.PHONY: all run build clean
.DEFAULT_GOAL := all