package userService

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	calculationService "CalculatorAppFrontendPantela-main/internal/calculationService"
)

// UserService — интерфейс бизнес-логики
type UserService interface {
	CreateUser(email, password string) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id, email, password string) (User, error)
	DeleteUser(id string) error
	GetTasksForUser(userID string) ([]calculationService.Calculation, error)
}

type userService struct {
	repo UserRepository
}

// NewUserService — конструктор сервиса
func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

// hashPassword — хеширование пароля
func (s *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *userService) CreateUser(email, password string) (User, error) {
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id string) (User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id, email, password string) (User, error) {
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:       id,
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *userService) GetTasksForUser(userID string) ([]calculationService.Calculation, error) {
	return s.repo.GetTasksForUser(userID)
}
