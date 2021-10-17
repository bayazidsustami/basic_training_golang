package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"os"
)

func main() {
	//gzip handler
	/*mux := new(http.ServeMux)

	mux.HandleFunc("/image", func(rw http.ResponseWriter, r *http.Request) {
		f, err := os.Open("sample.jpg")
		if f != nil {
			defer f.Close()
		}

		if err != nil{
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = io.Copy(rw, f)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = gziphandler.GzipHandler(mux)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/

	e := echo.New()
	e.Use(middleware.Gzip())

	e.GET("/image", func(ctx echo.Context) error {
		f, err := os.Open("sample.jpg")
		if err != nil {
			return err
		}

		_, err = io.Copy(ctx.Response(), f)
		if err != nil {
			return err
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))

}
