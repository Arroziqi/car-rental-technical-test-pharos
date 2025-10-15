package seed

import (
	"github.com/Arroziqi/car-rental-technical-test-pharos.git/features/customer/domain/entity"
	"gorm.io/gorm"
)

func SeedMemberships(db *gorm.DB) {
	memberships := []entity.Membership{
		{Name: "Bronze", DiscountRate: 4.00},
		{Name: "Silver", DiscountRate: 7.00},
		{Name: "Gold", DiscountRate: 15.00},
	}

	for _, m := range memberships {
		db.FirstOrCreate(&m, entity.Membership{Name: m.Name})
	}
}
