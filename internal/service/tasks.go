package service

import (
	"TaskHub/internal/repository"
	"TaskHub/pkg/model"
	"context"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, task *model.Task) error {
	return s.repo.Create(ctx, task)
}

func (s *TaskService) GetTasks(ctx context.Context) ([]model.Task, error) {
	return s.repo.Get(ctx)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	return s.repo.Update(ctx, task)
}

func (s *TaskService) GetTaskByID(ctx context.Context, id int) (*model.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}
