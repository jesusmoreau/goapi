package service

import (
	"context"
	"goapi/internal/models"
	"goapi/internal/repository"
)

//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(r repository.Repository) Service {
	return &serv{repo: r}
}
