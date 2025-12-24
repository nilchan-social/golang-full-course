package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, tasksIDs []int) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, tasksIDs)

	return err
}
