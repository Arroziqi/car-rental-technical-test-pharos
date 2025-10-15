package repository_sql

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
	"gorm.io/gorm"
)

type CustomerSQLRepository struct {
	db *gorm.DB
}

func NewCustomerSQLRepository(db *gorm.DB) *CustomerSQLRepository {
	return &CustomerSQLRepository{db: db}
}

func (r *CustomerSQLRepository) Create(ctx context.Context, c *entity.Customer) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *CustomerSQLRepository) GetByID(ctx context.Context, id int) (*entity.Customer, error) {
	var c entity.Customer
	if err := r.db.WithContext(ctx).Preload("Membership").First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CustomerSQLRepository) List(ctx context.Context) ([]*entity.Customer, error) {
	var list []*entity.Customer
	if err := r.db.WithContext(ctx).Preload("Membership").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *CustomerSQLRepository) Update(ctx context.Context, c *entity.Customer) error {
	if err := r.db.WithContext(ctx).Save(c).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).
		Preload("Membership").
		First(c, c.ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *CustomerSQLRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Customer{}, id).Error
}
