package calculationService

// Calculation — основная модель для таблицы в базе данных.
// Здесь хранятся выражение и его результат.
type Calculation struct {
	ID         string `gorm:"primaryKey" json:"id"` // Уникальный идентификатор записи
	Expression string `json:"expression"`           // Выражение (например, "2+2")
	Result     string `json:"result"`               // Результат вычисления (например, "4")
}

// CalculationRequest — структура для приёма данных от пользователя.
// Используется, когда фронтенд отправляет JSON с выражением.
type CalculationRequest struct {
	Expression string `json:"expression"` // Входное выражение для вычисления
}
