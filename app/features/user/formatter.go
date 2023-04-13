package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

type UserLoginFormatter struct {
	Token string `json:"token"`
}

type GetUserProfileFormatter struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

func FormatLoginUser(user User, token string) UserLoginFormatter {
	formatter := UserLoginFormatter{
		Token: token,
	}

	return formatter
}

func FormatGetProfile(user User) GetUserProfileFormatter {
	formatter := GetUserProfileFormatter{
		Username: user.Username,
		Name:     user.Name,
		Password: user.Password,
		Avatar:   user.AvatarFileName,
	}
	return formatter
}
