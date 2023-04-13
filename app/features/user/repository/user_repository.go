package repository

import "github.com/mujahxd/altabookbridge/app/features/user"

type Repository interface {
	Save(user user.User) (user.User, error)
}
