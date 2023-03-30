package service

import (
	"TestTask/internal/repo"
)

type UseCase struct {
	PathService PathService
}

func NewUseCase(repo *repo.PathMongo, cache repo.Cache) *UseCase {
	return &UseCase{
		PathService: NewUserService(repo, cache),
	}
}
