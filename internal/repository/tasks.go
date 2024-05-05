package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/elgntt/notes/internal/model/domain"
	"github.com/elgntt/notes/internal/model/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		pool: db,
	}
}

func (r *Repo) Create(ctx context.Context, task dto.NewTask) error {
	query := `INSERT INTO tasks (title, description, due_date, status)
			  VALUES ($1, $2, TO_TIMESTAMP($3, 'YYYY-MM-DD HH24:MI:SS'), $4)`

	_, err := r.pool.Exec(ctx, query, task.Title, task.Description, task.DueDate, task.Status)
	if err != nil {
		return fmt.Errorf(`SQL: CreateTask: Exec(): %w`, err)
	}

	return nil
}

func (r *Repo) Update(ctx context.Context, task domain.Task) error {
	query := `UPDATE tasks 
 			  SET title = $1,
 			      description = $2,
 			      due_date = $3,
 			      status = $4,
 			      updated_at = DEFAULT
			  WHERE id = $5`

	_, err := r.pool.Exec(ctx, query, task.Title, task.Description, task.DueDate, task.Status, task.ID)
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

func (r *Repo) GetAll(ctx context.Context) ([]domain.Task, error) {
	query := `SELECT 
     			id, 
     			title,
     			description,
     			due_date,
     			status
 			 FROM tasks
 			 WHERE deleted_at IS NULL`

	rows, err := r.pool.Query(ctx, query)
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
			&task.DueDate,
			&task.Status,
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
     			due_date,
     			status
			 FROM tasks
			 WHERE id = $1
			 AND deleted_at IS NULL`

	row := r.pool.QueryRow(ctx, query, taskId)

	var task domain.Task

	var dueDate pgtype.Timestamptz
	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&dueDate,
		&task.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf(`SQL: GetTask: Scan(): %w`, err)
	}

	if !dueDate.Time.IsZero() {
		task.DueDate = &dueDate.Time
	}

	return &task, nil
}
