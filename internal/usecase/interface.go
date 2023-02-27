package usecase

import (
	"context"
	"taskService/internal/entity"
)

type (
	User interface {
		CreateUser(context.Context, *entity.CreateUserReq) (*entity.CreateUserRes, error)
		Login(context.Context, *entity.LoginUserReq) (*entity.LoginUserRes, error)
	}

	UserRepo interface {
		CreateUser(context.Context, *entity.User) (*entity.User, error)
		GetUserByEmail(context.Context, string) (*entity.User, error)
	}
)
