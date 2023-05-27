docker.run: docker.dev-dependencies docker.stickerfy-api docker.stickerfy-webapp docker.import-products

docker.dev-dependencies: docker.network docker.redis docker.mongo docker.zookeeper docker.kafka docker.prometheus docker.grafana

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.stickerfy-webapp.build:
	docker build -t stickerfy-webapp ./webapp

docker.stickerfy-webapp: docker.stickerfy-webapp.build
	docker run --rm -d \
		--name stickerfy-webapp \
		--network dev-network \
		-p 8080:8080 \
		-e STICKERFY_SERVICE_URL="http://stickerfy-api:8000/v1" \
		stickerfy-webapp

docker.stickerfy-api.build:
	docker build -t stickerfy-api ./api

docker.stickerfy-api: docker.stickerfy-api.build
	docker run --rm -d \
		--name stickerfy-api \
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
		stickerfy-api

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
		mongo:4.4.16

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

docker.prometheus:
	docker run --rm -d \
		--name prometheus \
		--network dev-network \
		-p 9090:9090 \
		-v $(shell pwd)/api/config/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus

docker.grafana:
	docker run --rm -d \
		--name grafana \
		--network dev-network \
		-p 3000:3000 \
		-v $(shell pwd)/api/config/grafana/datasource.yml:/etc/grafana/provisioning/datasources/default.yml \
		-v $(shell pwd)/api/config/grafana/dashboards.yml:/etc/grafana/provisioning/dashboards/local.yml \
		-v $(shell pwd)/api/config/grafana/dashboard.json:/var/lib/grafana/dashboards/dashboard.json \
		-e "GF_INSTALL_PLUGINS=grafana-clock-panel,grafana-simple-json-datasource" \
    grafana/grafana-oss

docker.stop: docker.stop.stickerfy-webapp docker.stop.stickerfy-api docker.stop.dev-dependencies

docker.stop.dev-dependencies: docker.stop.redis docker.stop.mongo docker.stop.zookeeper docker.stop.kafka docker.stop.prometheus docker.stop.grafana

docker.stop.stickerfy-webapp:
	docker stop stickerfy-webapp

docker.stop.stickerfy-api:
	docker stop stickerfy-api

docker.stop.redis:
	docker stop redis

docker.stop.mongo:
	docker stop mongo

docker.stop.zookeeper:
	docker stop zookeeper

docker.stop.kafka:
	docker stop broker

docker.stop.prometheus:
	docker stop prometheus

docker.stop.grafana:
	docker stop grafana

docker.scout:
	docker scout recommendations stickerfy-api && \
	docker scout recommendations stickerfy-webapp

docker.import-products:
	docker run --rm \
		--network dev-network \
		-v $(shell pwd)/static/stickerfy.products.json:/stickerfy.products.json \
		mongo \
		mongoimport \
		-h mongo:27017 \
		-d stickerfy \
		-c products \
		-u mongoadmin \
		-p secret \
		--authenticationDatabase admin \
		--file /stickerfy.products.json
