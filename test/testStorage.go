package main

import (
	"clients_ms/internal/api"
	"context"
	"fmt"
	s "github.com/uan190176/storage"
)

const (
	dsn1 = "postgresql://adm:A159753a@192.168.100.25:5432/test?sslmode=disable"
)

type MTable1 struct {
	Name  string `db:"name"`
	Value uint64 `db:"value"`
	Id    uint64 `db:"id"`
}

func main() {
	ctx := context.Background()

	var err error
	var qParam *s.QueryParams
	var countries []*api.Country

	//(name, comment, created_by)
	var insertText = `INSERT INTO countries (name, comment, created_by) VALUES ($1, $2, $3) RETURNING
					id AS country_id,
    				name AS country_name,
    				concat(comment, '') AS country_comment,
    				to_char(created, 'YYYY-MM-DD HH24:MI:SS'::text) AS country_created,
    				created_by AS country_created_by,
					( SELECT u.fullname
           				FROM users u
          				WHERE created_by = u.id) AS country_created_by_name,
    				to_char(changed, 'YYYY-MM-DD HH24:MI:SS'::text) AS country_changed,
    				changed_by AS country_changed_by,
					( SELECT u.fullname
           				FROM users u
          				WHERE created_by = u.id) AS country_changed_by_name,
    				isdeleted AS country_isdeleted`

	_ = insertText

	//(name, comment, isdeleted, changed_by, id)
	var updateText = `UPDATE countries SET (name, comment, isdeleted, changed_by) = ($1, $2, $3, $4) WHERE id = $5 RETURNING
					id AS country_id,
    				name AS country_name,
    				concat(comment, '') AS country_comment,
    				to_char(created, 'YYYY-MM-DD HH24:MI:SS'::text) AS country_created,
    				created_by AS country_created_by,
					( SELECT u.fullname
           				FROM users u
          				WHERE created_by = u.id) AS country_created_by_name,
    				to_char(changed, 'YYYY-MM-DD HH24:MI:SS'::text) AS country_changed,
    				changed_by AS country_changed_by,
					( SELECT u.fullname
           				FROM users u
          				WHERE created_by = u.id) AS country_changed_by_name,
    				isdeleted AS country_isdeleted`

	//Insert
	qParam = &s.QueryParams{
		CTX:     ctx,
		DSN:     dsn1,
		AppName: "test",
		Dest:    &countries,
		QText:   updateText,
	}
	pName := "TEST1-updates"
	pComm := "test- del"
	pDel := true
	pAuthor := 7
	pId := 25

	err = s.RunQuerySelect(qParam, pName, pComm, pDel, pAuthor, pId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(countries)
}
