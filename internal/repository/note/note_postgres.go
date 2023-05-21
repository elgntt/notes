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

func (r *Repo) UpdateNote(ctx context.Context, note model.Note) error {
	_, err := r.pool.Exec(ctx,
		` UPDATE notes 
 			SET text = $1
				WHERE id = $2`, note.Text, note.Id)
	if err != nil {
		return fmt.Errorf(`SQL: update note:%w`, err)
	}

	return nil
}

func (r *Repo) DeleteNote(ctx context.Context, noteId int) error {
	_, err := r.pool.Exec(ctx,
		` DELETE 
 			FROM notes
 			WHERE id = $1`, noteId)
	if err != nil {
		return fmt.Errorf(`SQL: delete note:%w`, err)
	}

	return nil
}
