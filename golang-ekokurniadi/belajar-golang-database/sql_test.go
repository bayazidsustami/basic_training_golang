package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	querySelect := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(context, querySelect)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString // support nullable type
		var balance int32
		var rating float64
		var birthDate sql.NullTime // support nullable type
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id, " Name:", name, " Email:", email, " Balance", balance, " Rating:", rating, " Birth Date:", birthDate, " Married:", married, " Created At:", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	username := "admin'; #"
	password := "salah"

	sqlScript := "SELECT username FROM user WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(context, sqlScript)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("success login", username)
	} else {
		fmt.Println("failed login", username)
	}
}
