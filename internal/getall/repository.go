package getall

import (
	"context"
	"task-scheduler-api/db"
)

type GetAllTaskRepository interface {
	GetAllTask(ctx context.Context) ([]db.Task, error)
}

type getAllTaskRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAllTaskRepository(conn db.DBTX) GetAllTaskRepository {
	return &getAllTaskRepositoryImpl{
		queries: db.New(conn),
	}
}

func (r *getAllTaskRepositoryImpl) GetAllTask(ctx context.Context) ([]db.Task, error) {
	return r.queries.ListTasks(ctx)
}
