package videogames

import "errors"

var (
	ErrNotFound         = errors.New("videogame not found")
	ErrNameRequired     = errors.New("name of videogame required")
	ErrSizeRequired     = errors.New("size of the game required")
	ErrCategoryRequired = errors.New("category of game required")
)
