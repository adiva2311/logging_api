package repositories

import (
	"logging-api/models"

	"gorm.io/gorm"
)

type dbLogging struct {
	Conn *gorm.DB
}

// Create implements LoggingRepository.
func (db *dbLogging) Create(logging models.Logging) error {
	return db.Conn.Create(&logging).Error
}

// Delete implements LoggingRepository.
func (db *dbLogging) Delete(id int) error {
	return db.Conn.Where("id", id).Delete(&models.Logging{}).Error
}

// GetAll implements LoggingRepository.
func (db *dbLogging) GetAll() ([]models.Logging, error) {
	var data []models.Logging
	result := db.Conn.Find(&data)
	return data, result.Error
}

// Update implements LoggingRepository.
func (db *dbLogging) Update(id int, logging models.Logging) error {
	return db.Conn.Where("id", id).Updates(logging).Error
}

type LoggingRepository interface {
	Create(logging models.Logging) error
	Update(id int, logging models.Logging) error
	Delete(id int) error
	GetAll() ([]models.Logging, error)
}

//Untuk mengimplementasikan Semua yang ada di dalam interface
// Sebuah shortcut untuk membuat semua function Create, Delete, Update, dll

/*

func New`_NAMAFUNCTION` (`ISI STRUCT`) `NAMA_INTERFACE` {
	return &`NAMA STRUCT`{`NAMA VARIABEL STRUCT: NAMA_VARIABEL_PARAMETER`}
}

*/
func NewLoggingRepository(Conn *gorm.DB) LoggingRepository {
	return &dbLogging{Conn: Conn}
}
