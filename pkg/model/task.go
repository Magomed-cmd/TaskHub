package model

import (
	"time"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
	StatusExpired    TaskStatus = "expired"
)

type TaskPriority string

const (
	PriorityLow      TaskPriority = "low"
	PriorityMedium   TaskPriority = "medium"
	PriorityHigh     TaskPriority = "high"
	PriorityCritical TaskPriority = "critical"
)

type Task struct {
	ID           *int64        `db:"id" json:"id"`
	Title        *string       `db:"title" json:"title"`
	Detail       *string       `db:"detail" json:"detail"`
	Status       *TaskStatus   `db:"status" json:"status"`
	Priority     *TaskPriority `db:"priority" json:"priority"`
	DueDate      *time.Time    `db:"due_date" json:"due_date"`
	CreatedAt    time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time    `db:"updated_at" json:"updated_at"`
	UserID       *int64        `db:"user_id" json:"user_id"`
	AssigneeID   *int64        `db:"assignee_id" json:"assignee_id"`
	ParentTaskID *int64        `db:"parent_task_id" json:"parent_task_id"`
}
