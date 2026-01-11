package usecase

import (
	"context"

	"github.com/biploer/remindings/internal/adapter/postgres"
	"github.com/biploer/remindings/internal/domain"
)

type Postgres interface {
	CreateProfile(ctx context.Context, profile domain.Profile) error
	GetProfile(ctx context.Context, id int64) (domain.Profile, error)
}

type Profile struct {
	postgres Postgres
}

func NewProfile(postgres *postgres.Pool) *Profile {
	return &Profile{
		postgres: postgres,
	}
}
