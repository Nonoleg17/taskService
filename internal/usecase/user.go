package usecase

import (
	"context"
	"strconv"
	"taskService/internal/entity"
	"taskService/pkg/util"
	"time"
)

type UserCase struct {
	repo    UserRepo
	timeout time.Duration
}

func NewUserCase(repo UserRepo) *UserCase {
	return &UserCase{
		repo:    repo,
		timeout: time.Duration(2) * time.Second,
	}
}

func (uc *UserCase) CreateUser(c context.Context, req *entity.CreateUserReq) (*entity.CreateUserRes, error) {
	ctx, cansel := context.WithTimeout(c, uc.timeout)
	defer cansel()

	hashedPass, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPass,
	}
	r, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &entity.CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil

}

func (ur *UserCase) Login(c context.Context, email string) (*entity.LoginUserRes, error) {
	//TODO доделать
	return nil, nil
}
