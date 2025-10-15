package usecase

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/repository"
)

type CustomerUsecase struct {
	repo repository.CustomerRepository
}

func NewCustomerUsecase(r repository.CustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{repo: r}
}

func (u *CustomerUsecase) Create(ctx context.Context, c *entity.Customer) error {
	return u.repo.Create(ctx, c)
}

func (u *CustomerUsecase) GetByID(ctx context.Context, id int) (*entity.Customer, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *CustomerUsecase) List(ctx context.Context) ([]*entity.Customer, error) {
	return u.repo.List(ctx)
}

func (u *CustomerUsecase) Update(ctx context.Context, c *entity.Customer) error {
	return u.repo.Update(ctx, c)
}

func (u *CustomerUsecase) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
