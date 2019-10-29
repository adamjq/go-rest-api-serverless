BINARY_NAME=main

default: clean validate build bundle

validate:
	sam validate
	cfn-lint -t template.yml

#install:
#
#test:

clean:
	rm -f ./packaged.yml
	rm -f src/$(BINARY_NAME)
	rm -f src/$(BINARY_NAME).zip

build:
	cd src && GOARCH=amd64 GOOS=linux go build -o $(BINARY_NAME)

bundle:
	cd src && zip main.zip $(BINARY_NAME)

package:
	echo "package cloudformation template..."
	aws cloudformation package \
		--template-file template.yml \
		--output-template-file packaged.yml \
		--s3-bucket "${S3_BUCKET}" \
		--s3-prefix sam

deploy:
	echo "deploy stack ${STACK_NAME}..."
	sam deploy \
		--template-file packaged.yml \
		--stack-name "${STACK_NAME}" \
		--capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM

deploy-stack: clean validate build bundle package deploy
