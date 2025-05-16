package validator

import (
	"errors"
	"task-service/internal/repo"
)

func ValidatorTask(task repo.Task) error {
	if task.ID < 0 {
		return errors.New("invalid id")
	}
	if len(task.Description) == 0 {
		return errors.New("invalid description")
	}
	if len(task.Status) == 0 {
		return errors.New("invalid Status")
	}
	if len(task.Title) == 0 {
		return errors.New("invalid Title")
	}
	return nil
}
