package main

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

const (
	dsn = "postgresql://adm:A159753a@192.168.100.25:5432/test?sslmode=disable"
)

type MTable struct {
	Name  string `db:"name"`
	Value uint64 `db:"value"`
	Id    uint64 `db:"id"`
}

func main() {
	ctx := context.Background()

	// Vars for DB
	var (
		err    error
		config *pgx.ConnConfig
		conn   *pgx.Conn
		tx     pgx.Tx
	)

	// Create config to connect
	config, err = pgx.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	// Adding runtime params to config
	config.RuntimeParams = map[string]string{
		"application_name":                    "testDB:getRow(10)",
		"statement_timeout":                   "60s",
		"idle_in_transaction_session_timeout": "60s",
	}

	// Open connection to database
	conn, err = pgx.ConnectConfig(ctx, config)
	if err != nil {
		panic(err)
	}

	// Defer func to close connection
	defer func(conn *pgx.Conn, ctx context.Context) {
		err = conn.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(conn, ctx)

	// Var for result
	var mTable []*MTable

	// Beginning transaction
	tx, err = conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: pgx.NotDeferrable,
	})
	if err != nil {
		fmt.Println(err)
	}

	// Running query
	err = pgxscan.Select(ctx, tx, &mTable, "SELECT name, value, id FROM my_table WHERE id = $1", 10)
	if err != nil {
		fmt.Println(err)
		// Rollback transaction
		e := tx.Rollback(ctx)
		if e != nil {
			fmt.Println(e)
		}
	} else {
		// Commit transaction
		e := tx.Commit(ctx)
		if e != nil {
			fmt.Println(e)
		}
	}

	// Returning result
	fmt.Printf("name: %s, value: %d\n", mTable[0].Name, mTable[0].Value)
}
