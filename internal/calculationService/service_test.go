package calculationService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCalculation(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		mockSetup func(m *MockTaskRepository, input string)
		wantErr   bool
	}{
		{
			name:  "успешное создание задачи",
			input: "10+5",
			mockSetup: func(m *MockTaskRepository, input string) {
				m.On("CreateCalculation", mock.MatchedBy(func(c Calculation) bool {
					return c.Expression == input && c.Result == "15" && c.ID != ""
				})).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при создании",
			input: "20*2",
			mockSetup: func(m *MockTaskRepository, input string) {
				m.On("CreateCalculation", mock.MatchedBy(func(c Calculation) bool {
					return c.Expression == input && c.Result == "40" && c.ID != ""
				})).Return(errors.New("db error"))
			},
			wantErr: true,
		},
		{
			name:      "ошибка при вычислении выражения",
			input:     "invalid expression",
			mockSetup: func(m *MockTaskRepository, input string) {},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.input)

			service := NewCalculationService(mockRepo)
			result, err := service.CreateCalculation(tt.input, "")

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result.Expression)
				assert.NotEmpty(t, result.Result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllCalculations(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
		wantCount int
	}{
		{
			name: "успешное получение всех задач",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllCalculations").Return([]Calculation{
					{ID: "1", Expression: "10+5", Result: "15"},
					{ID: "2", Expression: "20*2", Result: "40"},
				}, nil)
			},
			wantErr:   false,
			wantCount: 2,
		},
		{
			name: "ошибка при получении",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllCalculations").Return(nil, errors.New("db error"))
			},
			wantErr:   true,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewCalculationService(mockRepo)
			result, err := service.GetAllCalculations()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result, tt.wantCount)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateCalculation(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		expression string
		mockSetup  func(m *MockTaskRepository, id, expression string)
		wantErr    bool
	}{
		{
			name:       "успешное обновление задачи",
			id:         "1",
			expression: "100+50",
			mockSetup: func(m *MockTaskRepository, id, expression string) {
				m.On("UpdateCalculation", Calculation{
					ID:         id,
					Expression: expression,
					Result:     "150",
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name:       "ошибка при обновлении",
			id:         "2",
			expression: "50-10",
			mockSetup: func(m *MockTaskRepository, id, expression string) {
				m.On("UpdateCalculation", Calculation{
					ID:         id,
					Expression: expression,
					Result:     "40",
				}).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id, tt.expression)

			service := NewCalculationService(mockRepo)
			result, err := service.UpdateCalculation(tt.id, tt.expression)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.id, result.ID)
				assert.Equal(t, tt.expression, result.Expression)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteCalculation(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(m *MockTaskRepository, id string)
		wantErr   bool
	}{
		{
			name: "успешное удаление задачи",
			id:   "1",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteCalculation", id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при удалении",
			id:   "2",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteCalculation", id).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewCalculationService(mockRepo)
			err := service.DeleteCalculation(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
