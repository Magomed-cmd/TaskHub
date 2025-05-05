package postgres

import (
	"TaskHub/internal/repository"
	"TaskHub/pkg/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) repository.TaskRepository {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(ctx context.Context, task *model.Task) error {
	err := r.db.WithContext(ctx).Create(task).Error
	if err != nil {
		log.Printf("failed to create task: %+v, error: %v", task, err)
		return err
	}
	return nil
}

func (r *TaskRepo) Get(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.WithContext(ctx).Order("created_at DESC").Find(&tasks)
	if result.Error != nil {
		log.Printf("failed to fetch tasks: %v", result.Error)
		return nil, result.Error
	}
	return tasks, nil
}

func (r *TaskRepo) GetTaskByID(ctx context.Context, id int) (*model.Task, error) {

	task := model.Task{}
	result := r.db.WithContext(ctx).First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		log.Printf("failed to fetch task with id=%d: %v", id, result.Error)
		return nil, result.Error
	}
	return &task, nil
}

func (r *TaskRepo) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&model.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Printf("no task found to delete with id=%d", id)
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *TaskRepo) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	updates := map[string]interface{}{}

	if task.Title != nil {
		updates["title"] = *task.Title
	}
	if task.Detail != nil {
		updates["detail"] = *task.Detail
	}
	if task.Status != nil {
		updates["status"] = string(*task.Status)
	}
	if task.Priority != nil {
		updates["priority"] = string(*task.Priority)
	}
	if task.DueDate != nil {
		updates["due_date"] = *task.DueDate
	}
	if task.AssigneeID != nil {
		updates["assignee_id"] = *task.AssigneeID
	}
	if task.ParentTaskID != nil {
		updates["parent_task_id"] = *task.ParentTaskID
	}

	if len(updates) == 0 {
		return nil, nil // ничего обновлять — ничего делать
	}

	updates["updated_at"] = time.Now()

	result := r.db.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ?", task.ID).
		Updates(updates)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var updated model.Task
	if err := r.db.WithContext(ctx).First(&updated, task.ID).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}
