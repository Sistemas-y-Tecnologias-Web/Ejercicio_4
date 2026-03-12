package videogames

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]Videogame, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM videogames")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var videogames []Videogame

	for rows.Next() {
		var v Videogame
		if err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.Category,
			&v.ActivePlayers,
			&v.Size,
			&v.Rating,
			&v.Downloads,
		); err != nil {
			return nil, err
		}
		videogames = append(videogames, v)
	}

	return videogames, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (Videogame, error) {
	var v Videogame
	err := r.db.QueryRow(ctx, "SELECT * FROM videogames WHERE id = $1", id).Scan(
		&v.ID,
		&v.Name,
		&v.Category,
		&v.ActivePlayers,
		&v.Size,
		&v.Rating,
		&v.Downloads,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return Videogame{}, ErrNotFound
	}
	return v, err
}

func (r *Repository) Create(ctx context.Context, req CreateVideogame) (Videogame, error) {
	var v Videogame
	err := r.db.QueryRow(ctx,
		`
	INSERT INTO videogames(name, category, active_players, size, rating, downloads) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, name, category, active_players, size, rating, downloads`,
		req.Name,
		req.Category,
		req.ActivePlayers,
		req.Size,
		req.Rating,
		req.Downloads,
	).Scan(
		&v.ID,
		&v.Name,
		&v.Category,
		&v.ActivePlayers,
		&v.Size,
		&v.Rating,
		&v.Downloads,
	)

	return v, err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, "DELETE FROM videogames WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil

}

func (r *Repository) Update(ctx context.Context, id int, req UpdateVideogame) (Videogame, error) {
	var v Videogame
	err := r.db.QueryRow(ctx,
		`
	UPDATE videogames
	SET name = $1, category = $2, active_players = $3, size = $4, rating = $5, downloads = $6
	WHERE id = $7
	RETURNING id, name, category, active_players, size, rating, downloads
	`,
		req.Name, req.Category, req.ActivePlayers, req.Size, req.Rating, req.Downloads, id,
	).Scan(
		&v.ID,
		&v.Name,
		&v.Category,
		&v.ActivePlayers,
		&v.Size,
		&v.Rating,
		&v.Downloads,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return Videogame{}, ErrNotFound
	}

	return v, err
}
