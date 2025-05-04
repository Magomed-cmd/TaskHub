package handler

import (
	"TaskHub/internal/service"
	"TaskHub/pkg/model"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) Create(c *gin.Context) {

	var task model.Task

	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTask(ctx, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"object": task,
	})

}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	ctx := c.Request.Context()

	tasks, err := h.service.GetTasks(ctx)
	if err != nil {
		log.Printf("ERROR: failed to get tasks: %v", err) // ← лог ошибки в консоль
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) Delete(c *gin.Context) {

	ctx := c.Request.Context()
	id_s := c.Param("id")

	id, err := strconv.Atoi(id_s)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if err = h.service.DeleteTask(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {

	task := model.Task{}
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	task.ID = int64(id)

	if err != nil {
		log.Printf("error to convert id to int: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		log.Printf("error to get json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := h.service.UpdateTask(ctx, &task)
	if err != nil {
		log.Printf("error to update task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": newTask})
}
