package models

import (
	"context"
	"fmt"
	"mpg/htmx_go_poc/webserver/db"
	"time"
)

type Tool struct {
	Name    string `db:"name"`
	Date    time.Time
	DateInt int64  `db:"date"`
	Color   string `db:"color"`
	Size    int    `db:"size"`
}

func ToolInsert(ctx context.Context, t Tool) error {
	t.DateInt = t.Date.UnixMilli()
	tx, err := db.BeginTx(db.BeginTxParams{
		Ctx: ctx,
	})
	if err != nil {
		return fmt.Errorf("on InserTool, BeginTx: %w", err)
	}

	tx.ExecContext(ctx, "INSERT INTO tool ")

	tx.Commit()
	return nil
}

func ToolSelectAll(ctx context.Context) ([]Tool, error) {
	rows, err := db.DB().QueryxContext(ctx, `SELECT * FROM tool`, nil)
	if err != nil {
		return nil, fmt.Errorf("on ToolSelectAll, Query: %w", err)
	}
	result := make([]Tool, 0)
	for rows.Next() {
		var t Tool
		rows.StructScan(&t)
		t.Date = time.Unix(0, t.DateInt*int64(time.Millisecond))
		result = append(result, t)
	}
	return result, nil
}
