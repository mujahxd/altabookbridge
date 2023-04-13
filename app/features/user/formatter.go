package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

func FormatLoginUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		Token: token,
	}

	return formatter
}
