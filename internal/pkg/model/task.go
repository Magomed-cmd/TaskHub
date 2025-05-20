package model

import (
	"time"
)

type TaskStatus string
type TaskPriority string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
	StatusExpired    TaskStatus = "expired"

	PriorityLow      TaskPriority = "low"
	PriorityMedium   TaskPriority = "medium"
	PriorityHigh     TaskPriority = "high"
	PriorityCritical TaskPriority = "critical"
)

type Task struct {
	ID           uint64 `gorm:"primaryKey"`
	Title        string `gorm:"not null"`
	Detail       *string
	Status       TaskStatus   `gorm:"type:task_status;default:'pending'"`
	Priority     TaskPriority `gorm:"type:task_priority;default:'medium'"`
	DueDate      *time.Time
	UserID       uint64 `gorm:"not null"`
	AssigneeID   *uint64
	ParentTaskID *uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
