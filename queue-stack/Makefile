.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf bin/queue-up
	
build:
	GOOS=linux GOARCH=amd64 go build -o bin/queue-up ./queue-up