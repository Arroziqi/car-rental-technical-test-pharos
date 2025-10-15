package usecase

import (
	"context"
	"errors"

	bookEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
	bookRepo "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/repository"
	bookingService "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/service"
	carEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
	custEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
)

type CarRepository interface {
	GetByID(context.Context, int) (*carEntity.Car, error)
	Update(context.Context, *carEntity.Car) error
}

type CustomerRepository interface {
	GetByID(context.Context, int) (*custEntity.Customer, error)
}

type BookingUsecase struct {
	bookingRepo bookRepo.BookingRepository
	carRepo     CarRepository
	custRepo    CustomerRepository
}

func NewBookingUsecase(b bookRepo.BookingRepository, cr CarRepository, crCust CustomerRepository) *BookingUsecase {
	return &BookingUsecase{
		bookingRepo: b,
		carRepo:     cr,
		custRepo:    crCust,
	}
}

func (u *BookingUsecase) Create(ctx context.Context, b *bookEntity.Booking) error {
	// validate customer
	cust, err := u.custRepo.GetByID(ctx, b.CustomerID)
	if err != nil {
		return errors.New("customer not found")
	}

	// validate car
	car, err := u.carRepo.GetByID(ctx, b.CarID)
	if err != nil {
		return errors.New("car not found")
	}
	if car.Stock <= 0 {
		return errors.New("car out of stock")
	}

	// attach customer for discount calculation
	b.Customer = cust

	// compute days and total cost
	days := bookingService.DaysOfRent(b)
	b.TotalCost = float64(days) * car.DailyRent
	b.Finished = false

	// calculate discount via domain service
	bookingService.CalculateDiscount(b)

	// calculate driver cost safely
	if b.DriverID != nil && b.Driver != nil && b.Driver.DailyCost != nil {
		b.TotalDriverCost = bookingService.CalculateDriverCost(b)
	} else {
		b.TotalDriverCost = 0
	}

	// create booking
	if err := u.bookingRepo.Create(ctx, b); err != nil {
		return err
	}

	// reduce car stock
	car.Stock--
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
	// attach customer jika belum ada
	if b.Customer == nil {
		cust, err := u.custRepo.GetByID(ctx, b.CustomerID)
		if err == nil {
			b.Customer = cust
		}
	}

	// recompute discount & driver cost
	bookingService.CalculateDiscount(b)
	if b.DriverID != nil && b.Driver != nil && b.Driver.DailyCost != nil {
		b.TotalDriverCost = bookingService.CalculateDriverCost(b)
	} else {
		b.TotalDriverCost = 0
	}

	return u.bookingRepo.Update(ctx, b)
}

func (u *BookingUsecase) Delete(ctx context.Context, id int) error {
	b, err := u.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if err := u.bookingRepo.Delete(ctx, id); err != nil {
		return err
	}

	// restore car stock
	car, err := u.carRepo.GetByID(ctx, b.CarID)
	if err == nil {
		car.Stock++
		_ = u.carRepo.Update(ctx, car)
	}

	return nil
}
