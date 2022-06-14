.PHONY: lint build test

profile := "tw-beach-push"

default: build

deps:
	go mod tidy

lint:
	go vet ./...

build:
	go build -v ./...

test:
	AWS_PROFILE=$(profile) go test ./...

login:
	terraform login

region := us-east-1
vpc_id := vpc-00d272c5db70c2527
show-vpc:
	aws --profile $(profile) --region $(region) ec2 describe-vpcs
