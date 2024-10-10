package service

import (
	"github.com/azevedoguigo/thermosync-api/internal/contract"
	"github.com/azevedoguigo/thermosync-api/internal/domain"
	"github.com/azevedoguigo/thermosync-api/internal/repository"
	"github.com/azevedoguigo/thermosync-api/pkg"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userDTO *contract.NewUserDTO) error
	GetUserByID(id uuid.UUID) (*domain.User, error)
	Login(email, password string) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) CreateUser(userDTO *contract.NewUserDTO) error {
	if err := pkg.ValidateStruct(userDTO); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		ID:        uuid.New(),
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		Password:  string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GetUserByID(id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

func (u *userService) Login(email string, password string) (string, error) {
	panic("unimplemented")
}
