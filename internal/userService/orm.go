package userService

import (
	"time"

	calculationService "CalculatorAppFrontendPantela-main/internal/calculationService"
)

// User — модель пользователя
type User struct {
	ID        string                              `gorm:"primaryKey" json:"id"`
	Email     string                              `gorm:"unique;not null" json:"email"`
	Password  string                              `gorm:"not null" json:"-"` // Не возвращаем пароль в JSON
	DeletedAt *time.Time                         `json:"deleted_at,omitempty"`
	CreatedAt time.Time                           `json:"created_at"`
	UpdatedAt time.Time                           `json:"updated_at"`
	Tasks     []calculationService.Calculation `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}

// UserRequest — структура для создания/обновления пользователя
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
