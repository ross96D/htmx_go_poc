package models

import (
	"context"
	"database/sql"
	"mpg/htmx_go_poc/webserver/db"
	"time"

	"github.com/pkg/errors"
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
		return errors.WithStack(err)
	}

	_, err = tx.ExecContext(
		ctx, "INSERT INTO tool (name, color, \"size\", date) VALUES($1, $2, $3, $4)",
		t.Name, t.Color, t.Size, t.DateInt,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ToolSelectAll(ctx context.Context) ([]Tool, error) {
	rows, err := db.DB().QueryxContext(ctx, `SELECT * FROM tool`, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	result := make([]Tool, 0)
	for rows.Next() {
		var t Tool
		rows.StructScan(&t)
		t.Date = sql.NullTime{
			Time:  time.Unix(0, t.DateInt.Int64*int64(time.Millisecond)),
			Valid: t.DateInt.Valid,
		}
		result = append(result, t)
	}
	return result, nil
}

func ToolDelete(ctx context.Context, id int) error {
	_, err := db.DB().ExecContext(ctx,
		`DELETE FROM tool WHERE id=$1`, id,
	)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
