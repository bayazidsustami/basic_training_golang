package app

import (
	"belajar-golang-database-migration/utils"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database_migration")
	utils.PanicErr(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	//migration command
	//migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
	//migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down
	//migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up 2 --migrate to two version
	//migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations version -> get latest version
	//migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20240108151351 -> rollback dirty state
	// migrate create -ext sql -dir db/migrations create_table_category -- create migration file

	return db
}
