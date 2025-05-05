package getall

import (
	"context"
)

type GetAllTaskUsecase interface {
	Execute(ctx context.Context) ([]GetAllResponse, error)
}

type getAllTaskUsecaseImpl struct {
	repo GetAllTaskRepository
}

func NewGetAllTaskUsecase(repo GetAllTaskRepository) GetAllTaskUsecase {
	return &getAllTaskUsecaseImpl{repo: repo}
}

func (u *getAllTaskUsecaseImpl) Execute(ctx context.Context) ([]GetAllResponse, error) {
	taskList, err := u.repo.GetAllTask(ctx)
	if err != nil {
		return nil, err
	}

	var responses []GetAllResponse
	for _, task := range taskList {
		responses = append(responses, GetAllResponse{
			Code:          task.Code.String,
			Name:          task.Name.String,
			FrequencyDate: task.FrequencyDate.String,
			FrequencyTime: task.FrequencyTime.String,
			Status:        task.Status,
			IsEnabled:     task.IsEnabled,
		})
	}

	return responses, nil
}
