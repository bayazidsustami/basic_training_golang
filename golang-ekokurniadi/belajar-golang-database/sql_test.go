package belajargolangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	queryInsert := "INSERT INTO customer(id, name) VALUES('baya', 'bayazid sustami')"
	_, err := db.ExecContext(context, queryInsert)
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	querySelect := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(context, querySelect)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id", id)
		fmt.Println("name", name)
	}
}
