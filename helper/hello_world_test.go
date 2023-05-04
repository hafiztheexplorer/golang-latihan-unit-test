package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//cara untuk melakukan unit testing dengan menggunakan slice/tabel sehingga tidak perlu lagi megngetest satu2
func TestTabelTesting(t *testing.T)  {

	Headertabeltests := []struct {
	name string
	request string
	expected string
	}{
	{
		name: "nama1",
		request: "nama1",
		expected: "Halo nama1",
	},
	{
		name: "nama2",
		request: "nama2",
		expected: "Halo nama2",
	},
	{
		name: "nama3",
		request: "nama3",
		expected: "Halo nama3",
	},
}

for _, row := range Headertabeltests{
	t.Run(row.name,func(t *testing.T)   {
		result:= HelloWorld(row.name)
		require.Equal(t, row.expected, result, "Result harusnya ",row.request)
	})
}
	
}

//functiom untuk mengetest function dengan subunittest
func TestSubTest(t *testing.T)()  {
	t.Run("Master Name 1",func(t *testing.T)   {
		result:= HelloWorld("Master Name 1")
		require.Equal(t, "Halo Master Name 1", result, "Result harusnya 'Halo Master Name 1'")
	})
	t.Run("Master Name 2",func(t *testing.T)  {
		result:= HelloWorld("Master Name 2")
		require.Equal(t, "Halo Master Name 2", result, "Result harusnya 'Halo Master Name 2'")
	})
}

//function untuk mengetest unit test, semua unit test dan dijalankan semua dari atas ke bawah, sekali
func TestMain(m *testing.M)()  {
	//before
	fmt.Println("BEFORE UNIT TEST")
	
	
	m.Run() //eksekusi semua unit test

	// after
	fmt.Println("AFTER UNIT TEST")
}

//function unit test agar mengskip pengetesan
func TestSkipHelloWorld(t *testing.T)  {
	if runtime.GOOS == "windows"{
		t.Skip("cannot run on windows")
	}
	result:= HelloWorld("Master Name Require")
	require.Equal(t, "Halo Master Name Require", result, "Result harusnya 'Halo Master Name Require'")
}

//function unit test custom menggunakan library assertion
func TestHelloWorldAssertion(t *testing.T)   {
	fmt.Println("-----------------------beginning of Assertion function----")
result:= HelloWorld("Master Name Assertion")
assert.Equal(t, "Halo Master Name Assertion", result, "Result harusnya 'Halo Master Name Assertion'")
	fmt.Println("-----------------------end of Assertion function----")
}

//function unit test custom menggunakan library require
func TestHelloWorldRequire(t *testing.T) {
	fmt.Println("-----------------------beginning of Require function----")
	result:= HelloWorld("Master Name Require")
	require.Equal(t, "Halo Master Name Require", result, "Result harusnya 'Halo Master Name Require'")
	fmt.Println("-----------------------end of Require function----")
}

// func TestHelloWorld1(t *testing.T) { //parameter akan dimasukkan ke package hello world
// 	fmt.Println("-----------------------beginning of t.Fail() function----")
// 	result:=HelloWorld("Master Name")
// 	if result!= "Hallo Master Name"{
// 		// saat error
// 		t.Fail()
// 	}
// 	fmt.Println("-----------------------end of t.Fail() function----")
// }

// func TestHelloWorld2(t *testing.T) {
// 	fmt.Println("-----------------------beginning of t.FailNow() function----")
// 	result:=HelloWorld("Master Name 2")
// 	if result!= "Hallo Master Name 2"{
// 		// saat error
// 		t.FailNow()
// 	}
// 	fmt.Println("-----------------------end of t.FailNow() function----")
// }

func TestHelloWorld3(t *testing.T) {
	fmt.Println("-----------------------beginning of t.Error() function----")
	result:=HelloWorld("Master Name 3")
	if result!= "Hallo Master Name 3"{
		// saat error
		t.Error("result must be 'Halo Master Name 3'")
	}
	fmt.Println("-----------------------end of t.Error() function----")
}

func TestHelloWorld4(t *testing.T) {
	fmt.Println("-----------------------beginning of t.Fatal() function----")
	result:=HelloWorld("Master Name 4")
	if result!= "Hallo Master Name 4"{
		// saat error
		t.Fatal("Result must be 'Hallo Master Name 4'")
	}
	fmt.Println("-----------------------end of t.Fatal() function----")
}