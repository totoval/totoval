package requests

type UserRegister struct {
	Email      string     `json:"email" validate:"required,email"`
	Password   string     `json:"password" validate:"required,min=8,max=24"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}