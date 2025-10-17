.SILENT:

# env var
include .env
export

# Build binary
build:
	go mod download && go build -o ./bin/app ./cmd/app/main.go

# Run server
run-server: build
	./bin/app

# Run docker-compose
run: build
	docker-compose up --remove-orphans --build -d

# Create new migration (make create-migration NAME=name)
create-migration:
	migrate create -ext sql -dir database/migrations -seq $(NAME)

# Apply all migrations
migrate-up-all:
	migrate -path ./database/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" up

# Apply migration
migrate-up:
	migrate -path ./database/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" up 1

# Rollback last migration
migrate-down:
	migrate -path ./database/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" down 1

# Show migration status
migrate-status:
	migrate -path ./database/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" version

# Rollback all migrations
migrate-reset:
	migrate -path ./database/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" down

# seed build
seed-build:
	go build -o ./bin/seed ./seed/seed.go

# 
run-seed: seed-build
	./bin/seed

# Stop containers
down:
	docker-compose down

setup: run migrate-up-all
	echo "Setup complete: containers running and migrations applied"