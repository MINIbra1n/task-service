package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	// "simple-service/internal/models"
	"task-service/internal/config"
)

// Слой репозитория, здесь должны быть все методы, связанные с базой данных

// SQL-запрос на вставку задачи
const (
	insertTaskQuery    = `INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id;`
	selectTaskQuery    = `SELECT title,description,status,created_at,updated_at FROM tasks WHERE id=$1`
	selectAllTaskQuery = `SELECT * FROM tasks;`
	updateTaskQuery    = `UPDATE tasks(title,description,status,updated_at) SET ($1,$2,$3,$4) WHERE id=$5`
	deleteTaskQuery    = `DELETE FROM tasks WHERE id=$1`
)

type repository struct {
	pool *pgxpool.Pool
}

// Repository - интерфейс с методом создания задачи
type Repository interface {
	GetTasks(ctx context.Context) (*[]Task, error)
	CreateTask(ctx context.Context, task Task) (int64, error)
	GetTaskID(ctx context.Context, id int64) (*Task, error)
	UpdateTaskID(ctx context.Context, id int64, task *TaskUpdate) (int64, error)
	DeleteTaskID(ctx context.Context, id int64) (int64, error)
}

// NewRepository - создание нового экземпляра репозитория с подключением к PostgreSQL
func NewRepository(ctx context.Context, cfg config.PostgreSQL) (Repository, error) {
	// Формируем строку подключения
	connString := fmt.Sprintf(
		`user=%s password=%s host=%s port=%d dbname=%s sslmode=%s 
        pool_max_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.PoolMaxConns,
		cfg.PoolMaxConnLifetime.String(),
		cfg.PoolMaxConnIdleTime.String(),
	)

	// Парсим конфигурацию подключения
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse PostgreSQL config")
	}

	// Оптимизация выполнения запросов (кеширование запросов)
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	// Создаём пул соединений с базой данных
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create PostgreSQL connection pool")
	}

	return &repository{pool}, nil
}

// CreateTask - вставка новой задачи в таблицу tasks
func (r *repository) CreateTask(ctx context.Context, task Task) (int64, error) {
	var id int64
	err := r.pool.QueryRow(ctx, insertTaskQuery, task.Title, task.Description).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert task")
	}
	return id, nil
}
func (r *repository) GetTaskID(ctx context.Context, id int64) (*Task, error) {
	resTask := Task{ID: id}
	err := r.pool.QueryRow(ctx, selectTaskQuery, id).Scan(&resTask.Title, &resTask.Description, &resTask.Status, &resTask.Created_at, &resTask.Updated_at)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select task id ")
	}
	return &resTask, nil
}
func (r *repository) GetTasks(ctx context.Context) (*[]Task, error) {
	resTasks := []Task{}
	task := Task{}
	row, err := r.pool.Query(ctx, selectAllTaskQuery)
	if err != nil {
		empty := &EmtyStorage{}
		return nil, empty
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&task)
		resTasks = append(resTasks, task)
	}
	return &resTasks, nil
}
func (r *repository) UpdateTaskID(ctx context.Context, id int64, task *TaskUpdate) (int64, error) {
	now := time.Now()
	_, err := r.pool.Exec(ctx, updateTaskQuery, task.Title, task.Description, now, id)
	if err != nil {
		return -1, errors.Wrap(err, "fail update task")
	}
	return id, nil
}
func (r *repository) DeleteTaskID(ctx context.Context, id int64) (int64, error) {
	_, err := r.pool.Exec(ctx, deleteTaskQuery, id)
	if err != nil {
		return -1, errors.Wrap(err, "fail delete task")
	}
	return id, nil
}
