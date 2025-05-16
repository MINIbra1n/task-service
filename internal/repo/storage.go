// package repo
package repo

// import (
// 	"context"
// 	"sync"
// 	"time"
// )

// var status = [3]string{"new", "in_progress", "done"}

// type tasks struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Status      string `json:"status"`
// 	Created_at  time.Time
// 	Updated_at  time.Time
// }
// type storage struct {
// 	Users map[string]string
// 	Tasks map[string]tasks
// 	mu    sync.Mutex
// }

// func NewRepositories() *storage {
// 	repo := storage{}
// 	repo.Users = make(map[string]string)
// 	repo.Tasks = make(map[string]tasks)
// 	repo.Users["admin"] = "admin"
// 	return &repo
// }

// func (r *storage) GetUsers() *map[string]string {
// 	return &r.Users
// }
// func (r *storage) GetStorageTask() *map[string]tasks {
// 	return &r.Tasks
// }
// func (r *storage) CreateTask(ctx context.Context, task *Task) (string, error) {
// 	repo := r.Tasks
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	if _, ok := repo[ask.ID]; !ok {
// 		if task.Status == status[0] || task.Status == status[1] || task.Status == status[2] {

// 			repo[task.ID] = tasks{
// 				Title:       task.Title,
// 				Description: task.Description,
// 				Status:      task.Status,
// 				Created_at:  task.Created_at,
// 				Updated_at:  task.Updated_at,
// 			}

// 			return task.ID, nil
// 		}
// 		return "", &FailedInsert{}
// 	}
// 	return "", &FailedInsert{}

// }

// func (r *storage) GetTasks(ctx context.Context) (*[]Task, error) {
// 	res := make([]Task, 0, 40)
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	if len(r.Tasks) == 0 {
// 		return nil, &EmtyStorage{}
// 	}
// 	for k, v := range r.Tasks {
// 		res = append(res, Task{
// 			ID:          k,
// 			Title:       v.Title,
// 			Description: v.Description,
// 			Status:      v.Status,
// 			Created_at:  v.Created_at,
// 			Updated_at:  v.Updated_at,
// 		})

// 	}
// 	return &res, nil
// }
// func (r *storage) GetTaskID(ctx context.Context, id string) (*Task, error) {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	if len(r.Tasks) == 0 {
// 		return nil, &EmtyStorage{}
// 	}
// 	if value, ok := r.Tasks[id]; ok {
// 		return &Task{
// 			ID:          id,
// 			Title:       value.Title,
// 			Description: value.Description,
// 			Status:      value.Status,
// 			Created_at:  value.Created_at,
// 			Updated_at:  value.Updated_at,
// 		}, nil
// 	} else {
// 		return nil, &NoSuchTaskStorage{}
// 	}
// }
// func (r *storage) UpdateTaskID(ctx context.Context, id string, task *TaskUpdate) (string, error) {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	if len(r.Tasks) == 0 {
// 		return "", &EmtyStorage{}
// 	}
// 	if task.Status == status[0] || task.Status == status[1] || task.Status == status[2] {
// 		if _, ok := r.Tasks[id]; ok {
// 			r.Tasks[id] = tasks{Title: task.Title, Description: task.Description, Status: task.Status, Updated_at: task.Updated_at}
// 			return id, nil
// 		}
// 		return "", &FailedToUpdate{}
// 	}
// 	return "", &FailedToUpdate{}
// }
// func (r *storage) DeleteTaskID(ctx context.Context, id string) (string, error) {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	if len(r.Tasks) == 0 {
// 		return "", &EmtyStorage{}
// 	}
// 	if _, ok := r.Tasks[id]; ok {
// 		delete(r.Tasks, id)
// 		return id, nil
// 	}
// 	return "", &NothingToDeleteById{}
// }
