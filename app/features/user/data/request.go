package data

type RegisterUserInput struct {
	Name     string `binding:"required" json:"name"`
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
