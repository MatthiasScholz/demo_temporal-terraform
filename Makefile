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

temporal-up: temporal-cleanup
	docker compose up -d
	docker compose logs
	docker compose ps
	open http://localhost:8080

temporal-cleanup:
	docker rm -f temporal-elasticsearch
	docker rm -f temporal-postgresql
	docker rm -f temporal
	docker rm -f temporal-admin-tools
	docker rm -f temporal-ui
	docker rm -f temporal-web

temporalite-install:
	go install github.com/DataDog/temporalite/cmd/temporalite@latest

temporalite-start:
	~/go/bin/temporalite start --namespace default -f temporalite.db

temporalite-ui:
	open http://localhost:8233

check:
	docker compose exec temporal-admin-tools tctl cluster health
	docker compose exec temporal-admin-tools tctl namespace list
	docker compose exec temporal-admin-tools tctl workflow listall

login:
	terraform login

region := us-east-1
vpc_id := vpc-00d272c5db70c2527
show-vpc:
	aws --profile $(profile) --region $(region) ec2 describe-vpcs

update-docker-compose:
	rm -f docker-compose.yml
	curl -O https://raw.githubusercontent.com/temporalio/docker-compose/main/docker-compose.yml
	curl -o dynamicconfig/development-sql.yaml https://raw.githubusercontent.com/temporalio/docker-compose/main/dynamicconfig/development-sql.yaml
