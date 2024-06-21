package categories

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

func (r *Repo) Create(ctx context.Context, draft dto.NewCategory) (domain.Category, error) {
	query := `
		INSERT INTO categories (title, project_id)
		VALUES ($1, $2)
		RETURNING id, title, project_id`

	var category domain.Category
	if err := r.pool.QueryRow(ctx, query, draft.Title, draft.ProjectID).Scan(
		&category.ID,
		&category.Title,
		&category.ProjectID,
	); err != nil {
		return domain.Category{}, fmt.Errorf("repo.Categories.Create: %w", err)
	}

	return category, nil
}

func (r *Repo) GetByID(ctx context.Context, id int) (*domain.Category, error) {
	query := `
		SELECT id, title, project_id
		FROM categories
		WHERE id = $1
		AND deleted_at IS NULL`

	var category domain.Category
	if err := r.pool.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Title,
		&category.ProjectID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("repo.Categories.GetByID: %w", err)
	}

	return &category, nil
}

func (r *Repo) FindByProjectID(ctx context.Context, projectId int) ([]domain.Category, error) {
	query := `
		SELECT c.id, c.title, c.project_id
		FROM categories c
		JOIN projects p ON p.id = c.project_id
		WHERE c.project_id = $1
		AND c.deleted_at IS NULL
		AND p.deleted_at IS NULL`

	var categories []domain.Category

	rows, err := r.pool.Query(ctx, query, projectId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category domain.Category

		if err := rows.Scan(
			&category.ID,
			&category.Title,
			&category.ProjectID,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo.Categories.FindByProjectID: %w", err)
	}

	return categories, nil
}

func (r *Repo) Delete(ctx context.Context, id int) error {
	query := `	
		UPDATE categories 
		SET deleted_at = now() 
		WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("repo.Categories.Delete: %w", err)
	}

	return nil
}

func (r *Repo) Update(ctx context.Context, draft domain.Category) (domain.Category, error) {
	query := `
		UPDATE categories
		SET title = $1
		WHERE id = $2
		RETURNING id, title, project_id`

	var category domain.Category
	if err := r.pool.QueryRow(ctx, query, draft.Title, draft.ID).Scan(
		&category.ID,
		&category.Title,
		&category.ProjectID,
	); err != nil {
		return domain.Category{}, fmt.Errorf("repo.Categories.Update: %w", err)
	}

	return category, nil
}
