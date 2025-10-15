package usecase

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
)

type CarRepository interface {
	Create(context.Context, *entity.Car) error
	GetByID(context.Context, int) (*entity.Car, error)
	List(context.Context) ([]*entity.Car, error)
	Update(context.Context, *entity.Car) error
	Delete(context.Context, int) error
}

type CarUsecase struct {
	repo CarRepository
}

func NewCarUsecase(r CarRepository) *CarUsecase {
	return &CarUsecase{repo: r}
}

func (u *CarUsecase) Create(ctx context.Context, c *entity.Car) error {
	return u.repo.Create(ctx, c)
}

func (u *CarUsecase) GetByID(ctx context.Context, id int) (*entity.Car, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *CarUsecase) List(ctx context.Context) ([]*entity.Car, error) {
	return u.repo.List(ctx)
}

func (u *CarUsecase) Update(ctx context.Context, c *entity.Car) error {
	return u.repo.Update(ctx, c)
}

func (u *CarUsecase) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
