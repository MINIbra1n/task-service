package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	logging "task-service/internal/logger"
	"task-service/internal/repo"
	"task-service/internal/repo/mocks"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateTast(t *testing.T) {
	app := fiber.New(fiber.Config{})
	l, _ := logging.NewLogger("debug")

	r := mocks.NewRepository(t)
	s := NewService(r, l)

	app.Post("/task", s.CreateTask)
	t.Run("/task . Добовление задачи", func(t *testing.T) {
		r.On("CreateTask", mock.Anything, mock.AnythingOfType("repo.Task")).Return(int64(1), nil)
		req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(`{"id":100,"title":"TestTask","description":"first_description","status":"new"}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("/task", func(t *testing.T) {
		// r.On("CreateTask", mock.Anything, mock.AnythingOfType("repo.Task")).Return(int64(1), nil)
		req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(`{"id":0,"title":"","description":"","status":""}`))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		fmt.Println(resp.StatusCode)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

}

func TestGetTasks(t *testing.T) {
	mockDB := mocks.NewRepository(t)
	logger, _ := logging.NewLogger("debug")
	service := NewService(mockDB, logger)

	app := fiber.New()
	app.Get("/tasks", service.GetTasks)

	t.Run("success case", func(t *testing.T) {
		tasks := []repo.Task{
			{ID: 1, Title: "Task 1", Description: "Description 1", Status: "new"},
			{ID: 2, Title: "Task 2", Description: "Description 2", Status: "new"},
		}

		mockDB.On("GetTasks", mock.Anything).Return(&tasks, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	})
	t.Run("empty Db case", func(t *testing.T) {
		emptyError := &repo.EmtyStorage{}
		mockDB.On("GetTasks", mock.Anything).Return(nil, emptyError).Once()

		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

func TestGetTaskID(t *testing.T) {
	mockDB := mocks.NewRepository(t)
	logger, _ := logging.NewLogger("debug")
	service := NewService(mockDB, logger)

	app := fiber.New()
	app.Get("/task/:id", service.GetTaskID)

	t.Run("succes case", func(t *testing.T) {
		task := &repo.Task{ID: 1, Title: "Task 1", Description: "Description 1", Status: "new"}
		mockDB.On("GetTaskID", mock.Anything, task.ID).Return(task, nil)
		req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
	t.Run("empty Db case", func(t *testing.T) {
		emptyError := &repo.EmtyStorage{}
		task := &repo.Task{ID: 1, Title: "Task 1", Description: "Description 1", Status: "new"}

		mockDB.On("GetTaskID", mock.Anything, task.ID).Return(nil, emptyError).Once()

		req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

func TestUpdateTaskID(t *testing.T) {
	mockDB := mocks.NewRepository(t)
	logger, _ := logging.NewLogger("debug")
	service := NewService(mockDB, logger)

	app := fiber.New()
	app.Put("/task/:id", service.UpdateTaskID)

	t.Run("succes case", func(t *testing.T) {
		// task := &repo.Task{ID: 1, Title: "Task 1", Description: "Description 1", Status: "new"}
		// task := &repo.TaskUpdate{Title: "Task 1", Description: "Description 1", Status: "new"}
		id := 1
		updateData := repo.TaskUpdate{
			Title:       "Updated Task",
			Description: "Updated Description",
			Status:      "done",
		}

		// Преобразуем данные в JSON для запроса
		updateJSON, err := json.Marshal(updateData)
		require.NoError(t, err)
		mockDB.On("UpdateTaskID", mock.Anything, int64(id), mock.AnythingOfType("*repo.TaskUpdate")).Return(int64(1), nil)
		req := httptest.NewRequest(http.MethodPut, "/task/1", bytes.NewReader(updateJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
