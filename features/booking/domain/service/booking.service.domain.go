package service

import (
	"time"

	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
)

func DaysOfRent(b *entity.Booking) int {
	start := time.Date(b.StartRent.Year(), b.StartRent.Month(), b.StartRent.Day(), 0, 0, 0, 0, b.StartRent.Location())
	end := time.Date(b.EndRent.Year(), b.EndRent.Month(), b.EndRent.Day(), 0, 0, 0, 0, b.EndRent.Location())

	d := int(end.Sub(start).Hours() / 24)
	if d <= 0 {
		return 0
	}
	return d
}

func CalculateDiscount(b *entity.Booking) float64 {
	if b.Customer != nil && b.Customer.Membership != nil {
		discountRate := b.Customer.Membership.DiscountRate
		b.Discount = b.TotalCost * (discountRate / 100)
		return b.Discount
	}
	b.Discount = 0
	return 0
}

func FinalCost(b *entity.Booking) float64 {
	return (b.TotalCost - b.Discount) + b.TotalDriverCost
}

func CalculateDriverCost(b *entity.Booking) float64 {
	if b.Driver != nil && b.Driver.DailyCost != nil {
		days := DaysOfRent(b)
		return float64(days) * *b.Driver.DailyCost
	}
	return 0
}
