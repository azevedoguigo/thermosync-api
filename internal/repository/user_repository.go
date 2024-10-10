package repository

import (
	"github.com/azevedoguigo/thermosync-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (u *userRepository) FindByEmail(email string) (*domain.User, error) {
	panic("unimplemented")
}

func (u *userRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}
