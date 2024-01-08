package main

import (
	"belajar-golang-restful/utils"
)

func main() {

	server := InitializedServer()
	err := server.ListenAndServe()
	utils.PanicErr(err)

}
