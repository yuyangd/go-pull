.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf bin/queue-up
	
build:
	GOOS=linux GOARCH=amd64 go build -o bin/queue-up ./queue-up

pkg:
	aws cloudformation package --template-file template.yaml --output-template-file packaged.yaml --s3-bucket finx-lambda-dev

ship:
	aws cloudformation deploy --template-file packaged.yaml --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM --stack-name test-fifo-queue --tags expires=2019-05-15