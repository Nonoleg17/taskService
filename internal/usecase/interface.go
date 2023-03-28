package usecase

import (
	"context"
	"taskService/internal/entity"
)

type (
	User interface {
		CreateUser(context.Context, *entity.CreateUserReq) (*entity.CreateUserRes, error)
		Login(context.Context, *entity.LoginUserReq) (*entity.LoginUserRes, *entity.Session, error)
		Check(string) error
	}
	Task interface {
		Create(context.Context, *entity.Task) (*entity.Task, error)
		Delete(context.Context, int) error
		Get(context.Context, int) (*entity.Task, error)
		Update(context.Context, *entity.Task) error
	}

	UserRepo interface {
		CreateUser(context.Context, *entity.User) (*entity.User, error)
		GetUserByEmail(context.Context, string) (*entity.User, error)
	}
	SessionRepo interface {
		Set(string, string) (*entity.Session, error)
		Get(string) error
	}
	TaskRepo interface {
		Create(context.Context, *entity.Task) (*entity.Task, error)
		Delete(context.Context, int) error
		Get(context.Context, int) (*entity.Task, error)
		Update(context.Context, *entity.Task) error
	}
)
