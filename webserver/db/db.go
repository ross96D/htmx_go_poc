package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func DB() *sqlx.DB {
	return db
}

func Connect() {
	path := "proof.db"

	var err error
	db, err = sqlx.Open(
		"sqlite3",
		path,
	)
	if err != nil {
		log.Fatal(fmt.Sprint("Error connectando con sqlite", err))
	}
}

type BeginTxParams struct {
	Ctx  context.Context
	Opts *sql.TxOptions
}

func BeginTx(params BeginTxParams) (*sqlx.Tx, error) {
	if params.Opts == nil {
		params.Opts = &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  true,
		}
	}
	return db.BeginTxx(params.Ctx, params.Opts)
}
