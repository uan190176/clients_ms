package main

import (
	"clients_ms/internal/api"
	"context"
	"fmt"
	fs "github.com/uan190176/fstorage"
)

const (
	dsn2 = "postgresql://adm:A159753a@192.168.100.25:5432/test?sslmode=disable"
)

func main() {
	//selectCountries()
	//insertCountries()
	//updateCountries()
	copyCountries()
}

func selectCountries() {
	var countries []*api.Country

	q := "SELECT * FROM countries_full"

	qp := &fs.QueryParams{
		Ctx:         context.Background(),
		Dsn:         dsn2,
		AppName:     "test",
		Dest:        &countries,
		QueryString: q,
		Args:        []interface{}{},
	}

	st, err := fs.NewStorage(qp)
	if err != nil {
		fmt.Println(err)
	}

	err = st.RunQuery(qp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(countries)

	err = st.Release(qp)
	if err != nil {
		fmt.Println(err)
	}
}

func insertCountries() {
	var countries []*api.Country

	//(name, comment, created_by)
	var q = `INSERT INTO countries (name, comment, created_by) VALUES ($1, $2, $3) RETURNING
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

	pName := "TEST2-ins"
	pComm := "test- del"
	pAuthor := 7

	qp := &fs.QueryParams{
		Ctx:         context.Background(),
		Dsn:         dsn2,
		AppName:     "test",
		Dest:        &countries,
		QueryString: q,
		Args:        []interface{}{pName, pComm, pAuthor},
	}

	st, err := fs.NewStorage(qp)
	if err != nil {
		fmt.Println(err)
	}

	err = st.RunQuery(qp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(countries)

	err = st.Release(qp)
	if err != nil {
		fmt.Println(err)
	}
}

func updateCountries() {
	var countries []*api.Country

	//(name, comment, isdeleted, changed_by, id)
	var q = `UPDATE countries SET (name, comment, isdeleted, changed_by) = ($1, $2, $3, $4) WHERE id = $5 RETURNING
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

	pName := "TEST999-updates"
	pComm := "test- del"
	pDel := true
	pAuthor := 7
	pId := 25

	qp := &fs.QueryParams{
		Ctx:         context.Background(),
		Dsn:         dsn2,
		AppName:     "test",
		Dest:        &countries,
		QueryString: q,
		Args:        []interface{}{pName, pComm, pDel, pAuthor, pId},
	}

	st, err := fs.NewStorage(qp)
	if err != nil {
		fmt.Println(err)
	}

	err = st.RunQuery(qp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(countries)

	err = st.Release(qp)
	if err != nil {
		fmt.Println(err)
	}
}

func copyCountries() {

	var countries []*api.Country
	var cc []*api.RequestCountry

	c := &api.RequestCountry{
		CountryName:    "TEST123",
		CountryComment: "TEST123",
		AuthorId:       7,
	}
	cc = append(cc, c)
	c = &api.RequestCountry{
		CountryName:    "TEST321",
		CountryComment: "TEST321",
		AuthorId:       7,
	}
	cc = append(cc, c)

	data1 := [][]interface{}{
		{"data1", 100},
		{"data2", 200},
	}

	qp := &fs.QueryParams{
		Ctx:     context.Background(),
		Dsn:     dsn2,
		AppName: "test",
		Dest:    &countries,
		TabName: "my_table",
		Fields:  []string{"name", "value"},
		Source:  data1,
	}

	st, err := fs.NewStorage(qp)
	if err != nil {
		fmt.Println(err)
	}

	err = st.CopyToDatabase(qp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(countries)

	err = st.Release(qp)
	if err != nil {
		fmt.Println(err)
	}
}
