package usecase

import (
	"context"
	"taskService/internal/entity"
	"time"
)

type TaskCase struct {
	taskRepo TaskRepo
	timeout  time.Duration
}

func NewTaskCase(tr TaskRepo) *TaskCase {
	return &TaskCase{
		taskRepo: tr,
		timeout:  time.Duration(2) * time.Second,
	}
}

func (tc *TaskCase) Create(c context.Context, task *entity.Task) (*entity.Task, error) {
	ctx, cansel := context.WithTimeout(c, tc.timeout)
	defer cansel()
	res, err := tc.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (tc *TaskCase) Delete(c context.Context, id int) error {
	ctx, cansel := context.WithTimeout(c, tc.timeout)
	defer cansel()
	if err := tc.taskRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (tc *TaskCase) Get(c context.Context, id int) (*entity.Task, error) {
	ctx, cansel := context.WithTimeout(c, tc.timeout)
	defer cansel()
	res, err := tc.taskRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (tc *TaskCase) Update(c context.Context, updateTask *entity.Task) error {
	ctx, cansel := context.WithTimeout(c, tc.timeout)
	defer cansel()
	if err := tc.taskRepo.Update(ctx, updateTask); err != nil {
		return err
	}
	return nil

}
