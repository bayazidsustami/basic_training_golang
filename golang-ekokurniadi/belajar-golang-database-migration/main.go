package main

import "belajar-golang-database-migration/utils"

func main() {

	server := InitializedServer()
	err := server.ListenAndServe()
	utils.PanicErr(err)

}
