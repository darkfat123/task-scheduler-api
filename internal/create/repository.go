package create

import (
	"context"
	"task-scheduler-api/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateTaskRepository interface {
	Create(ctx context.Context, req CreateRequest) error
}

type createTaskRepositoryImpl struct {
	queries *db.Queries
}

func NewCreateTaskRepository(conn db.DBTX) CreateTaskRepository {
	return &createTaskRepositoryImpl{
		queries: db.New(conn),
	}
}

func (r *createTaskRepositoryImpl) Create(ctx context.Context, req CreateRequest) error {
	params := db.CreateTaskParams{
		Code:          pgtype.Text{String: req.Code, Valid: true},
		Name:          pgtype.Text{String: req.Name, Valid: true},
		FrequencyDate: pgtype.Text{String: req.FrequencyDate, Valid: true},
		FrequencyTime: pgtype.Text{String: req.FrequencyTime, Valid: true},
		MaxRetries:    pgtype.Int4{Int32: req.MaxRetries, Valid: true},
	}

	_, err := r.queries.CreateTask(ctx, params)
	return err
}
