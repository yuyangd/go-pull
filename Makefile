STACK?=test-fifo-queue

url: # SQS URL
	aws cloudformation describe-stacks --stack-name $(STACK) --query 'Stacks[0].Outputs[?OutputKey==`QueueURL`].OutputValue'

build: build-darwin build-linux

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/go-pull-darwin ./

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/go-pull-linux ./

build-windows:
	GOOS=windows GOARCH=386 go build -o bin/go-pull-windows ./	