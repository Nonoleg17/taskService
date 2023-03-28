package repo

import (
	"context"
	"taskService/internal/entity"
	"taskService/pkg/postgres"
)

type TaskRepo struct {
	pg *postgres.Postgres
}

func NewTaskRepo(pg *postgres.Postgres) *TaskRepo {
	return &TaskRepo{
		pg: pg,
	}
}

func (tr *TaskRepo) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	if err := tr.pg.DbConnect.Create(&task).Error; err != nil {
		return nil, err
	}
	if err := tr.pg.DbConnect.WithContext(ctx).Last(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *TaskRepo) Delete(ctx context.Context, id int) error {
	if err := tr.pg.DbConnect.WithContext(ctx).Delete(&entity.Task{}, id).Error; err != nil {
		return err
	}
	return nil

}
func (tr *TaskRepo) Get(ctx context.Context, id int) (*entity.Task, error) {
	var task entity.Task
	if err := tr.pg.DbConnect.WithContext(ctx).First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *TaskRepo) Update(ctx context.Context, task *entity.Task) error {
	if err := tr.pg.DbConnect.WithContext(ctx).Save(task).Error; err != nil {
		return err
	}
	return nil
}
