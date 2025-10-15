package entity

type Membership struct {
	ID           int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string     `json:"name" gorm:"type:varchar(50);not null;unique"`
	DiscountRate float64    `json:"discount_rate" gorm:"type:decimal(5,2);not null"`
	Customers    []Customer `json:"customers" gorm:"foreignKey:MembershipID"`
}
