# Set up tools.
install:
	go install github.com/cosmtrek/air@v1.27.3

# Start dev server.
start:
	air

# Set up database.
setup_db:
	./bin/init_db.sh

# Migrate scheme to database.
migrate_schema:
	go run ./cmd/migration/main.go

.PHONY: install setup_db start migrate_schema
