APP?=get2json

default:
	echo ${APP}

.PHONEY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o bin/${APP} main.go

all: build test

.PHONY: run
## run: runs go run main.go
run:
	go run main.go get


.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@rm -rf bin/${APP}

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

.PHONY: docker-build
## docker-build: build the docker image
docker-build:
	docker buildx build --platform linux/amd64 -t get2json .

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
