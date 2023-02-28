package repo

import (
	"context"
	"taskService/internal/entity"
	"taskService/pkg/postgres"
)

type UserRepo struct {
	pg *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{
		pg: pg,
	}
}

func (ur *UserRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := ur.pg.DbConnect.Create(&user).Error; err != nil {
		return nil, err
	}
	if err := ur.pg.DbConnect.WithContext(ctx).Last(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := ur.pg.DbConnect.WithContext(ctx).Find(&user).Where("email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
