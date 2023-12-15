package helper

var version = "1.0.0"                  // tidak bisa diakses karena huruf kecil (private)
var ApplicationName = "simple-package" // bisa diakses karena huruf besar (public)

// (public)
func SayHello(name string) string {
	return "hellow " + name
}

// (private)
func sayGoodBye(name string) string {
	return "bye " + name
}
