.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgresql://postgres:password@0.0.0.0:5432/postgres?sslmode=disable

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) cmd/cmd.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

swag:
	swag init

docker.run: docker.network swag docker.stickerfy docker.redis docker.mongo docker.zookeeper docker.kafka

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.stickerfy.build:
	docker build -t stickerfy .

docker.stickerfy: docker.stickerfy.build
	docker run --rm -d \
		--name stickerfy \
		--network dev-network \
		-p 5000:5000 \
		stickerfy

docker.postgres:
	docker run --rm -d \
		--name postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev/postgres-dev/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.redis:
	docker run --rm -d \
		--name redis \
		--network dev-network \
		-p 6379:6379 \
		redis

docker.mongo:
	docker run --rm -d \
		--name mongo \
		--network dev-network \
		-p 27017:27017 \
		mongo

docker.zookeeper:
	docker run --rm -d \
		--name zookeeper \
		--network dev-network \
		-p 2181:2181 \
    -e ZOOKEEPER_CLIENT_PORT: 2181 \
    -e ZOOKEEPER_TICK_TIME: 2000 \
		confluentinc/cp-zookeeper:7.3.0

docker.kafka:
	docker run --rm -d \
		--name broker \
		--network dev-network \
		-p 9092:9092 \
    -e KAFKA_BROKER_ID: 1 \
    -e KAFKA_ZOOKEEPER_CONNECT: 'zookeeper:2181' \
    -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT \
    -e KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://broker:29092 \
    -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1 \
    -e KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1 \
    -e KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1 \
		confluentinc/cp-kafka:7.3.0

docker.stop: docker.stop.stickerfy docker.stop.redis docker.stop.mongo docker.stop.zookeeper docker.stop.kafka

docker.stop.stickerfy:
	docker stop stickerfy

docker.stop.postgres:
	docker stop postgres

docker.stop.redis:
	docker stop redis

docker.stop.mongo:
	docker stop mongo

docker.stop.zookeeper:
	docker stop zookeeper

docker.stop.kafka:
	docker stop broker
