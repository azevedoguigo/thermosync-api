package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}
