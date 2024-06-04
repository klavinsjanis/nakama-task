package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// SaveUserRequestDetails saves request details to the database.
// Parameters:
//   - ctx: context.Context
//   - db: *sql.DB
//   - userID: string
//   - rpcName: strings
//   - requestedAt: time.Time
//
// Returns:
//   - error: any error that occurred
func SaveUserRequestDetails(ctx context.Context, db *sql.DB, userID, rpcName string, requestedAt time.Time) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO user_requests (user_id, rpc_name, requested_at)
	VALUES ($1, $2, $3) RETURNING id`)

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("preapre statement: %w", err)
	}
	defer stmt.Close()

	var id int64
	if err = stmt.QueryRow(userID, rpcName, requestedAt).Scan(&id); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return id, nil
}
