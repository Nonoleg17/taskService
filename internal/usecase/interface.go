package usecase

import (
	"context"
	"taskService/internal/entity"
)

type (
	User interface {
		CreateUser(context.Context, *entity.CreateUserReq) (*entity.CreateUserRes, error)
		Login(context.Context, *entity.LoginUserReq) (*entity.LoginUserRes, *entity.Session, error)
	}

	UserRepo interface {
		CreateUser(context.Context, *entity.User) (*entity.User, error)
		GetUserByEmail(context.Context, string) (*entity.User, error)
	}
	SessionRepo interface {
		Set(string, string) (*entity.Session, error)
		Get(string) error
	}
)
