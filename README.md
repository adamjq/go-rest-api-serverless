# go-rest-api-serverless
Serverless Golang REST API hosted on AWS Lambda.

## Requirements
- Go 1.18
- Docker
- AWS CDK

## Features

- Golang Echo REST server running on API Gateway backed by an AWS Lambda proxy integration
- Schema-first, code-generation using [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- Adapts [AWS Lambda Proxy response to Go Echo server request](https://github.com/awslabs/aws-lambda-go-api-proxy)
- Starts server in Lambda global context on Cold Start for reuse between invocations
- OpenAPI spec validation in CI [with swagger-cli](https://github.com/APIDevTools/swagger-cli)

### Not included:
- Auth - add BYO auth handling inside Echo server
- Doesn't build an API Gateway from the OpenAPI spec, for example with [a SpecRestAPI](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_apigateway.SpecRestApi.html).

## Development

Launch a development server with Docker:
```
docker-compose up
```

Run unit tests:
```
make test
```

Validate OpenAPI spec with:
```
make validate
```

## AWS Deployment

```
export AWS_PROFILE=<YOUR_PROFILE_NAME>
make deploy
```
