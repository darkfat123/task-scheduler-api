package create

import (
	"context"
	"task-scheduler-api/pkg/common"
)

type CreateTaskUsecase interface {
	Execute(ctx context.Context, req CreateRequest) error
}

type createTaskUsecaseImpl struct {
	repo CreateTaskRepository
}

func NewCreateTaskUsecase(repo CreateTaskRepository) CreateTaskUsecase {
	return &createTaskUsecaseImpl{repo: repo}
}

func (u *createTaskUsecaseImpl) Execute(ctx context.Context, req CreateRequest) error {
	if req.Code == "" || req.Name == "" || req.FrequencyDate == "" || req.FrequencyTime == "" || req.MaxRetries < 0 {
		return common.ErrInvalidTask
	}

	return u.repo.Create(ctx, req)
}
