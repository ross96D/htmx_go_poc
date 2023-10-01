package models

import (
	"context"
	"database/sql"
	"fmt"
	"mpg/htmx_go_poc/webserver/db"
	"time"
)

type Tool struct {
	ID      int           `db:"id"`
	Name    string        `db:"name"`
	DateInt sql.NullInt64 `db:"date"`
	Date    sql.NullTime
	Color   string        `db:"color"`
	Size    sql.NullInt32 `db:"size"`
}

func ToolInsert(ctx context.Context, t Tool) error {
	t.DateInt = sql.NullInt64{
		Int64: t.Date.Time.UnixMilli(),
		Valid: t.Date.Valid,
	}
	tx, err := db.BeginTx(db.BeginTxParams{
		Ctx: ctx,
	})
	if err != nil {
		return fmt.Errorf("on InserTool, BeginTx: %w", err)
	}

	_, err = tx.ExecContext(
		ctx, "INSERT INTO tool (name, color, \"size\", date) VALUES($1, $2, $3, $4)",
		t.Name, t.Color, t.Size, t.DateInt,
	)
	if err != nil {
		return fmt.Errorf("on InserTool, ExecContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("on InserTool, Commit: %w", err)
	}
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
		// t.Date = time.Unix(0, t.DateInt*int64(time.Millisecond))
		println(t.DateInt.Int64)
		println(t.DateInt.Valid)
		t.Date = sql.NullTime{
			Time:  time.Unix(0, t.DateInt.Int64*int64(time.Millisecond)),
			Valid: t.DateInt.Valid,
		}
		result = append(result, t)
	}
	return result, nil
}
