include vars.mk
include $(ENV_PATH)

# !!! This variables should be defined in vars.mk !!!
# DAYS?=30
# VPS?=127.0.0.1
# SERVER?=server
# COMPANY?=Company
# USER=
# SUDOPASS=
# BASE_VERSION?=0.0a
# DOCKER_NETWORK_NAME=
# ENV_PATH=
# REMOTE_DB_IP=
# REMOTE_DB_PORT=

GO=$(GOROOT)/bin/go

RUN_ON_VPS=echo $(SUDOPASS) | ssh -tt $(SERVER)

DATETIME=$(shell date "+%Y%m%d%H%M%S")
VERSION=$(shell echo $(BASE_VERSION).$(DATETIME))

.PHONY: build
build:
	@echo "Building..."
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GO) build -o dist/app ./cmd/bot
	@echo "Done"

.PHONY: generate-certificates
generate-certificates:
	@echo "Generate new SSL certificates..."
	@openssl req -newkey rsa:2048 -sha256 -nodes -keyout certificates/key.pem -x509 -days $(DAYS) -out certificates/cert.pem -subj "/C=KZ/ST=Almaty/L=Almaty/O=$(COMPANY)/CN=$(VPS)"
	@echo "Done"

.PHONY: deploy
deploy: generate-certificates build #check-migration-on-server
	@echo "Copying binary..."
	@scp dist/app $(SERVER):bot/dist/app
	@echo "Copying ENV file..."
	@scp $(ENV_PATH) $(SERVER):bot/00.env
	@echo "Copying Dockerfile file..."
	@scp Dockerfile $(SERVER):bot/Dockerfile
	@echo "Copying dockerignore..."
	@scp .dockerignore $(SERVER):bot/.dockerignore
	@echo "Copying certificates..."
	@scp certificates/cert.pem $(SERVER):bot/certificates/cert.pem
	@scp certificates/key.pem $(SERVER):bot/certificates/key.pem
	@echo "Docker building and starting..."
	@$(RUN_ON_VPS) "sudo docker image build -f bot/Dockerfile --build-arg EXPOSE_PORT=$(CHGKBOT_PORT) -t chgkbot bot"
	@$(RUN_ON_VPS) "sudo docker tag chgkbot:latest chgkbot:$(VERSION)"
	@$(RUN_ON_VPS) "sudo docker container rm -f chgk && sudo docker run -p 443:$(CHGKBOT_PORT) --network $(DOCKER_NETWORK_NAME) --env-file bot/00.env -d --name chgk chgkbot:$(VERSION)"
	@echo "Deploy completed"

include .test.env

.PHONY: run-local-docker-db
run-local-docker-db:
	@docker container rm -f database && docker run --name database -p $(TEST_DB_PORT):5432\
 		-e POSTGRES_PASSWORD=$(TEST_DB_PASS) -e POSTGRES_DB=$(TEST_DB_NAME) -e POSTGRES_USER=$(TEST_DB_USER)\
 		-d postgres
	@sleep 2

.PHONY: migrate
migrate:
	@migrate -database "postgres://$(TEST_DB_USER):$(TEST_DB_PASS)@$(TEST_DB_HOST):$(TEST_DB_PORT)/$(TEST_DB_NAME)?sslmode=disable" -path internal/migrations up

VARS=$(shell cat .test.env | xargs)

.PHONY: test
test: run-local-docker-db migrate
	@env $(VARS) $(GO) test -v ./...

.PHONY: run-db-on-server
run-db-on-server:
	@$(RUN_ON_VPS) "sudo docker network create $(DOCKER_NETWORK_NAME) || true"
	@$(RUN_ON_VPS) "sudo docker container rm -f database && sudo docker run --name database -p $(REMOTE_DB_PORT):5432\
 		--network $(DOCKER_NETWORK_NAME) -e POSTGRES_PASSWORD=$(CHGKBOT_DB_PASS) -e POSTGRES_DB=$(CHGKBOT_DB_NAME) -e POSTGRES_USER=$(CHGKBOT_DB_USER)\
 		-e PGDATA=/var/lib/postgresql/data/pgdata -v /home/$(USER)/db:/var/lib/postgresql/data -d postgres"

.PHONY: migrate-on-server
migrate-on-server:
	@echo "!!! Apply migrations on the remote server !!!"
	@migrate -database "postgres://$(CHGKBOT_DB_USER):$(CHGKBOT_DB_PASS)@$(REMOTE_DB_IP):$(REMOTE_DB_PORT)/$(CHGKBOT_DB_NAME)?sslmode=disable" -path internal/migrations up

.PHONY: check-migration-on-server
check-migration-on-server:
	@echo "======================================="
	@echo "Migration version on the remote server!"
	@migrate -database "postgres://$(CHGKBOT_DB_USER):$(CHGKBOT_DB_PASS)@$(REMOTE_DB_IP):$(REMOTE_DB_PORT)/$(CHGKBOT_DB_NAME)?sslmode=disable" -path internal/migrations version
	@echo "======================================="
