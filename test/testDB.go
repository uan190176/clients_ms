package main

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"math/rand"
	"time"
)

type Ttable struct {
	Name  string `db:"name"`
	Value uint64 `db:"value"`
	Id    uint64 `db:"id"`
}

const ISO_LEVEL = pgx.ReadCommitted
const DB_URL = "postgresql://adm:A159753a@192.168.100.25:5432/test?sslmode=disable"

func main() {
	start := time.Now()
	ctx := context.Background()

	err := Ping(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping db is ok")
	}

	//var wg sync.WaitGroup
	//
	//inc := 30
	//wg.Add(inc)
	//cnt := 0
	//for i := 0; i < inc/3; i++ {
	//	go func(i int) {
	//		data, err := readData(ctx)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		cnt++
	//		fmt.Println(i, cnt, data)
	//		wg.Done()
	//	}(i)
	//	go func(i int) {
	//		err := writeData(ctx)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		cnt++
	//		fmt.Println(i, cnt)
	//		wg.Done()
	//	}(i)
	//	go func(i int) {
	//		err := updateData(ctx)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		cnt++
	//		fmt.Println(i, cnt)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %v\n", elapsed)
}

func Ping(ctx context.Context) error {

	config, err := pgxpool.ParseConfig(DB_URL)
	if err != nil {
		return err
	}
	config.MinConns = 1
	config.MaxConns = 1
	config.HealthCheckPeriod = time.Second * 60
	config.MaxConnIdleTime = time.Minute * 5
	config.MaxConnLifetime = time.Minute * 10
	config.MaxConnLifetimeJitter = time.Minute * 10

	connConfig, err := pgx.ParseConfig(DB_URL)
	if err != nil {
		return err
	}
	connConfig.RuntimeParams = map[string]string{
		"application_name":                    "testDB:getRows()",
		"statement_timeout":                   "600s",
		"idle_in_transaction_session_timeout": "600s",
	}
	config.ConnConfig = connConfig

	//conn, err := pgx.ConnectConfig(ctx, config)
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       ISO_LEVEL,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.Deferrable,
	})

	var tRows []*Ttable

	err = tx.QueryRow(ctx, "SELECT * FROM my_table WHERE id = 10").Scan(&tRows)

	if err != nil {
		e := tx.Rollback(ctx)
		if e != nil {
			return e
		}
		return err
	}

	e := tx.Commit(ctx)
	if e != nil {
		return e
	}

	return nil
}

func readData(ctx context.Context) ([]*Ttable, error) {

	conn, err := pgxpool.New(ctx, DB_URL)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	var tRows []*Ttable

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: ISO_LEVEL})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = pgxscan.Select(context.Background(), conn, &tRows, "SELECT * FROM my_table WHERE id = 10")
	if err != nil {
		fmt.Println(err)
		e := tx.Rollback(ctx)
		if e != nil {
			fmt.Println(e)
			return nil, e
		}
		return nil, err
	}
	e := tx.Commit(ctx)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	return tRows, err
}

func writeData(ctx context.Context) error {

	conn, err := pgxpool.New(ctx, DB_URL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: ISO_LEVEL})
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	_, err = conn.Exec(context.Background(), "INSERT INTO my_table (name, value) VALUES ($1, $2)", "test", uint64(rand.Intn(100)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func updateData(ctx context.Context) error {

	conn, err := pgxpool.New(ctx, DB_URL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: ISO_LEVEL})
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = conn.Exec(context.Background(), "UPDATE my_table SET (name, value) = ($1, $2) WHERE id = $3", "test_update", rand.Intn(1000), 11)

	if err != nil {
		fmt.Println(err)
		e := tx.Rollback(ctx)
		if e != nil {
			fmt.Println(e)
			return e
		}
		return err
	}
	e := tx.Commit(ctx)
	if e != nil {
		fmt.Println(e)
		return e
	}
	return nil
}
