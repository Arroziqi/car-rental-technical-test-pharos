package repository

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
)

type CustomerRepository interface {
	Create(context.Context, *entity.Customer) error
	GetByID(context.Context, int) (*entity.Customer, error)
	List(context.Context) ([]*entity.Customer, error)
	Update(context.Context, *entity.Customer) error
	Delete(context.Context, int) error
}
