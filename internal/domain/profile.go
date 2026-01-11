package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID
	TGId      int64
	CreatedAt time.Time
	FirstName string
	LastName  string
	Username  string
}

func NewProfile(tgId int64, firstName, lastName, username string) (Profile, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return Profile{}, fmt.Errorf("create uuid.NewV7: %w", err)
	}

	return Profile{
		ID:        newId,
		TGId:      tgId,
		CreatedAt: time.Now(),
	}, nil
}
