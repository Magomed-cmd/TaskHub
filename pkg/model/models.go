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
    ID          int64         // Уникальный идентификатор
    Title       string        // Название задачи
    Description *string       // Описание (опционально)
    Status      TaskStatus    // enum: pending, in_progress, completed, expired
    Priority    TaskPriority  // enum: low, medium, high, critical
    DueDate     *time.Time    // Дедлайн (может быть nil)
    CreatedAt   time.Time     // Метка создания
    UpdatedAt   time.Time     // Метка обновления
    UserID      int64         // Владелец задачи
    AssigneeID  *int64        // Кому назначена (может быть nil)
    Subtasks    []Task        // Вложенные задачи (один-ко-многим)
}
