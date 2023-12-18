package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestMultiplicationAssertion(t *testing.T) {
	result := Multiplication(1, 2)
	assert.Equal(t, 2, result) // mirip t.Fail()
}

func TestMultiplicationRequire(t *testing.T) {
	result := Multiplication(1, 2)
	require.Equal(t, 2, result) // mirip t.FailNow()
}

func TestMultiplicationSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("can't run on mac")
	}

	result := Multiplication(1, 2)
	assert.Equal(t, 2, result)
}
