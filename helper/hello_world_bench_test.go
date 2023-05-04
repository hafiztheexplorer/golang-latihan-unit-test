package helper

import (
	"testing"
)

// code program untuk benchmark hello world
func BenchmarkHelloWorld1(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		benchinput:="name 1"
		HelloWorld(benchinput)
	}
	
}
func BenchmarkHelloWorld2(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		// Name:="name 1"
		HelloWorld("name 1")
	}
	
}


//code program untuk sub benchmark

func BenchmarkHelloWorldSub(b *testing.B)  {
	benchsubinput1:="name 1"
	benchsubinput2:="name 2"
	b.Run("name 1",func(b *testing.B)   {
		for i := 0; i < b.N; i++ {
		
			HelloWorld(benchsubinput1)
		}
	})

	b.Run("name 2",func(b *testing.B)   {
		for i := 0; i < b.N; i++ {

			HelloWorld(benchsubinput2)
		}
	})
}

//code program untuk tabel benchmark

func BenchmarkHelloWorldTabel(b *testing.B)  {
	Headertabelbenchmark := []struct {
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
	
	for _, row := range Headertabelbenchmark{
		b.Run(row.name,func(b *testing.B)   {
			for i := 0; i < b.N; i++{
				//simple output
				HelloWorld(row.name)
				//agak rumit output
				// result:= HelloWorld(row.name)
				// require.Equal(b, row.expected, result, "Result harusnya ",row.request)
			}
			
		})
	}
}