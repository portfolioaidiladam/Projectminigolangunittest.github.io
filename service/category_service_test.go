// Package service - mendefinisikan package untuk testing service
package service

// Import package yang dibutuhkan
import (
	"belajar-golang-unit-test/entity"     // Import entity untuk struktur Category
	"belajar-golang-unit-test/repository" // Import repository untuk interface CategoryRepository
	"testing"                             // Import testing package dari Go

	"github.com/stretchr/testify/assert" // Import assert untuk assertion dalam testing
	"github.com/stretchr/testify/mock"   // Import mock untuk membuat mock object
)

// Membuat instance mock repository dan service yang akan digunakan di semua test
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// Test case untuk skenario ketika kategori tidak ditemukan
func TestCategoryService_GetNotFound(t *testing.T) {
	// Setup mock untuk mengembalikan nil saat FindById dipanggil dengan parameter "1"
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	
	// Eksekusi method Get
	category, err := categoryService.Get("1")
	
	// Verifikasi hasil
	assert.Nil(t, category)      // Memastikan category adalah nil
	assert.NotNil(t, err)        // Memastikan error tidak nil
}

// Test case untuk skenario ketika kategori berhasil ditemukan
func TestCategoryService_GetSuccess(t *testing.T) {
	// Membuat data dummy category
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	
	// Setup mock untuk mengembalikan category saat FindById dipanggil dengan parameter "2"
	categoryRepository.Mock.On("FindById", "2").Return(category)
	
	// Eksekusi method Get
	result, err := categoryService.Get("2")
	
	// Verifikasi hasil
	assert.Nil(t, err)                      // Memastikan tidak ada error
	assert.NotNil(t, result)                // Memastikan result tidak nil
	assert.Equal(t, category.Id, result.Id)     // Memastikan ID sesuai
	assert.Equal(t, category.Name, result.Name) // Memastikan Name sesuai
}