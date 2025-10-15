package sql

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(ctx context.Context, c *entity.Customer) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *CustomerRepository) GetByID(ctx context.Context, id int) (*entity.Customer, error) {
	var c entity.Customer
	if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CustomerRepository) List(ctx context.Context) ([]*entity.Customer, error) {
	var list []*entity.Customer
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *CustomerRepository) Update(ctx context.Context, c *entity.Customer) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *CustomerRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Customer{}, id).Error
}
