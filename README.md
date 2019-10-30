# go-rest-api-serverless
Serverless Golang REST API hosted on AWS Lambda.

## Development

Run the Lambda API locally for testing:
```
make invoke-no-event
```

## Packaging and deployment

An S3 bucket must be created before deployment to hold the lambda code:

```bash
aws s3 mb s3://BUCKET_NAME
```

Set the follow environment variables:
```bash
export S3_BUCKET=golang-serverless-proxy-api
export STACK_NAME=golang-serverless-proxy-api
```

```bash
# Build and test during development
make

# Build, Package and Deploy
make deploy-stack
```
