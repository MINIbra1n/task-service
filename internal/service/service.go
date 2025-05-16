package service

import (
	"encoding/json"
	"errors"
	"strconv"
	"task-service/internal/dto"
	"task-service/internal/repo"
	"task-service/pkg/validator"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type service struct {
	repo   Storage
	logger *zap.SugaredLogger
}

func NewService(repo Storage, logger *zap.SugaredLogger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// func (s *service) GetStorageUser() *map[string]string {
// 	return s.repo.GetUsers()
// }

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
		s.logger.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, 0, "Invalid request body")
	}

	if err := validator.ValidatorTask(task); err != nil {
		s.logger.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, 0, "Invalid request body")
	}
	t := time.Now()
	task.Created_at = t
	task.Updated_at = t
	taskID, err := s.repo.CreateTask(ctx.Context(), task)
	if err != nil {
		if errors.Is(&repo.FailedInsert{}, err) {

			return dto.BadResponseError(ctx, 400, "Invalid request body")
		}
		s.logger.Error("Failed to insert task", zap.Error(err))
		return dto.InternalServerError(ctx)
	}
	ID := strconv.Itoa(int(taskID))
	response := dto.Response{
		Status: "success",
		Data:   map[string]string{"task_id": (ID)},
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
func (s *service) GetTasks(ctx *fiber.Ctx) error {

	tasks, err := s.repo.GetTasks(ctx.Context())
	if err != nil {
		if errors.Is(&repo.EmtyStorage{}, err) {
			return ctx.Status(fiber.StatusNotFound).JSON(dto.Error{
				Code: fiber.StatusNotFound,
				Desc: "Empty storage",
			})

		}
		s.logger.Error("Failed to get task", zap.Error(err))
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
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	task, err := s.repo.GetTaskID(c.Context(), int64(intID))
	if err != nil {
		if errors.Is(&repo.EmtyStorage{}, err) {
			return c.Status(fiber.StatusNotFound).JSON(dto.Error{
				Code: fiber.StatusNotFound,
				Desc: "Empty storage",
			})
		} else if errors.Is(&repo.NoSuchTaskStorage{}, err) {
			return dto.BadResponseError(c, (fiber.StatusNotFound), err.Error())

		}
		s.logger.Error("Failed to selects task", zap.Error(err))
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
		return dto.BadResponseError(c, 0, "Invalid request body")
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	taskId, err := s.repo.UpdateTaskID(c.Context(), int64(intID), taskUp)
	if err != nil {
		if errors.Is(&repo.EmtyStorage{}, err) {

			return dto.BadResponseError(c, (fiber.StatusNotFound), err.Error())

		} else if errors.Is(&repo.FailedToUpdate{}, err) {
			return dto.BadResponseError(c, (fiber.StatusNotFound), err.Error())

		}
		s.logger.Error("Failed to update task", zap.Error(err))
		return dto.InternalServerError(c)
	}
	id = strconv.Itoa(int(taskId))
	response := dto.Response{
		Status: "update",
		Data:   map[string]string{"task_id": (id)},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (s *service) DeleteTaskID(c *fiber.Ctx) error {

	id := c.Params("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	taskId, err := s.repo.DeleteTaskID(c.Context(), int64(intID))
	if err != nil {
		if errors.Is(&repo.EmtyStorage{}, err) {
			return dto.BadResponseError(c, (fiber.StatusNotFound), err.Error())

		} else if errors.Is(&repo.NothingToDeleteById{}, err) {
			return dto.BadResponseError(c, (fiber.StatusNotFound), err.Error())

		}
		s.logger.Error("Failed to delete task", zap.Error(err))
		return dto.InternalServerError(c)
	}
	idtask := strconv.Itoa(int(taskId))
	response := dto.Response{
		Status: "delete",
		Data:   map[string]string{"task_id": (idtask)},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
