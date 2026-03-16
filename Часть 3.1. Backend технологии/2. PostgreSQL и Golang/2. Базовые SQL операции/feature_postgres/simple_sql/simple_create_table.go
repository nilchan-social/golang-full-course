package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP,

		UNIQUE(title)
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
