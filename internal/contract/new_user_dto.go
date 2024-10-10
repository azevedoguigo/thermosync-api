package contract

type NewUserDTO struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email,min=6,max=60"`
	Password  string `json:"password" validate:"required,min=6,max=30"`
}
