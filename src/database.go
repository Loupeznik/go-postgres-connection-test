package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() (int32, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return 0, err
	}
	defer conn.Close(context.Background())

	var n int32
	err = conn.QueryRow(context.Background(), "SELECT 1").Scan(&n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed failed: %v\n", err)
		return 0, err
	}

	return n, nil
}
