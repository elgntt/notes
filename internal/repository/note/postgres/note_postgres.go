package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func (r *Repo) GetAllNotes(ctx context.Context) ([]model.NoteInfo, error) {
	rows, err := r.pool.Query(ctx,
		` SELECT 
     			id, 
     			text 
 			FROM notes`)

	if err != nil {
		return nil, fmt.Errorf(`SQL: get all notes:%w`, err)
	}

	var notes []model.NoteInfo

	for rows.Next() {
		note := &model.NoteInfo{}

		err := rows.Scan(&note.Id, &note.Text)
		if err != nil {
			return nil, fmt.Errorf(`SQL: get all notes:%w`, err)
		}

		notes = append(notes, *note)
	}

	return notes, nil
}

func (r *Repo) GetNote(ctx context.Context, noteId int) (*model.NoteInfo, error) {
	row := r.pool.QueryRow(ctx,
		` SELECT 
     			id,
     			text
			FROM notes
			WHERE id = $1`, noteId)

	var note model.NoteInfo

	err := row.Scan(&note.Id, &note.Text)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf(`SQL: get note:%w`, err)
	}

	return &note, nil
}
