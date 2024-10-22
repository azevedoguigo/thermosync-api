package pkg

import (
	"errors"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var tokenAuth = jwtauth.New("HS256", []byte("secretKey"), nil)

func ValidateStruct(data interface{}) error {
	validate := validator.New()

	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErros := err.(validator.ValidationErrors)
	validationError := validationErros[0]

	switch validationError.Tag() {
	case "required":
		return errors.New(validationError.StructField() + " is required")
	case "max":
		return errors.New(validationError.StructField() + " is required with max: " + validationError.Param())
	case "min":
		return errors.New(validationError.StructField() + " is required with min: " + validationError.Param())
	case "email":
		return errors.New(validationError.StructField() + " is invalid.")
	}

	return nil
}

func GenerateJWT(userID uuid.UUID) (string, error) {
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
