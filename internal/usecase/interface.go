package usecase

import (
	"context"
	"taskService/internal/entity"
)

type (
	User interface {
		CreateUser(context.Context, *entity.User) (*entity.User, error)
		Login(context.Context, string) (*entity.User, error)
	}

	UserRepo interface {
		CreateUser(context.Context, *entity.User) (*entity.User, error)
		GetUserByEmail(context.Context, string) (*entity.User, error)
	}
)
