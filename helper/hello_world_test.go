package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// func BenchmarkTable(b *testing.B) {
// 	benchmarks := []struct {
// 		name    string
// 		request string
// 	}{
// 		{
// 			name:    "Aidil",
// 			request: "Aidil",
// 		},
// 		{
// 			name:    "Adam",
// 			request: "Adam",
// 		},
// 		{
// 			name:    "AidilAdamBaik",
// 			request: "Aidil Adam Baik",
// 		},
// 		{
// 			name:    "Budi",
// 			request: "Budi Nugraha",
// 		},
// 	}

// 	for _, benchmark := range benchmarks {
// 		b.Run(benchmark.name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				HelloWorld(benchmark.request)
// 			}
// 		})
// 	}
// }

// func BenchmarkSub(b *testing.B) {
// 	b.Run("Aidil", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			HelloWorld("Aidil")
// 		}
// 	})
// 	b.Run("Adam", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			HelloWorld("Adam")
// 		}
// 	})
// }

// func BenchmarkHelloWorld(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		HelloWorld("Aidil")
// 	}
// }

// func BenchmarkHelloWorldAdam(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		HelloWorld("Adam")
// 	}
// }

// TestTableHelloWorld adalah fungsi untuk menguji fungsi HelloWorld menggunakan table test
// Table test memungkinkan kita menguji beberapa test case sekaligus dalam satu fungsi
func TestTableHelloWorld(t *testing.T) {
	// Mendefinisikan slice struct yang berisi test case
	tests := []struct {
		name     string  // Nama test case
		request  string  // Parameter input yang akan diuji
		expected string  // Hasil yang diharapkan
	}{
		{
			name:     "Aidil",
			request:  "Aidil", 
			expected: "Hello Aidil", // Mengharapkan output "Hello Aidil"
		},
		{
			name:     "Adam",
			request:  "Adam",
			expected: "Hello Adam", // Mengharapkan output "Hello Adam" 
		},
		{
			name:     "Baik",
			request:  "Baik",
			expected: "Hello Baik", // Mengharapkan output "Hello Baik"
		},
		{
			name:     "Budi", 
			request:  "Budi",
			expected: "Hello Budi", // Mengharapkan output "Hello Budi"
		},
		{
			name:     "Joko",
			request:  "Joko", 
			expected: "Hello Joko", // Mengharapkan output "Hello Joko"
		},
	}

	// Iterasi untuk setiap test case
	for _, test := range tests {
		// Menjalankan sub-test untuk setiap test case
		t.Run(test.name, func(t *testing.T) {
			// Memanggil fungsi HelloWorld dengan parameter dari test case
			result := HelloWorld(test.request)
			// Memastikan hasil sesuai dengan yang diharapkan menggunakan require.Equal
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	// Menjalankan sub-test untuk nama "Aidil"
	t.Run("Aidil", func(t *testing.T) {
		// Memanggil fungsi HelloWorld dengan parameter "Aidil"
		result := HelloWorld("Aidil")
		// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
		// yaitu "Hello Aidil" menggunakan require.Equal
		require.Equal(t, "Hello Aidil", result, "Result must be 'Hello Aidil'")
	})
	
	// Menjalankan sub-test untuk nama "Adam" 
	t.Run("Adam", func(t *testing.T) {
		// Memanggil fungsi HelloWorld dengan parameter "Adam"
		result := HelloWorld("Adam")
		// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
		// yaitu "Hello Adam" menggunakan require.Equal
		require.Equal(t, "Hello Adam", result, "Result must be 'Hello Adam'")
	})
	
	// Menjalankan sub-test untuk nama "Baik"
	t.Run("Baik", func(t *testing.T) {
		// Memanggil fungsi HelloWorld dengan parameter "Baik" 
		result := HelloWorld("Baik")
		// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
		// yaitu "Hello Baik" menggunakan require.Equal
		require.Equal(t, "Hello Baik", result, "Result must be 'Hello Baik'")
	})
}

func TestMain(m *testing.M) {
	// before
	// Mencetak pesan sebelum menjalankan unit test
	fmt.Println("BEFORE UNIT TEST")

	// Menjalankan semua unit test
	m.Run()

	// Mencetak pesan setelah menjalankan unit test
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
			
	// Skip test jika berjalan di sistem operasi Mac OS (darwin)
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	// Memanggil fungsi HelloWorld dengan parameter "Aidil"
	result := HelloWorld("Aidil")
	// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
	// yaitu "Hello Aidil" menggunakan require.Equal
	require.Equal(t, "Hello Aidil", result, "Result must be 'Hello Aidil'")
}

func TestHelloWorldRequire(t *testing.T) {
	// Memanggil fungsi HelloWorld dengan parameter "Aidil"
	result := HelloWorld("Aidil")
	// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
	// yaitu "Hello Aidil" menggunakan require.Equal
	require.Equal(t, "Hello Aidil", result, "Result must be 'Hello Aidil'")
	// Mencetak pesan bahwa pengujian telah selesai
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	// Memanggil fungsi HelloWorld dengan parameter "Aidil"
	result := HelloWorld("Aidil")
	// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
	// yaitu "Hello Aidil" menggunakan assert.Equal
	assert.Equal(t, "Hello Aidil", result, "Result must be 'Hello Aidil'")
	// Mencetak pesan bahwa pengujian telah selesai
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAidil(t *testing.T) {
	// Memanggil fungsi HelloWorld dengan parameter "Aidil"
	result := HelloWorld("Aidil")

	// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
	// yaitu "Hello Aidil", jika tidak sesuai akan menampilkan error
	if result != "Hello Aidil" {
		// Menampilkan pesan error jika hasil tidak sesuai
		t.Error("Result must be 'Hello Aidil'")
	}

	// Mencetak pesan bahwa pengujian telah selesai
	fmt.Println("TestHelloWorldAidil Done")
}

func TestHelloWorldBaik(t *testing.T) {
	// Memanggil fungsi HelloWorld dengan parameter "Baik"
	result := HelloWorld("Baik")

	// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
	// yaitu "Hello Baik", jika tidak sesuai akan menampilkan error fatal
	if result != "Hello Baik" {
		// Menampilkan pesan error fatal jika hasil tidak sesuai
		t.Fatal("Result must be 'Hello Baik'")
	}

	// Mencetak pesan bahwa pengujian telah selesai
	fmt.Println("TestHelloWorldBaik Done")
}

func TestHelloWorldAdam(t *testing.T) {
	// Memanggil fungsi HelloWorld dengan parameter "Adam"
	result := HelloWorld("Adam")

	// Memastikan hasil yang dikembalikan sesuai dengan yang diharapkan
	// yaitu "Hello Adam", jika tidak sesuai akan menampilkan error fatal
	if result != "Hello Adam" {
		// Menampilkan pesan error fatal jika hasil tidak sesuai
		t.Fatal("Result must be 'Hello Adam'")
	}

	// Mencetak pesan bahwa pengujian telah selesai
	fmt.Println("TestHelloWorldAdam Done")
}
