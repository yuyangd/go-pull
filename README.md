# go-pull

## queue-stack

- SQS FIFO queue
- Source s3 bucket
- Lambda function that put object message to FIFO queue

Provision the queue stack

```bash
cd queue-stack
make provision -e STACK=<stack name> -e LAMBDA_BUCKET=<bucket name>
```

## go-pull CLI Usage

Download the binary from release

```bash
go-pull ls # Inspect the s3 bucket

go-pull get # Download the object based on the FIFO queue
```

## Sample config

Generate the config $HOME/.go-pull.yaml

```bash
make url # Get SQS URL
```

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
