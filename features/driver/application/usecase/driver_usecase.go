package usecase

import (
	"context"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/driver/domain/entity"
	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/driver/domain/repository"
)

type DriverUsecase struct {
	repo repository.DriverRepository
}

func NewDriverUsecase(r repository.DriverRepository) *DriverUsecase {
	return &DriverUsecase{repo: r}
}

func (u *DriverUsecase) Create(ctx context.Context, d *entity.Driver) error {
	return u.repo.Create(ctx, d)
}

func (u *DriverUsecase) GetByID(ctx context.Context, id int) (*entity.Driver, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *DriverUsecase) List(ctx context.Context) ([]*entity.Driver, error) {
	return u.repo.List(ctx)
}

func (u *DriverUsecase) Update(ctx context.Context, d *entity.Driver) error {
	return u.repo.Update(ctx, d)
}

func (u *DriverUsecase) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
