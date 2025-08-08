# Build targets
.PHONY: all
all: build test lint coverage

build:
	go mod verify && go mod tidy
	cd cmd/cli && go build main.go && rm main
	

test:
	cd cmd/cli && go test -covermode=atomic
	#cd auth && go test -covermode=atomic
	#cd cli && go test -covermode=atomic
	#cd config && go test -covermode=atomic
	#go test -covermode=atomic
	#go test ./... -v
	
lint:
	golangci-lint run > lint.txt

coverage:
	go test ./... -coverprofile=coverage.txt -covermode=atomic
	go tool cover -func=coverage.txt

clean:
	rm -f coverage.txt lint.txt
	go clean
