DOCKER_IMAGE_TAG = $(git describe --abbrev=0 --tags)
DOCKER_REGISTRY ?= docker.io
DOCKER_IMAGE_NAME ?= clouina/go-rest-api

.PHONY: build_local_binary
build_local_binary: 
	go mod tidy
	go build -o ./cmd/bin/go-rest-api ./cmd/main.go

.PHONY: clean_local_binary
clean_local_binary:	
	rm -rf ./cmd/bin

.PHONY: build 
build:
	go mod tidy
	docker build -t ${DOCKER_REGISTRY}/${DOCKER_IMAGE_NAME}:latest .

.PHONY: publish 
publish:
	docker image tag ${DOCKER_REGISTRY}/${DOCKER_IMAGE_NAME}:latest ${DOCKER_REGISTRY}/${DOCKER_IMAGE_NAME}:$(DOCKER_IMAGE_TAG)
	docker push ${DOCKER_REGISTRY}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG}
	docker push ${DOCKER_REGISTRY}/${DOCKER_IMAGE_NAME}:latest

.PHONY: run 
run:
	go mod tidy
	docker-compose up --build

.PHONY: stop 
stop: 
	docker-compose down

.PHONY: cleanup
cleanup:
	docker-compose down --remove-orphans
	docker volume rm go-rest-api_pgdata