package helper

import "testing"

func TestHelloWorldSuccess(t *testing.T) {
	result := HelloWorld("bay")
	if result != "Hello bay" {
		//t.Fail() // masih melanjutkan eksekusi code
		t.Error("harusnya bays") // memanggil t.Fail() dengan message
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("bay")
	if result != "Hello bays" {
		//t.Fail() // masih melanjutkan eksekusi code
		t.Error("harusnya bays") // memanggil t.Fail() dengan message
	}
}

func TestHelloWorldWithString(t *testing.T) {
	result := HelloWorld("bayazid sustami")
	if result != "Hello bayazid sustamis" {
		//t.FailNow() // langsung  menghentikan eksekusi code
		t.Fatal("harusnya bayazid sustamis") // memanggil t.FailNow() dengan message
	}
}
