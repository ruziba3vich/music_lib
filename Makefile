# Load environment variables from .env
export $(shell sed 's/=.*//' .env)

# Migration directory
MIGRATION_DIR = migrations

# Database URL
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Migration name (default: "new_migration")
NAME ?= new_migration

# Create a new migration file
migrate-create:
	docker run --rm -v $(PWD)/$(MIGRATION_DIR):/migrations migrate/migrate \
	  create -ext sql -dir /migrations -seq $(NAME)

# Apply all migrations
migrate-up:
	docker run --rm --network=host -v $(PWD)/$(MIGRATION_DIR):/migrations migrate/migrate \
	  -path=/migrations -database "$(DB_URL)" up

# Rollback the last migration
migrate-down:
	docker run --rm --network=host -v $(PWD)/$(MIGRATION_DIR):/migrations migrate/migrate \
	  -path=/migrations -database "$(DB_URL)" down 1

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

test:
	go test ./...

# Clean up binary files
clean:
	rm -f music_lib
