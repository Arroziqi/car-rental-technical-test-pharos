package repository

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
)

type BookingRepository interface {
	Create(ctx context.Context, b *entity.Booking) error
	GetByID(ctx context.Context, id int) (*entity.Booking, error)
	List(ctx context.Context) ([]*entity.Booking, error)
	Update(ctx context.Context, b *entity.Booking) error
	Delete(ctx context.Context, id int) error
}
