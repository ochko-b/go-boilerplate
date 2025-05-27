.PHONY: build run test clean docker-up docker-down migrate-up migrate-down sqlc-generate

build:
	go build -o bin/main cmd/server/main.go

run: 
	go run cmd/server/main.go

test: go test -v ./...

clean:
	rm -rf bin/

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down -v

docker-logs:
	docker-compose logs -f

migrate-up: 
	migrate -path internal/database/migrations -database "postgres://postgres:password@localhost:5432/goappdb?sslmode=disable" up

migrate-down:
	migrate -path internal/database/migrations -database "postgres://postgres:password@localhost:5432/goappdb?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir internal/database/migrations -seq $(name)

sqlc-generate:
	sqlc generate

dev-setup: docker-up migrate-up sqlc-generate
	@echo "Development environment ready!"

install-tools:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
