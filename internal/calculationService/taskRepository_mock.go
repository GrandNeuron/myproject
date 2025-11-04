package calculationService

import (
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository — поддельный репозиторий
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateCalculation(task Calculation) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllCalculations() ([]Calculation, error) {
	args := m.Called()
	if res := args.Get(0); res != nil {
		return res.([]Calculation), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskRepository) GetCalculationByID(id string) (Calculation, error) {
	args := m.Called(id)
	return args.Get(0).(Calculation), args.Error(1)
}

func (m *MockTaskRepository) UpdateCalculation(task Calculation) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteCalculation(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
