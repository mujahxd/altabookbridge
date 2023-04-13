package user

import "github.com/mujahxd/altabookbridge/app/features/user/data"

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

type UserLoginFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

func FormatLoginUser(user User, token string) data.LoginResponse {
	formatter := data.LoginResponse{
		Token: token,
	}

	return formatter
}
