package repository

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/driver/domain/entity"
)

type DriverRepository interface {
	Create(ctx context.Context, d *entity.Driver) error
	GetByID(ctx context.Context, id int) (*entity.Driver, error)
	List(ctx context.Context) ([]*entity.Driver, error)
	Update(ctx context.Context, d *entity.Driver) error
	Delete(ctx context.Context, id int) error
}
