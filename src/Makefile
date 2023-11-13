BINARY_NAME=weather


build:
	go build -o ${BINARY_NAME}

run: build
	bin/${BINARY_NAME}

test:
	go test -v main.go

clean:
	go clean
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

deps:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all

compile:
	echo "Compiling for every os"
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux main.go
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows main.go

all: 
	build test