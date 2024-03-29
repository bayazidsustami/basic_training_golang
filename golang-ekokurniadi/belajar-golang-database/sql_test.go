package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSqlInjectionFixed(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	username := "admin'; #"
	password := "salah"

	sqlScript := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1"
	rows, err := db.QueryContext(context, sqlScript, username, password)
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

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	username := "admin1"
	password := "admin123"

	queryInsert := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(context, queryInsert, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new customer")
}

func TestSqlAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	email := "eko@mail.com"
	comment := "tes comment test"

	queryInsert := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(context, queryInsert, email, comment)
	if err != nil {
		panic(err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new comment with id", lastId)
}

func TestSqlPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	queryInsert := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(context, queryInsert)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@mail.com"
		comment := "ini comment ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(context, email, comment)
		if err != nil {
			panic(err)
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("success insert new comment with id", lastId)
	}
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	queryInsert := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@mail.com"
		comment := "ini comment ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(context, queryInsert, email, comment)
		if err != nil {
			panic(err)
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("success insert new comment with id", lastId)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
