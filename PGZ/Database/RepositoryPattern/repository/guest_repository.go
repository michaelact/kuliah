package repository

import (
	"repository_pattern/entity"
	"context"
)

type GuestRepository interface {
	Insert(ctx context.Context, guest entity.Guest) (entity.Guest, error)
	FindById(ctx context.Context, id int) (entity.Guest, error)
	FindAll(ctx context.Context) ([]entity.Guest, error)
}
