BINARY_NAME=bqt

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} ./src/main/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ./src/main/main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ./src/main/main.go

run:
	go run ./src/main/main.go

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

test:
	go test ./src/main/
	go test ./src/templater/
	go test ./src/mapper/

dep:
	go mod download

