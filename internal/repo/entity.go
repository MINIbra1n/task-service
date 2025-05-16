package repo

import "time"

type Task struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	Created_at  time.Time `json:"created_at" db:"created_at"`
	Updated_at  time.Time `json:"updated_at" db:"updated_at" `
}

type TaskUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`

	Updated_at time.Time
}

// type ResponseTask struct {
// 	ID          int64     `json:"id" db:"id"`
// 	Title       string    `json:"title" db:"title"`
// 	Description string    `json:"description" db:"description"`
// 	Status      string    `json:"status db:"status"`
// 	Created_at  time.Time `json:"created_at" db:"created_at"`
// 	Updated_at  time.Time `json:"updated_at" db:"updated_at" `
// }
