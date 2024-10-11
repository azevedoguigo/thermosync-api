package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azevedoguigo/thermosync-api/internal/contract"
	"github.com/azevedoguigo/thermosync-api/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) CreateUser(userDTO *contract.NewUserDTO) error {
	args := m.Called(userDTO)
	return args.Error(0)
}

func (m *mockUserService) FindUserByID(id uuid.UUID) (*domain.User, error) {
	args := m.Called(id)
	if user := args.Get(0); user != nil {
		return user.(*domain.User), args.Error(0)
	}

	return nil, args.Error(1)
}

func (m *mockUserService) Login(email string, password string) (string, error) {
	panic("unimplemented")
}

func TestUserHandler_Create_Success(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{
		"first_name": "Ayrton",
		"last_name":  "Senna",
		"email":      "senna@example.com",
		"password":   "supersenha",
	})

	mockService := new(mockUserService)
	mockService.On("CreateUser", mock.Anything).Return(nil)

	handler := NewUserHandler(mockService)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	recorderResponse := httptest.NewRecorder()
	handler.CreateUser(recorderResponse, req)

	assert.Equal(t, http.StatusCreated, recorderResponse.Code)
}
