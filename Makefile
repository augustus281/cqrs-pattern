run:
	go run cmd/main.go
migrate_up:
	goose -dir database/migrations postgres "postgres://root:mysecurepassword@localhost:5432/shop?sslmode=disable" up
sqlc:
	sqlc generate --file=database/sqlc.yaml
docker-up:
	docker compose up -d
generate-proto:
	cd api && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. order.proto
.PHONY: run migrate_up sqlc docker-up generate-proto