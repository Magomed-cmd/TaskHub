package repository

import (
	"TaskHub/pkg/model"
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	Get(ctx context.Context) ([]model.Task, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, task *model.Task) (*model.Task, error)
	GetTaskByID(ctx context.Context, id int) (*model.Task, error)
}
