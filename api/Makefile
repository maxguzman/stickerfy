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

docker.dev-dependencies:
	$(MAKE) -C ../ docker.dev-dependencies

dev:
	@ENV="dev" \
	SERVER_HOST="0.0.0.0" \
	SERVER_PORT="8000" \
	READ_TIMEOUT=15 \
	WRITE_TIMEOUT=15 \
	IDLE_TIMEOUT=60 \
	MONGO_USER="mongoadmin" \
	MONGO_PASSWORD="secret" \
	MONGO_HOST="0.0.0.0" \
	MONGO_PORT="27017" \
	MONGO_DATABASE="stickerfy" \
	REDIS_HOST="0.0.0.0" \
	REDIS_PORT="6379" \
	REDIS_PASSWORD="secret" \
	KAFKA_BROKERS="localhost:9092" \
	TOPIC_NAME="stickerfy_order_added" \
	go run main.go
