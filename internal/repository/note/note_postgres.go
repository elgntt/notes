package note

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/elgntt/notes/internal/model"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		pool: db,
	}
}

func (r *Repo) CreateNote(ctx context.Context, note model.Note) error {
	_, err := r.pool.Exec(ctx,
		` INSERT INTO notes (text)
			VALUES ($1)`, note.Text)
	if err != nil {
		return fmt.Errorf(`SQL: create note:%w`, err)
	}

	return nil
}
