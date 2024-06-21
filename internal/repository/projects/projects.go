package projects

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

func (r *Repo) Create(ctx context.Context, project dto.Project) (domain.Project, error) {
	query := `
		INSERT INTO projects (name, description) 
		VALUES ($1, $2)
		RETURNING id, name, description`

	var proj domain.Project
	if err := r.pool.QueryRow(ctx, query, project.Name, project.Description).Scan(
		&proj.ID,
		&proj.Name,
		&proj.Description,
	); err != nil {
		return proj, fmt.Errorf("repo.Projects.Create: %w", err)
	}

	return proj, nil
}

func (r *Repo) GetByID(ctx context.Context, id int) (*domain.Project, error) {
	query := `
		SELECT id, name, description
		FROM projects
		WHERE id = $1
		AND deleted_at IS NULL`

	var proj domain.Project
	if err := r.pool.QueryRow(ctx, query, id).Scan(
		&proj.ID,
		&proj.Name,
		&proj.Description,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("repo.Projects.GetByID: %w", err)
	}

	return &proj, nil
}

func (r *Repo) Delete(ctx context.Context, id int) error {
	query := `
		UPDATE projects SET deleted_at = now()
		WHERE id = $1`

	if _, err := r.pool.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("repo.Projects.Delete: %w", err)
	}

	return nil
}

func (r *Repo) Update(ctx context.Context, project domain.Project) (domain.Project, error) {
	query := `
		UPDATE projects SET name = $1, description = $2
		WHERE id = $3
		AND deleted_at IS NULL
		RETURNING id, name, description`

	var proj domain.Project
	if err := r.pool.QueryRow(ctx, query, project.Name, project.Description, project.ID).Scan(
		&proj.ID,
		&proj.Name,
		&proj.Description,
	); err != nil {
		return domain.Project{}, fmt.Errorf("repo.Projects.Update: %w", err)
	}

	return proj, nil
}
