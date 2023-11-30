package schema

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func MigrateUp(db *sql.DB) error {
	// goose has a lot of dependencies with ClickHouse and other DB drivers
	goose.SetBaseFS(migrations)
	goose.SetTableName("schema_migrations")
	// PostgreSQL by default
	// goose.SetDialect("postgres")

	return goose.Up(db, "migrations")
}

func MigrateDown(db *sql.DB) error {
	// goose has a lot of dependencies with ClickHouse and other DB drivers
	goose.SetBaseFS(migrations)
	goose.SetTableName("schema_migrations")
	// PostgreSQL by default
	// goose.SetDialect("postgres")
	return goose.DownTo(db, "migrations", 0)
}
