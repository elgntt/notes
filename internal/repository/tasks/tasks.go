package tasks

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		pool: db,
	}
}

func (r *Repo) Create(ctx context.Context, task dto.NewTask) (int, error) {
	query := `INSERT INTO tasks (title, description, status, category_id, project_id)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`

	var id int
	err := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Status,
		task.CategoryID,
		task.ProjectID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf(`SQL: CreateTask: Exec(): %w`, err)
	}

	return id, nil
}

func (r *Repo) Update(ctx context.Context, task domain.Task) error {
	query := `UPDATE tasks 
 			  SET title = $1,
 			      description = $2,
 			      status = $3,
 			      updated_at = DEFAULT
			  WHERE id = $4`

	_, err := r.pool.Exec(ctx, query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return fmt.Errorf(`SQL: UpdateTask: Exec(): %w`, err)
	}

	return nil
}

func (r *Repo) Delete(ctx context.Context, taskId int) error {
	query := `UPDATE tasks
			  SET deleted_at = CURRENT_TIMESTAMP
			  WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, taskId)
	if err != nil {
		return fmt.Errorf(`SQL: DeleteTask: Exec(): %w`, err)
	}

	return nil
}

func (r *Repo) Get(ctx context.Context, categoryID int) ([]domain.Task, error) {
	query := `SELECT 
     			t.id, 
     			t.title,
     			t.description,
     			t.status,
     			t.category_id,
     			t.project_id
 			 FROM tasks t
 			 JOIN categories c ON t.category_id = c.id
 			 WHERE t.deleted_at IS NULL
 			 AND c.deleted_at IS NULL
			 AND t.category_id = $1`

	rows, err := r.pool.Query(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf(`SQL: GetAllTasks: Query(): %w`, err)
	}

	var tasks []domain.Task

	for rows.Next() {
		task := domain.Task{}

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CategoryID,
			&task.ProjectID,
		)
		if err != nil {
			return nil, fmt.Errorf(`SQL: GetAllTasks: Scan(): %w`, err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *Repo) GetByID(ctx context.Context, taskId int) (*domain.Task, error) {
	query := `SELECT 
     			id,
     			title,
     			description,
     			status,
     			category_id,
     			project_id
			 FROM tasks
			 WHERE id = $1
			 AND deleted_at IS NULL`

	row := r.pool.QueryRow(ctx, query, taskId)

	var task domain.Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CategoryID,
		&task.ProjectID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf(`SQL: GetTask: Scan(): %w`, err)
	}

	return &task, nil
}
