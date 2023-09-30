package models

import (
	"context"
	"fmt"
	"mpg/htmx_go_poc/webserver/db"
	"time"
)

type Tool struct {
	Name  string
	Date  time.Time
	Color string
	Size  int
}

func InsertTool(ctx context.Context, t Tool) error {
	tx, err := db.BeginTx(db.BeginTxParams{
		Ctx: ctx,
	})
	if err != nil {
		return fmt.Errorf("on InserTool, BeginTx: %w", err)
	}
	defer tx.Commit()

	return nil
}
