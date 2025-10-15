package sql

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/driver/domain/entity"
	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/driver/domain/repository"
	"gorm.io/gorm"
)

type DriverSQLRepository struct {
	db *gorm.DB
}

func NewDriverSQLRepository(db *gorm.DB) repository.DriverRepository {
	return &DriverSQLRepository{db: db}
}

func (r *DriverSQLRepository) Create(ctx context.Context, d *entity.Driver) error {
	return r.db.WithContext(ctx).Create(d).Error
}

func (r *DriverSQLRepository) GetByID(ctx context.Context, id int) (*entity.Driver, error) {
	var d entity.Driver
	if err := r.db.WithContext(ctx).First(&d, id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DriverSQLRepository) List(ctx context.Context) ([]*entity.Driver, error) {
	var list []*entity.Driver
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *DriverSQLRepository) Update(ctx context.Context, d *entity.Driver) error {
	return r.db.WithContext(ctx).Save(d).Error
}

func (r *DriverSQLRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Driver{}, id).Error
}
