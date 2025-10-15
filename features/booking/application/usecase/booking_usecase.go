package usecase

import (
	"context"
	"errors"
	"time"

	bookEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
	carEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
	custEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
)

type BookingRepository interface {
	Create(context.Context, *bookEntity.Booking) error
	GetByID(context.Context, int) (*bookEntity.Booking, error)
	List(context.Context) ([]*bookEntity.Booking, error)
	Update(context.Context, *bookEntity.Booking) error
	Delete(context.Context, int) error
}

type CarRepository interface {
	GetByID(context.Context, int) (*carEntity.Car, error)
	Update(context.Context, *carEntity.Car) error
}

type CustomerRepository interface {
	GetByID(context.Context, int) (*custEntity.Customer, error)
}

type BookingUsecase struct {
	bookingRepo BookingRepository
	carRepo     CarRepository
	custRepo    CustomerRepository
}

func NewBookingUsecase(b BookingRepository, cr CarRepository, crCust CustomerRepository) *BookingUsecase {
	return &BookingUsecase{
		bookingRepo: b,
		carRepo:     cr,
		custRepo:    crCust,
	}
}

func daysBetween(a, b time.Time) int {
	start := time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, a.Location())
	end := time.Date(b.Year(), b.Month(), b.Day(), 0, 0, 0, 0, b.Location())
	d := int(end.Sub(start).Hours() / 24)
	if d < 0 {
		return 0
	}
	if d == 0 {
		return 1 // treat same-day as 1 day rent
	}
	return d
}

func (u *BookingUsecase) Create(ctx context.Context, b *bookEntity.Booking) error {
	// validate customer
	if _, err := u.custRepo.GetByID(ctx, b.CustomerID); err != nil {
		return errors.New("customer not found")
	}

	// validate car
	car, err := u.carRepo.GetByID(ctx, b.CarID)
	if err != nil {
		return errors.New("car not found")
	}
	// check stock
	if car.Stock <= 0 {
		return errors.New("car out of stock")
	}

	// compute days and total cost
	days := daysBetween(b.StartRent, b.EndRent)
	base := float64(days) * car.DailyRent
	b.TotalCost = base

	// mark finished default false
	b.Finished = false

	// create booking
	if err := u.bookingRepo.Create(ctx, b); err != nil {
		return err
	}

	// reduce stock
	car.Stock = car.Stock - 1
	if err := u.carRepo.Update(ctx, car); err != nil {
		return err
	}

	return nil
}

func (u *BookingUsecase) GetByID(ctx context.Context, id int) (*bookEntity.Booking, error) {
	return u.bookingRepo.GetByID(ctx, id)
}

func (u *BookingUsecase) List(ctx context.Context) ([]*bookEntity.Booking, error) {
	return u.bookingRepo.List(ctx)
}

func (u *BookingUsecase) Update(ctx context.Context, b *bookEntity.Booking) error {
	return u.bookingRepo.Update(ctx, b)
}

func (u *BookingUsecase) Delete(ctx context.Context, id int) error {
	// when delete booking, optionally increment stock back (naive)
	b, err := u.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if err := u.bookingRepo.Delete(ctx, id); err != nil {
		return err
	}
	// increase stock
	car, err := u.carRepo.GetByID(ctx, b.CarID)
	if err == nil {
		car.Stock = car.Stock + 1
		_ = u.carRepo.Update(ctx, car)
	}
	return nil
}
