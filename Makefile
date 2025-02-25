run:
	go run cmd/main.go & envoy -c gateway/envoy.yaml

# Migrate database in postgresql
DB_NAME = shop
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable

force_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations force 1

version_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path migrations down 1

sqlc:
	sqlc generate --file=database/sqlc.yaml
docker-up:
	docker compose up -d
generate-proto:
	cd api && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. order.proto
build-gateway:
	envoy -c gateway/envoy.yaml
.PHONY: run migrate_up sqlc docker-up generate-proto

# Migration