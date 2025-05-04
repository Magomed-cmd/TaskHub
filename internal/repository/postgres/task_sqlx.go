package postgres

import (
	"TaskHub/internal/repository"
	"TaskHub/pkg/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) repository.TaskRepository {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(ctx context.Context, task *model.Task) error {
	query := `
		INSERT INTO task (title, detail, status, priority, due_date, user_id, assignee_id)
		VALUES (:title, :detail, :status, :priority, :due_date, :user_id, :assignee_id)
		RETURNING id, created_at, updated_at
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	return stmt.GetContext(ctx, task, task)
}

func (r *TaskRepo) Get(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.SelectContext(ctx, &tasks, "SELECT * FROM task ORDER BY created_at DESC")
	return tasks, err
}

func (r *TaskRepo) Delete(ctx context.Context, id int) error {
	var task model.Task

	err := r.db.GetContext(ctx, &task, "SELECT * FROM task WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}

	_, err = r.db.ExecContext(ctx, "DELETE FROM task WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepo) Update(ctx context.Context, task *model.Task) error {
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
		return nil
	}

	updates["updated_at"] = time.Now()
	updates["id"] = task.ID

	set := make([]string, 0, len(updates))
	for k := range updates {
		if k != "id" {
			set = append(set, fmt.Sprintf("%s = :%s", k, k))
		}
	}

	query := fmt.Sprintf("UPDATE task SET %s WHERE id = :id", strings.Join(set, ", "))

	_, err := r.db.NamedExecContext(ctx, query, updates)
	return err
}
