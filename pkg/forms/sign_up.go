package forms

type SignUp struct {
	Name     string `json:"name" validate:"required,max=45,min=1"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
