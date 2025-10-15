package sql

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (r *CarRepository) Create(ctx context.Context, c *entity.Car) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *CarRepository) GetByID(ctx context.Context, id int) (*entity.Car, error) {
	var c entity.Car
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CarRepository) List(ctx context.Context) ([]*entity.Car, error) {
	var list []*entity.Car
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *CarRepository) Update(ctx context.Context, c *entity.Car) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *CarRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Car{}, id).Error
}
