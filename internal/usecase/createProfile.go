package usecase

import (
	"context"
	"fmt"

	"github.com/biploer/remindings/internal/domain"
	"github.com/biploer/remindings/internal/dto"
)

func (p *Profile) CreateProfile(ctx context.Context, input dto.CreateProfileInput) (dto.CreateProfileOutput, error) {
	var output dto.CreateProfileOutput

	// TODO: Проверить в кэше существование профиля

	profile, err := domain.NewProfile(input.TGId, input.FirstName, input.LastName, input.Username)
	if err != nil {
		return output, fmt.Errorf("create profile model: %w", err)
	}

	err = p.repository.CreateProfile(ctx, profile)
	if err != nil {
		return output, fmt.Errorf("create pofile in db: %w", err)
	}

	return dto.CreateProfileOutput{
		ID: profile.ID,
	}, nil
}
