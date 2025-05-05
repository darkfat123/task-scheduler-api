package get

import (
	"context"
	"task-scheduler-api/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetTaskByCodeRepository interface {
	GetTaskByCode(ctx context.Context, code string) (db.Task, error)
}

type getTaskByCodeRepositoryImpl struct {
	queries *db.Queries
}

func NewGetTaskByCodeRepository(conn db.DBTX) GetTaskByCodeRepository {
	return &getTaskByCodeRepositoryImpl{
		queries: db.New(conn),
	}
}

func (r *getTaskByCodeRepositoryImpl) GetTaskByCode(ctx context.Context, code string) (db.Task, error) {
	return r.queries.GetTasksByCode(ctx, pgtype.Text{String: code, Valid: true})
}
