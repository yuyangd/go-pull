# go-pull

## queue-stack

- SQS FIFO queue
- Source s3 bucket
- Lambda function that put object message to FIFO queue

## go-pull CLI Usage

```bash
go-pull ls # Inspect the s3 bucket

go-pull get # Download the object based on the FIFO queue
```

## Sample config

Generate the Config

```bash
go-pull config --stack-name `STACK_NAME` > $HOME/.go-pull.yaml
```

Inspect the config

```yaml
---
SQS_URL: "https://sqs.ap-southeast-2.amazonaws.com/<aws-account-id>/test-fifo-queue-ModelUpdatesSQSQueue-13LVZGQ6PGZTS.fifo"
SOURCE_BUCKET: "<bucket name>"
```

Or environment variable

```bash
export SQS_URL="https://sqs.ap-southeast-2.amazonaws.com/<aws-account-id>/test-fifo-queue-ModelUpdatesSQSQueue-13LVZGQ6PGZTS.fifo"
export SOURCE_BUCKET="<bucket name>"
```
