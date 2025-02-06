run:
	go run cmd/main.go
migrate_up:
	goose -dir database/migrations postgres "postgres://root:mysecurepassword@localhost:5432/shop?sslmode=disable" up
sqlc:
	sqlc generate --file=database/sqlc.yaml
docker_up:
	docker compose up -d
.PHONY: run migrate_up sqlc docker_up