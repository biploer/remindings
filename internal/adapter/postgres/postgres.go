package postgres

import (
	"context"

	"github.com/biploer/remindings/internal/domain"
)

type Config struct {
	User     string
	Password string
	Port     string
	Host     string
	DBName   string
}

type Pool struct{}

func New(ctx context.Context, c Config) (*Pool, error) {
	// Делаем настройки подключения и пингуем БД на доступность
	return &Pool{}, nil
}

func (p *Pool) CreateProfile(ctx context.Context, profile domain.Profile) error {
	return nil
}

func (p *Pool) GetProfile(ctx context.Context, id int64) (domain.Profile, error) {
	return domain.Profile{}, nil
}

func (p *Pool) Close() {
	// Shutdown
}
