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
