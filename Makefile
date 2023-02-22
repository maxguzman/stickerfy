.PHONY: clean critic security lint test build run

APP_NAME = stickerfy
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build cover.out

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
	GOARCH=amd64 go build -tags musl -o $(BUILD_DIR)/$(APP_NAME) .

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
		-p 8000:8000 \
		-e ENV="dev" \
		-e SERVER_HOST="0.0.0.0" \
		-e SERVER_PORT="8000" \
		-e READ_TIMEOUT=60 \
		-e WRITE_TIMEOUT=15 \
		-e IDLE_TIMEOUT=60 \
		-e MONGO_USER="mongoadmin" \
		-e MONGO_PASSWORD="secret" \
		-e MONGO_HOST="mongo" \
		-e MONGO_PORT="27017" \
		-e MONGO_DATABASE="stickerfy" \
		-e REDIS_HOST="redis" \
		-e REDIS_PORT="6379" \
		-e REDIS_PASSWORD="secret" \
		-e KAFKA_BROKERS="broker:29092" \
		-e TOPIC_NAME="stickerfy_order_added" \
		stickerfy

docker.redis:
	docker run --rm -d \
		--name redis \
		--network dev-network \
		-p 6379:6379 \
		redis \
		-- requirepass "secret"

docker.mongo:
	docker run --rm -d \
		--name mongo \
		--network dev-network \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
		-e MONGO_INITDB_ROOT_PASSWORD=secret \
		mongo

docker.zookeeper:
	docker run --rm -d \
		--name zookeeper \
		--network dev-network \
		-p 2181:2181 \
    -e ZOOKEEPER_CLIENT_PORT=2181 \
    -e ZOOKEEPER_TICK_TIME=2000 \
		confluentinc/cp-zookeeper:7.3.0

docker.kafka:
	docker run --rm -d \
		--name broker \
		--network dev-network \
		-p 9092:9092 \
		-p 19092:19092 \
    -e KAFKA_BROKER_ID=1 \
    -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
    -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://broker:29092 \
    -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
    -e KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1 \
    -e KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1 \
		confluentinc/cp-kafka:7.3.0

docker.stop: docker.stop.stickerfy docker.stop.redis docker.stop.mongo docker.stop.zookeeper docker.stop.kafka

docker.stop.stickerfy:
	docker stop stickerfy

docker.stop.redis:
	docker stop redis

docker.stop.mongo:
	docker stop mongo

docker.stop.zookeeper:
	docker stop zookeeper

docker.stop.kafka:
	docker stop broker

docker.scan:
	docker scan stickerfy

refresh: docker.stop.stickerfy swag docker.stickerfy
