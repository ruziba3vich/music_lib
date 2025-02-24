include .env
export $(shell sed 's/=.*//' .env)

name ?= new_migration
dir = migrations

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

migrate-create:
	docker run --rm -v $(PWD)/migrations:/migrations migrate/migrate \
		create -ext sql -dir $(dir) -seq $(name)

migrate-up:
	docker run --rm --network=bridge -v $(PWD)/migrations:/migrations migrate/migrate \
		-path=/migrations -database $(DB_URL) up

migrate-down:
	docker run --rm --network=bridge -v $(PWD)/migrations:/migrations migrate/migrate \
		-path=/migrations -database $(DB_URL) down 1

migrate-down-all:
	docker run --rm --network=bridge -v $(PWD)/migrations:/migrations migrate/migrate \
		-path=/migrations -database $(DB_URL) down

# Rollback all migrations
migrate-down-all:
	docker run --rm --network=host -v $(PWD)/$(MIGRATION_DIR):/migrations migrate/migrate \
	  -path=/migrations -database "$(DB_URL)" down

# Run the application
run:
	go run cmd/main.go

# Build the application
build:
	go build -o music_lib main.go

swag-gen:
	swag init -g internal/http/handler.go -o docs --parseDependency --parseInternal
test:
	go test ./...

# Clean up binary files
clean:
	rm -f music_lib
