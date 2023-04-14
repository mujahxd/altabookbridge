package data

type RegisterUserInput struct {
	Name     string `binding:"required" json:"name"`
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type LoginInput struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type UpdateInput struct {
	Name           string `json:"name" binding:"required"`
	Password       string `json:"password" binding:"required"`
	AvatarFileName string `json:"avatar" binding:"required"`
}
