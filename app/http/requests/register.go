package requests

type UserRegister struct {
	Email      string     `json:"email" binding:"required,email"`
	Password   string     `json:"password" binding:"required,min=8,max=24"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}