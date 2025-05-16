package validator

// import (
// 	"errors"
// 	"task-service/internal/repo"
// 	"testing"
// )

// func TestValidator(t *testing.T) {
// 	var test = []struct {
// 		task repo.Task
// 		want error
// 	}{
// 		{task: repo.Task{
// 			ID:          0,
// 			Title:       "",
// 			Description: "",
// 			Status:      ""},
// 			want: errors.New("")},
// 	}
// 	var arrTask = []repo.Task{
// 		{
// 			ID:          0,
// 			Title:       "",
// 			Description: "",
// 			Status:      "",
// 		},
// 		{
// 			ID:          1,
// 			Title:       "",
// 			Description: "",
// 			Status:      "",
// 		},
// 		{
// 			ID:          1,
// 			Title:       "title",
// 			Description: "",
// 			Status:      "",
// 		},
// 		{
// 			ID:          1,
// 			Title:       "title",
// 			Description: "desc",
// 			Status:      "",
// 		},
// 		{
// 			ID:          2,
// 			Title:       "title",
// 			Description: "desc",
// 			Status:      "status",
// 		},
// 	}

// }
