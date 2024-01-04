package main

import (
	"belajar-golang-restful/middleware"
	"belajar-golang-restful/utils"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(autMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: autMiddleware,
	}
}

func main() {

	server := InitializedServer()
	err := server.ListenAndServe()
	utils.PanicErr(err)

}
