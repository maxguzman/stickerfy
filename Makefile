REDIS_VERSION :=7.0.11
MONGO_VERSION := 4.4.6
CONFLUENT_VERSION := 7.3.0
PROMETHEUS_VERSION := v2.44.0
GRAFANA_VERSION := 9.5.2
JAEGER_VERSION := 1.47
OTEL_COLLECTOR_VERSION := 0.83.0

docker.run: docker.dev-dependencies docker.stickerfy-api docker.stickerfy-webapp docker.import-products

docker.dev-dependencies: docker.network docker.redis docker.mongo docker.zookeeper docker.kafka docker.prometheus docker.grafana docker.jaeger docker.otel-collector

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
		-e OTEL_COLLECTOR_OTPL_TRACES_ENDPOINT="grpc://otel-collector:4317" \
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
		-e OTEL_COLLECTOR_ENDPOINT="localhost:4317" \
		stickerfy-api

docker.redis:
	docker run --rm -d \
		--name redis \
		--network dev-network \
		-p 6379:6379 \
		redis:$(REDIS_VERSION) \
		-- requirepass "secret"

docker.mongo:
	docker run --rm -d \
		--name mongo \
		--network dev-network \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
		-e MONGO_INITDB_ROOT_PASSWORD=secret \
		mongo:$(MONGO_VERSION)

docker.zookeeper:
	docker run --rm -d \
		--name zookeeper \
		--network dev-network \
		-p 2181:2181 \
    -e ZOOKEEPER_CLIENT_PORT=2181 \
    -e ZOOKEEPER_TICK_TIME=2000 \
		confluentinc/cp-zookeeper:$(CONFLUENT_VERSION)

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
		confluentinc/cp-kafka:$(CONFLUENT_VERSION)

docker.prometheus:
	docker run --rm -d \
		--name prometheus \
		--network dev-network \
		-p 9090:9090 \
		-v $(shell pwd)/api/config/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus:$(PROMETHEUS_VERSION)

docker.grafana:
	docker run --rm -d \
		--name grafana \
		--network dev-network \
		-p 3000:3000 \
		-v $(shell pwd)/api/config/grafana/datasource.yml:/etc/grafana/provisioning/datasources/default.yml \
		-v $(shell pwd)/api/config/grafana/dashboards.yml:/etc/grafana/provisioning/dashboards/local.yml \
		-v $(shell pwd)/api/config/grafana/dashboard.json:/var/lib/grafana/dashboards/dashboard.json \
		-e "GF_INSTALL_PLUGINS=grafana-clock-panel,grafana-simple-json-datasource" \
    grafana/grafana-oss:$(GRAFANA_VERSION)

docker.jaeger:
	docker run --rm -d \
		--name jaeger \
		--network dev-network \
		-e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
		-e COLLECTOR_OTLP_ENABLED=true \
		-e METRICS_STORAGE_TYPE=prometheus \
		-e PROMETHEUS_SERVER_URL=http://prometheus:9090 \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 5778:5778 \
		-p 16686:16686 \
		-p 14250:14250 \
		-p 14268:14268 \
		-p 14269:14269 \
		-p 9411:9411 \
		jaegertracing/all-in-one:$(JAEGER_VERSION)

docker.otel-collector:
	docker run --rm -d \
    --name otel-collector \
    --network dev-network \
    -p 4317:4317/tcp \
    -p 4318:4318/tcp \
    -v $(PWD)/api/config/otel/otel-config.yaml:/tmp/otel-config.yaml \
    otel/opentelemetry-collector:$(OTEL_COLLECTOR_VERSION) --config /tmp/otel-config.yaml

docker.stop: docker.stop.stickerfy-webapp docker.stop.stickerfy-api docker.stop.dev-dependencies

docker.stop.dev-dependencies: docker.stop.redis docker.stop.mongo docker.stop.zookeeper docker.stop.kafka docker.stop.prometheus docker.stop.grafana docker.stop.jaeger docker.stop.otel-collector

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

docker.stop.jaeger:
	docker stop jaeger

docker.stop.otel-collector:
	docker stop otel-collector

docker.scout:
	docker scout recommendations stickerfy-api && \
	docker scout recommendations stickerfy-webapp

docker.import-products:
	docker run --rm \
		--network dev-network \
		-v $(shell pwd)/static/stickerfy.products.json:/stickerfy.products.json \
		mongo:$(MONGO_VERSION) \
		mongoimport \
		-h mongo:27017 \
		-d stickerfy \
		-c products \
		-u mongoadmin \
		-p secret \
		--authenticationDatabase admin \
		--file /stickerfy.products.json
