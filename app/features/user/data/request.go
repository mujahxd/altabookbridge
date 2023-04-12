package data

type RegisterUserInput struct {
	Name     string `validate:"required,min=2,max=100" json:"name"`
	Username string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}
