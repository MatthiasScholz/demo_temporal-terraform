.PHONY: lint build test

profile := "tw-beach-push"

default: build

deps:
	go mod tidy

lint:
	go vet ./...

build: deps
	go build -v ./...

worker:
	go run cmd/main.go

signal: deps
	go run client/network.go

test: deps
	AWS_PROFILE=$(profile) go test ./...

# FIXME NOT WORKING Temporal not starting up - need to recheck after colima restart !
temporal-up:
	docker compose up -d
	docker compose logs
	docker compose ps
	open http://localhost:8088

temporalite-install:
	go install github.com/DataDog/temporalite/cmd/temporalite@latest

temporalite-start:
	~/go/bin/temporalite start --namespace default -f temporalite.db

temporalite-ui:
	open http://localhost:8233

check:
	tctl namespace list
	tctl workflow list

login:
	terraform login

region := us-east-1
vpc_id := vpc-00d272c5db70c2527
show-vpc:
	aws --profile $(profile) --region $(region) ec2 describe-vpcs
