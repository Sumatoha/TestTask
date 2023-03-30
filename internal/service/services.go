package service

import (
	"TestTask/internal/models"
	"TestTask/internal/repo"
	"context"
)

type PathService struct {
	pathRepo    repo.PathRepository
	cacheMemory repo.Cache
}

type Service interface {
	GetAllPaths(ctx context.Context) ([]models.Path, error)
	GetPathByID(ctx context.Context, id string) (models.Path, error)
	UpdatePath(ctx context.Context, id string, param models.Path) error
	DeletePath(ctx context.Context, id string) error
	AddPath(ctx context.Context, params models.Path) (string, error)
	RedirectPath(ctx context.Context, link string) (string, bool)
}

func NewUserService(pathRepo repo.PathRepository, cache repo.Cache) PathService {
	return PathService{
		pathRepo:    pathRepo,
		cacheMemory: cache,
	}
}

func (ps *PathService) GetAllPaths(ctx context.Context) ([]models.Path, error) {
	return ps.pathRepo.GetAllPaths(ctx)
}

func (ps *PathService) GetPathByID(ctx context.Context, id string) (models.Path, error) {
	paths, err := ps.pathRepo.GetPathByID(ctx, id)
	if err != nil {
		return models.Path{}, err
	}

	return paths, nil
}

func (ps *PathService) UpdatePath(ctx context.Context, id string, param models.Path) error {

	return nil
}

func (ps *PathService) DeletePath(ctx context.Context, id string) error {

	return nil
}

func (ps *PathService) AddPath(ctx context.Context, param models.Path) (string, error) {
	id, err := ps.pathRepo.AddPath(ctx, param)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ps *PathService) RedirectPath(ctx context.Context, link string) (string, bool) {
	//what is key? old or new link
	//example for map: [old_link]new_link

	activeLink, ok := ps.cacheMemory.Get(link)
	if !ok {
		path, iok := ps.pathRepo.GetPathByActiveLink(ctx, link)
		if !iok {
			hpath, dok := ps.pathRepo.GetPathByHistoryLink(ctx, link)
			if !dok {
				return "", false
			}
			return hpath.ActiveLink, false
		}
		return path.ActiveLink, true
	}
	return activeLink, false
}
