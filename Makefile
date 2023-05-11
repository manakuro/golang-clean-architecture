# Set up tools.
install:
	go install github.com/cosmtrek/air@v1.27.3
	go install github.com/pressly/goose/v3/cmd/goose@v3.11.2

# Start dev server.
start:
	air

# Set up database.
setup_db:
	./bin/init_db.sh

# Migrate scheme to database.
migrate_schema_up:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" up

migrate_schema_down:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" down

migrate_schema_reset:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" reset

migrate_schema_status:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" status

seed:
	go run ./cmd/seed/main.go

.PHONY: install setup_db start migrate_schema seed
