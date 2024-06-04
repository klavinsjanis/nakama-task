package main

import (
	"context"
	"database/sql"
	"fmt"
	"nakama-task/rpc"

	"github.com/heroiclabs/nakama-common/runtime"
)

// InitModule initializes module.
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("initializing custom RPC functions")

	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS user_requests (
	id SERIAL PRIMARY KEY,
	user_id VARCHAR NOT NULL,
	rpc_name VARCHAR NOT NULL,
	requested_at TIMESTAMP NOT NULL
);`); err != nil {
		return fmt.Errorf("create user_requests db table")
	}

	if err := initializer.RegisterRpc("GetFileContents", rpc.GetFileContents); err != nil {
		return fmt.Errorf("register GetFileContents RPC: %w", err)
	}

	logger.Info("initialization of custom RPC functions complete")

	return nil
}
