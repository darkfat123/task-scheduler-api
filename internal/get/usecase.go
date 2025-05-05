package get

import (
	"context"
	"task-scheduler-api/pkg/common"
)

type GetTaskByCodeUsecase interface {
	Execute(ctx context.Context, param string) (GetByCodeResponse, error)
}

type getTaskByCodeUsecaseImpl struct {
	repo GetTaskByCodeRepository
}

func NewGetTaskByCodeUsecase(repo GetTaskByCodeRepository) GetTaskByCodeUsecase {
	return &getTaskByCodeUsecaseImpl{repo: repo}
}

func (u *getTaskByCodeUsecaseImpl) Execute(ctx context.Context, param string) (GetByCodeResponse, error) {
	if param == "" {
		return GetByCodeResponse{}, common.ErrInvalidTask
	}
	task, err := u.repo.GetTaskByCode(ctx, param)
	if err != nil {
		return GetByCodeResponse{}, err
	}

	return GetByCodeResponse{
		Code:          task.Code.String,
		Name:          task.Name.String,
		FrequencyDate: task.FrequencyDate.String,
		FrequencyTime: task.FrequencyTime.String,
		NextRunAt:     task.NextRunAt.Time,
		LastRunAt:     task.LastRunAt.Time,
		MaxRetries:    task.MaxRetries.Int32,
		Status:        task.Status,
		IsEnabled:     task.IsEnabled,
	}, nil
}
