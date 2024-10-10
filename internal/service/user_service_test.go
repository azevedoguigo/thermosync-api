package service

import (
	"errors"
	"testing"

	"github.com/azevedoguigo/thermosync-api/internal/contract"
	"github.com/azevedoguigo/thermosync-api/internal/domain"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) FindByEmail(email string) (*domain.User, error) {
	panic("unimplemented")
}

func (m *mockUserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

func TestUserService_CreateUser_Success(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	mockRepo.On("Create", mock.Anything).Return(nil)

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.NoError(t, err)

	mockRepo.AssertNumberOfCalls(t, "Create", 1)
}

func TestUserService_CreateUser_Error(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	mockRepo.On("Create", mock.Anything).Return(errors.New("database error"))

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Error(t, err)

	mockRepo.AssertNumberOfCalls(t, "Create", 1)
}

func TestUserService_CreateUser_FirstNameIsRequired(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "",
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "FirstName is required", err.Error())
}

func TestUserService_CreateUser_MustValidFirstNameMinLenght(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "A",
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "FirstName is required with min: 2", err.Error())
}

func TestUserService_CreateUser_MustValidFirstNameMaxLenght(t *testing.T) {
	fake := faker.New()

	userDTO := &contract.NewUserDTO{
		FirstName: fake.Lorem().Text(52),
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "FirstName is required with max: 50", err.Error())
}

func TestUserService_CreateUser_LastNameIsRequired(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "",
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "LastName is required", err.Error())
}

func TestUserService_CreateUser_MustValidLastNameMinLenght(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "S",
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "LastName is required with min: 2", err.Error())
}

func TestUserService_CreateUser_MustValidLastNameMaxLenght(t *testing.T) {
	fake := faker.New()

	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  fake.Lorem().Text(52),
		Email:     "senna@example.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)
	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "LastName is required with max: 50", err.Error())
}

func TestUserService_CreateUser_EmailIsRequired(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     "",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "Email is required", err.Error())
}

func TestUserService_CreateUser_InvalidEmail(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     "invalid.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "Email is invalid.", err.Error())
}

func TestUserService_CreateUser_MustValidEmailMaxLength(t *testing.T) {
	fake := faker.New()

	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     fake.Lorem().Text(58) + "@gmail.com",
		Password:  "supersenha",
	}

	mockRepo := new(mockUserRepository)

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "Email is required with max: 60", err.Error())
}

func TestUserService_CreateUser_PasswordIsRequired(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  "",
	}

	mockRepo := new(mockUserRepository)

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "Password is required", err.Error())
}

func TestUserService_CreateUser_MustValidPasswordMinLength(t *testing.T) {
	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  "12345",
	}

	mockRepo := new(mockUserRepository)

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "Password is required with min: 6", err.Error())
}

func TestUserService_CreateUser_MustValidPasswordMaxLength(t *testing.T) {
	fake := faker.New()

	userDTO := &contract.NewUserDTO{
		FirstName: "Ayrton",
		LastName:  "Senna",
		Email:     "senna@example.com",
		Password:  fake.Lorem().Text(52),
	}

	mockRepo := new(mockUserRepository)

	userService := NewUserService(mockRepo)

	err := userService.CreateUser(userDTO)

	assert.Equal(t, "Password is required with max: 30", err.Error())
}
