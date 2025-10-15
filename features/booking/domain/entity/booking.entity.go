package entity

import (
	"time"

	carEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
	customerEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
	driverEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/driver/domain/entity"
)

type Booking struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID      int       `json:"customer_id" gorm:"not null;index"`
	CarID           int       `json:"car_id" gorm:"not null;index"`
	StartRent       time.Time `json:"start_rent" gorm:"type:date;not null"`
	EndRent         time.Time `json:"end_rent" gorm:"type:date;not null"`
	TotalCost       float64   `json:"total_cost" gorm:"type:numeric(12,2)"`
	Discount        float64   `json:"discount" gorm:"type:numeric(12,2);default:0"`
	BookingTypeID   int       `json:"booking_type_id" gorm:"index"` // e.g. 1 = self-drive, 2 = with driver
	DriverID        *int      `json:"driver_id"`
	TotalDriverCost float64   `json:"total_driver_cost" gorm:"type:numeric(12,2);default:0"`
	Finished        bool      `json:"finished" gorm:"default:false"`

	// Relations
	Customer  *customerEntity.Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Car       *carEntity.Car           `json:"car,omitempty" gorm:"foreignKey:CarID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Driver    *driverEntity.Driver     `json:"driver,omitempty" gorm:"foreignKey:DriverID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time                `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time                `json:"updated_at" gorm:"autoUpdateTime"`
}
