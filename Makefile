STACK?=test-fifo-queue

url: # SQS URL
	aws cloudformation describe-stacks --stack-name $(STACK) --query 'Stacks[0].Outputs[?OutputKey==`QueueURL`].OutputValue'
