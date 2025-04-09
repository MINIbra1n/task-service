package service

import (
	"encoding/json"
	"errors"
	"task-service/internal/dto"
	"task-service/internal/repo"
	"time"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	repo Storage
}

func NewService(repo Storage) *service {
	return &service{
		repo: repo,
	}
}
func (s *service) GetStorageUser() *map[string]string {
	return s.repo.GetUsers()
}

// Пример запроса:
//
//	{
//	    "id":"1",
//	    "title":"first_title",
//	    "description":"first_description",
//	    "status":"new"
//	}
func (s *service) CreateTask(ctx *fiber.Ctx) error {
	var task repo.Task
	if err := json.Unmarshal(ctx.Body(), &task); err != nil {
		//loger
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}
	t := time.Now()
	task.Created_at = t
	task.Updated_at = t
	taskID, err := s.repo.CreateTask(ctx.Context(), &task)
	if err != nil {
		if errors.Is(&repo.FailedInsert{}, err) {
			dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
		}
		// s.log.Error("Failed to insert task", zap.Error(err))
		return dto.InternalServerError(ctx)
	}
	response := dto.Response{
		Status: "success",
		Data:   map[string]string{"task_id": taskID},
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
func (s *service) GetTasks(ctx *fiber.Ctx) error {

	tasks, err := s.repo.GetTasks(ctx.Context())
	if err != nil {
		if errors.Is(&repo.EmtyStorage{}, err) {
			dto.BadResponseError(ctx, string(fiber.StatusNotFound), "Empty storage")
		}
		//loger
		return dto.InternalServerError(ctx)
	}
	response := dto.Response{
		Status: "success",
		Data:   map[string][]repo.Task{"All_task": *tasks},
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *service) GetTaskID(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := s.repo.GetTaskID(c.Context(), id)
	if err != nil {
		if errors.Is(&repo.NoSuchTaskStorage{}, err) {
			dto.BadResponseError(c, string(fiber.StatusNotFound), "Not Found")
		}
		//loger
		return dto.InternalServerError(c)
	}
	response := dto.Response{
		Status: "success",
		Data:   map[string]repo.Task{"task_id": *task},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (s *service) UpdateTaskID(c *fiber.Ctx) error {
	var taskUp *repo.TaskUpdate
	id := c.Params("id")
	if err := json.Unmarshal(c.Body(), &taskUp); err != nil {
		//loger
		return dto.BadResponseError(c, dto.FieldBadFormat, "Invalid request body")
	}
	taskId, err := s.repo.UpdateTaskID(c.Context(), id, taskUp)
	if err != nil {
		//loger
		return dto.InternalServerError(c)
	}
	response := dto.Response{
		Status: "update",
		Data:   map[string]string{"task_id": taskId},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (s *service) DeleteTaskID(c *fiber.Ctx) error {
	id := c.Params("id")
	taskId, err := s.repo.DeleteTaskID(c.Context(), id)
	if err != nil {
		//loger
		return dto.InternalServerError(c)
	}
	response := dto.Response{
		Status: "update",
		Data:   map[string]string{"task_id": taskId},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
