package calculationService

import (
	"gorm.io/gorm"
)

// Основные методы CRUD — Create, Read, Update, Delete.
type CalculationRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

// calcRepository — структура, которая реализует интерфейс CalculationRepository.
// Внутри хранится объект базы данных GORM.
type calcRepository struct {
	db *gorm.DB
}

// NewCalculationRepository — конструктор, который создаёт новый репозиторий.
func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRepository{db: db}
}

// CreateCalculation — создаёт (добавляет) новую запись в БД.
func (r *calcRepository) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

// GetAllCalculations — возвращает все записи из таблицы calculations.
func (r *calcRepository) GetAllCalculations() ([]Calculation, error) {
	var calculations []Calculation
	err := r.db.Find(&calculations).Error
	return calculations, err
}

// GetCalculationByID — ищет конкретную запись по ID.
func (r *calcRepository) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation
	err := r.db.First(&calc, "id = ?", id).Error
	return calc, err
}

// UpdateCalculation — обновляет существующую запись (по ID из calc).
func (r *calcRepository) UpdateCalculation(calc Calculation) error {
	return r.db.Save(&calc).Error
}

// DeleteCalculation — удаляет запись по ID.
func (r *calcRepository) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{}, "id = ?", id).Error
}
