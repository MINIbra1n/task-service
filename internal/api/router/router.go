package router

import (
	"task-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Repositories interface {
	GetUsers() *map[string]string
}
type Router struct {
	Service service.Service
}

func NewRouter(rout *Router) *fiber.App {
	r := fiber.New(fiber.Config{})

	rGroupV1 := r.Group("/v1")
	//  Handler
	rGroupV1.Post("/task", rout.Service.CreateTask)
	rGroupV1.Get("/tasks", rout.Service.GetTasks)
	rGroupV1.Get("/task/:id", rout.Service.GetTaskID)
	rGroupV1.Put("/task/:id", rout.Service.UpdateTaskID)
	rGroupV1.Delete("/task/:id", rout.Service.DeleteTaskID)

	return r

}
