package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	bookEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/booking/domain/entity"
	carEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/car/domain/entity"
	custEntity "github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
)

func NewPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&custEntity.Customer{},
		&carEntity.Car{},
		&bookEntity.Booking{},
	)
}
