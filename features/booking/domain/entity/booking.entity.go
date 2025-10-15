package entity

import (
	"time"

	carEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
	customerEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
)

// Booking represents a car booking/reservation.
type Booking struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID int       `json:"customer_id" gorm:"not null;index"`
	CarID      int       `json:"car_id" gorm:"not null;index"`
	StartRent  time.Time `json:"start_rent" gorm:"type:date;not null"`
	EndRent    time.Time `json:"end_rent" gorm:"type:date;not null"`
	TotalCost  float64   `json:"total_cost" gorm:"type:numeric(12,2)"`
	Finished   bool      `json:"finished" gorm:"default:false"`

	// Associations (pointers to avoid large copies and to allow nil)
	Customer *customerEntity.Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Car      *carEntity.Car           `json:"car,omitempty" gorm:"foreignKey:CarID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// DaysOfRent returns number of days for the booking (>=1)
// It uses truncation to date to avoid time-of-day issues.
func (b *Booking) DaysOfRent() int {
	start := time.Date(b.StartRent.Year(), b.StartRent.Month(), b.StartRent.Day(), 0, 0, 0, 0, b.StartRent.Location())
	end := time.Date(b.EndRent.Year(), b.EndRent.Month(), b.EndRent.Day(), 0, 0, 0, 0, b.EndRent.Location())

	// if end before start, return 0 (you can also validate earlier)
	d := int(end.Sub(start).Hours() / 24)
	if d <= 0 {
		return 0
	}
	return d
}
