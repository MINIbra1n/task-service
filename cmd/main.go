package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"task-service/internal/api/router"
	"task-service/internal/config"
	logging "task-service/internal/logger"
	"task-service/internal/repo"
	"task-service/internal/service"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env file", err)
	}
}

func main() {
	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(errors.Wrap(err, "failed to load configuration"))

	}
	loger, err := logging.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal(errors.Wrap(err, "No log level"))
	}

	repositories := repo.NewRepositories()
	service := service.NewService(repositories, loger)
	app := router.NewRouter(&router.Router{Service: service}, *service.GetStorageUser())
	go func() {
		log.Fatal(app.Listen(cfg.Rest.ServerName + cfg.Rest.ListenAddress))
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

}
