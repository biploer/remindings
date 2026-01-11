package dto

import "github.com/google/uuid"

type CreateProfileOutput struct {
	ID uuid.UUID
}

type CreateProfileInput struct {
	TGId      int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}
