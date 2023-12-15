package database

var connection string

// auto init ketika pertama kali package di akses
func init() {
	connection = "MYSQL"
}

func GetDatabase() string {
	return connection
}
