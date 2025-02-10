package repository

// Package repository berisi implementasi mock untuk keperluan unit testing
// Mock repository memungkinkan kita untuk mensimulasikan perilaku repository
// tanpa perlu mengakses database atau storage yang sebenarnya
import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

// CategoryRepositoryMock adalah struct yang mengimplementasikan interface CategoryRepository
// untuk keperluan testing. Struct ini menggunakan mock.Mock dari testify
type CategoryRepositoryMock struct {
	// Mock adalah instance dari mock.Mock yang akan menyimpan ekspektasi dan panggilan
	Mock mock.Mock
}

// FindById adalah implementasi mock untuk method FindById
// Parameter:
//   - id: string - ID category yang dicari
// Return:
//   - *entity.Category - pointer ke Category jika ditemukan, nil jika tidak
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	// Called() merekam bahwa method ini dipanggil dengan parameter id
	// dan mengembalikan nilai sesuai ekspektasi yang telah di-setup
	arguments := repository.Mock.Called(id)
	
	// Cek apakah nilai yang dikembalikan nil
	if arguments.Get(0) == nil {
		return nil
	} else {
		// Konversi nilai yang dikembalikan ke tipe entity.Category
		// dan kembalikan pointer ke nilai tersebut
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
