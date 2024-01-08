package app

import (
	"belajar-golang-restful/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(autMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: autMiddleware,
		}
}
