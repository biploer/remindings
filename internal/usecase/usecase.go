package usecase

import (
	"context"

	"github.com/biploer/remindings/internal/domain"
)

type Repository interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
}

type Profile struct {
	repository Repository
}

func NewProfile(repo Repository) *Profile {
	return &Profile{
		repository: repo,
	}
}
