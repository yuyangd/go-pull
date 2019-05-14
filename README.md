# go-pull

queue-stack:

- SQS FIFO queue
- Source s3 bucket
- Lambda function that put object event to FIFO queue

go-pull cmd:

Command to pull from FIFO queue

## Sample config

$HOME/.go-pull.yaml

```yaml
---
SQS_URL: "https://sqs.ap-southeast-2.amazonaws.com/<aws-account-id>/test-fifo-queue-ModelUpdatesSQSQueue-13LVZGQ6PGZTS.fifo"
```

Or set environment variable

```bash
export SQS_URL="https://sqs.ap-southeast-2.amazonaws.com/<aws-account-id>/test-fifo-queue-ModelUpdatesSQSQueue-13LVZGQ6PGZTS.fifo"
```

## Usage