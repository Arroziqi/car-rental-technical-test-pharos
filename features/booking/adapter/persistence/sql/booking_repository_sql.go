package sql

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/repository"
	"gorm.io/gorm"
)

type BookingSQLRepository struct {
	db *gorm.DB
}

var _ repository.BookingRepository = (*BookingSQLRepository)(nil)

func NewBookingSQLRepository(db *gorm.DB) *BookingSQLRepository {
	return &BookingSQLRepository{db: db}
}

func (r *BookingSQLRepository) Create(ctx context.Context, b *entity.Booking) error {
	return r.db.WithContext(ctx).Create(b).Error
}

func (r *BookingSQLRepository) GetByID(ctx context.Context, id int) (*entity.Booking, error) {
	var b entity.Booking
	if err := r.db.WithContext(ctx).Preload("Customer.Membership").Preload("Car").Preload("Driver").First(&b, id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookingSQLRepository) List(ctx context.Context) ([]*entity.Booking, error) {
	var list []*entity.Booking
	if err := r.db.WithContext(ctx).Preload("Customer.Membership").Preload("Car").Preload("Driver").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *BookingSQLRepository) Update(ctx context.Context, b *entity.Booking) error {
	return r.db.WithContext(ctx).Save(b).Error
}

func (r *BookingSQLRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Booking{}, id).Error
}
