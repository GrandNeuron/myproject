package handlers

import (
	"context"
	"fmt"

	calculationService "CalculatorAppFrontendPantela-main/internal/calculationService"
	"CalculatorAppFrontendPantela-main/internal/web/tasks"
)

// TaskHandler — структура, адаптирующая CalculationService для tasks API
type TaskHandler struct {
	service calculationService.CalculationService
}

// NewTaskHandler — конструктор для создания нового task хендлера
func NewTaskHandler(s calculationService.CalculationService) *TaskHandler {
	return &TaskHandler{service: s}
}

// GetTasks - реализация получения всех задач (вычислений)
func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	calculations, err := h.service.GetAllCalculations()
	if err != nil {
		return nil, err
	}

	// Конвертируем Calculation в Task
	result := make([]tasks.Task, 0, len(calculations))
	for _, calc := range calculations {
		isDone := calc.Result != ""
		task := tasks.Task{
			Id:     nil, // можно добавить парсинг ID если нужно
			IsDone: &isDone,
			Task:   &calc.Expression,
			Result: &calc.Result,
		}
		result = append(result, task)
	}

	return tasks.GetTasks200JSONResponse(result), nil
}

// PostTasks - реализация создания новой задачи (вычисления)
func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil || request.Body.Task == nil {
		return nil, nil
	}

	// TODO: Extract user_id from request body when frontend supports it
	calc, err := h.service.CreateCalculation(*request.Body.Task, "")
	if err != nil {
		return nil, err
	}

	isDone := calc.Result != ""
	result := tasks.Task{
		Id:     nil,
		IsDone: &isDone,
		Task:   &calc.Expression,
		Result: &calc.Result,
	}

	return tasks.PostTasks201JSONResponse(result), nil
}

// PatchTasksId - реализация обновления задачи (вычисления)
func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	if request.Body == nil || request.Body.Task == nil {
		return nil, nil
	}

	// Конвертируем uint в string для ID
	idStr := fmt.Sprintf("%d", request.Id)

	calc, err := h.service.UpdateCalculation(idStr, *request.Body.Task)
	if err != nil {
		return nil, err
	}

	isDone := calc.Result != ""
	result := tasks.Task{
		Id:     &request.Id,
		IsDone: &isDone,
		Task:   &calc.Expression,
		Result: &calc.Result,
	}

	return tasks.PatchTasksId200JSONResponse(result), nil
}

// DeleteTasksId - реализация удаления задачи
func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Конвертируем uint в string для ID
	idStr := fmt.Sprintf("%d", request.Id)

	if err := h.service.DeleteCalculation(idStr); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}
