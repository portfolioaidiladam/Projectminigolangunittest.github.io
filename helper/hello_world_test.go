// Package helper berisi fungsi-fungsi pembantu dan pengujiannya
package helper

// Mengimpor package yang diperlukan
import (
	"fmt"     // untuk fungsi print
	"runtime" // untuk mengakses informasi runtime sistem
	"testing" // untuk melakukan unit testing

	"github.com/stretchr/testify/assert"  // untuk assertion yang lebih baik
	"github.com/stretchr/testify/require" // untuk assertion yang strict
)

// BenchmarkTable adalah fungsi benchmark yang menggunakan table-driven test
// untuk menguji performa fungsi HelloWorld dengan berbagai input
func BenchmarkTable(b *testing.B) {
	// Mendefinisikan slice dari struct yang berisi data test
	// Setiap struct memiliki field name untuk nama benchmark dan request untuk parameter input
	benchmarks := []struct {
		name    string
		request string
	}{
		// Test case 1: Menguji dengan input "Aidil"
		{
			name:    "Aidil",
			request: "Aidil",
		},
		// Test case 2: Menguji dengan input "Adam"
		{
			name:    "Adam",
			request: "Adam",
		},
		// Test case 3: Menguji dengan input yang lebih panjang "Aidil Adam Baik"
		{
			name:    "AidilAdamBaik",
			request: "Aidil Adam Baik",
		},
		// Test case 4: Menguji dengan input "Budi Nugraha"
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	// Melakukan iterasi untuk setiap test case dalam benchmarks
	for _, benchmark := range benchmarks {
		// Menjalankan sub-benchmark untuk setiap test case
		// menggunakan nama yang telah ditentukan
		b.Run(benchmark.name, func(b *testing.B) {
			// Melakukan perulangan sebanyak b.N kali untuk mengukur performa
			// b.N adalah nilai yang ditentukan oleh Go testing framework
			for i := 0; i < b.N; i++ {
				// Memanggil fungsi HelloWorld dengan parameter dari test case
				HelloWorld(benchmark.request)
			}
		})
	}
}

// BenchmarkSub adalah fungsi benchmark yang menguji performa fungsi HelloWorld
// dengan beberapa sub-benchmark yang berbeda
func BenchmarkSub(b *testing.B) {
	// Menjalankan sub-benchmark untuk nama "Aidil"
	b.Run("Aidil", func(b *testing.B) {
		// Melakukan perulangan sebanyak b.N kali untuk mengukur performa
		for i := 0; i < b.N; i++ {
			HelloWorld("Aidil")
		}
	})
	// Menjalankan sub-benchmark untuk nama "Adam"
	b.Run("Adam", func(b *testing.B) {
		// Melakukan perulangan sebanyak b.N kali untuk mengukur performa
		for i := 0; i < b.N; i++ {
			HelloWorld("Adam")
		}
	})
}

// BenchmarkHelloWorld adalah fungsi benchmark yang menguji performa
// fungsi HelloWorld dengan parameter "Aidil"
func BenchmarkHelloWorld(b *testing.B) {
	// Melakukan perulangan sebanyak b.N kali untuk mengukur performa
	for i := 0; i < b.N; i++ {
		HelloWorld("Aidil")
	}
}

// BenchmarkHelloWorldAdam adalah fungsi benchmark yang menguji performa
// fungsi HelloWorld dengan parameter "Adam"
func BenchmarkHelloWorldAdam(b *testing.B) {
	// Melakukan perulangan sebanyak b.N kali untuk mengukur performa
	for i := 0; i < b.N; i++ {
		HelloWorld("Adam")
	}
}

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
