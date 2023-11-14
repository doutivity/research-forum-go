# Creates new migration file with the current timestamp
# Example: make create-new-migration-file NAME=<name>
create-new-migration-file:
	goose -dir .\migrations\ create $(NAME) sql

POSTGRES_URI=postgres://user:secretpassword@localhost:5434/forum-db?sslmode=disable

migrate-up:
	cmd /C goose -dir .\migrations\ -table schema_migrations postgres $(POSTGRES_URI) up
