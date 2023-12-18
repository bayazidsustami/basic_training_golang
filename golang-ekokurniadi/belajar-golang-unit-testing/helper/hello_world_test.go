package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("run before")

	m.Run() // eksekusi semua code

	fmt.Println("run after")
}

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

func TestHelloName(t *testing.T) {
	t.Run("Bay", func(t *testing.T) {
		result := HelloWorld("Bay")
		assert.Equal(t, "Hello Bay", result)
	})
	t.Run("Yazid", func(t *testing.T) {
		result := HelloWorld("Yazid")
		assert.Equal(t, "Hello Yazid", result)
	})
	t.Run("Bayazid", func(t *testing.T) {
		result := HelloWorld("Bayazid")
		assert.Equal(t, "Hello Bayazid", result)
	})
}

func TestHelloNameWithTableTest(t *testing.T) {
	testCases := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Bay",
			Request:  "Bay",
			Expected: "Hello Bay",
		},
		{
			Name:     "Bayazid",
			Request:  "Bayazid",
			Expected: "Hello Bayazid",
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			assert.Equal(t, test.Expected, result)
		})
	}
}
