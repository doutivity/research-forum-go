package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

const (
	dataSourceName = "postgresql://user:secretpassword@postgres1:5432/forum-db?sslmode=disable"
)

func TestPing(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	connection, err := sql.Open("postgres", dataSourceName)
	require.NoError(t, err)
	defer connection.Close()

	require.NoError(t, connection.Ping())
}
