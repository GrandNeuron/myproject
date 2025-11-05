package calculationService

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

// Интерфейс описывает все операции для бизнес-логики.
type CalculationService interface {
	CreateCalculation(expression, userID string) (Calculation, error)
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

// calcService — структура, реализующая интерфейс CalculationService.
// Здесь мы храним зависимость от репозитория.
type calcService struct {
	repo CalculationRepository
}

// NewCalculationService — конструктор, создающий новый сервис.
func NewCalculationService(repo CalculationRepository) CalculationService {
	return &calcService{repo: repo}
}

// calculateExpression — вспомогательная функция для вычислений.
// Принимает строку (например, "2+2"), возвращает результат ("4").
func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err // Ошибка при создании выражения
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err // Ошибка при вычислении
	}

	return fmt.Sprintf("%v", result), nil
}

// CreateCalculation — создаёт новую запись: вычисляет и сохраняет результат.
func (s *calcService) CreateCalculation(expression, userID string) (Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
		UserID:     userID,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

// GetAllCalculations — возвращает все записи из БД.
func (s *calcService) GetAllCalculations() ([]Calculation, error) {
	return s.repo.GetAllCalculations()
}

// GetCalculationByID — возвращает конкретную запись по ID.
func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

// UpdateCalculation — пересчитывает выражение и обновляет запись в БД.
func (s *calcService) UpdateCalculation(id, expression string) (Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         id,
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

// DeleteCalculation — удаляет запись по ID.
func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}
