package videogames

import "errors"

var (
	ErrNotFound         = errors.New("Videogame not found.")
	ErrNameRequired     = errors.New("Name of videogame required.")
	ErrSizeRequired     = errors.New("Size of the game required.")
	ErrCategoryRequired = errors.New("Category of game required.")
)
