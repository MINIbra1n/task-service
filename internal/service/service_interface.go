package service

import (
	"context"
	"task-service/internal/repo"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	CreateTask(ctx *fiber.Ctx) error
	GetTasks(ctx *fiber.Ctx) error
	GetTaskID(ctx *fiber.Ctx) error
	UpdateTaskID(ctx *fiber.Ctx) error
	DeleteTaskID(ctx *fiber.Ctx) error
}

type Storage interface {
	GetTasks(ctx context.Context) (*[]repo.Task, error)
	CreateTask(ctx context.Context, task repo.Task) (int64, error)
	GetTaskID(ctx context.Context, id int64) (*repo.Task, error)
	UpdateTaskID(ctx context.Context, id int64, task *repo.TaskUpdate) (int64, error)
	DeleteTaskID(ctx context.Context, id int64) (int64, error)
}
