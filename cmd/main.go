package main

import (
	"os"
	"os/signal"
	"syscall"
	"task-service/internal/api/repo"
	"task-service/internal/api/router"
	"task-service/internal/service"
)

func main() {
	repositories := repo.NewRepositories()
	service := service.NewService(repositories)
	app := router.NewRouter(&router.Router{Service: service}, *service.GetStorageUser())
	go func() {
		app.Listen(":8080")
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

}
