package sql

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Create(ctx context.Context, b *entity.Booking) error {
	return r.db.WithContext(ctx).Create(b).Error
}

func (r *BookingRepository) GetByID(ctx context.Context, id int) (*entity.Booking, error) {
	var b entity.Booking
	if err := r.db.WithContext(ctx).Preload("Customer").Preload("Car").First(&b, id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookingRepository) List(ctx context.Context) ([]*entity.Booking, error) {
	var list []*entity.Booking
	if err := r.db.WithContext(ctx).Preload("Customer").Preload("Car").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *BookingRepository) Update(ctx context.Context, b *entity.Booking) error {
	return r.db.WithContext(ctx).Save(b).Error
}

func (r *BookingRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Booking{}, id).Error
}
