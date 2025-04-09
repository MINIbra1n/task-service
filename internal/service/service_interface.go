package service

import (
	"context"
	"task-service/internal/api/repo"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	CreateTask(ctx *fiber.Ctx) error
	GetTasks(ctx *fiber.Ctx) error
	GetTaskID(ctx *fiber.Ctx) error
	UpdateTaskID(ctx *fiber.Ctx) error
	DeleteTaskID(ctx *fiber.Ctx) error
	GetStorageUser() *map[string]string
}

type Storage interface {
	GetUsers() *map[string]string
	GetTasks(ctx context.Context) (*[]repo.Task, error)
	CreateTask(ctx context.Context, task *repo.Task) (string, error)
	GetTaskID(ctx context.Context, id string) (*repo.Task, error)
	UpdateTaskID(ctx context.Context, id string, task *repo.TaskUpdate) (string, error)
	DeleteTaskID(ctx context.Context, id string) (string, error)
}
