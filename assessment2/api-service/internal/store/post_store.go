package store

import (
	"api-service/internal/domain"
	"context"
)

// PostStore represent interface for blog storage
type PostStore interface {
	Get(ctx context.Context, title string, page int, limit int) (*[]domain.Post, error)
	GetOne(ctx context.Context, id int) (*domain.Post, error)
	Insert(ctx context.Context, post domain.Post) (int, error)
	Update(ctx context.Context, id int, post domain.Post) error
	Delete(ctx context.Context, id int) error
}
