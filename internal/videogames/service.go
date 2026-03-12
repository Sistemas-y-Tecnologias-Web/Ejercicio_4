package videogames

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(ctx context.Context) ([]Videogame, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int) (Videogame, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, req CreateVideogame) (Videogame, error) {
	if err := req.Validate(); err != nil {
		return Videogame{}, err
	}
	return s.repo.Create(ctx, req)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, id int, req UpdateVideogame) (Videogame, error) {
	return s.repo.Update(ctx, id, req)
}
