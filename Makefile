POSTGRES_URI=postgresql://user:secretpassword@localhost:5432/forum-db?sslmode=disable

env-up:
	docker-compose up -d

env-down:
	docker-compose down --remove-orphans -v

go-test:
	docker exec research-forum-go-app go test ./... -v -count=1

docker-go-version:
	docker exec research-forum-go-app go version
	docker exec research-forum-go-app go run main.go

docker-pg-version:
	docker exec research-forum-go-postgres-1 psql -U user -d forum-db -c "SELECT VERSION();"

# Example: make test WAIT_POSTGRES_LAUNCH=15s
test:
	$(eval WAIT_POSTGRES_LAUNCH ?= 0s)
	make env-up
	sleep $(WAIT_POSTGRES_LAUNCH)
	make docker-go-version
	make docker-pg-version
	make go-test
	make env-down

# go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# sqlc generate
#
# alternative
# docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate
generate-sqlc:
	sqlc generate

# Creates new migration file with the current timestamp
# Example: make create-new-migration-file NAME=<name>
create-new-migration-file:
	$(eval NAME ?= noname)
	mkdir -p ./schema/migrations/
	goose -dir ./schema/migrations/ create $(NAME) sql

migrate-up:
	goose -dir ./schema/migrations/ -table schema_migrations postgres $(POSTGRES_URI) up
migrate-redo:
	goose -dir ./schema/migrations/ -table schema_migrations postgres $(POSTGRES_URI) redo
migrate-down:
	goose -dir ./schema/migrations/ -table schema_migrations postgres $(POSTGRES_URI) down
migrate-reset:
	goose -dir ./schema/migrations/ -table schema_migrations postgres $(POSTGRES_URI) reset
migrate-status:
	goose -dir ./schema/migrations/ -table schema_migrations postgres $(POSTGRES_URI) status

install-sqlc:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest
