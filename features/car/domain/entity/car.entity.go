package entity

type Car struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string  `json:"name" gorm:"type:varchar(100);not null"`
	Stock     int     `json:"stock" gorm:"not null;default:0"`
	DailyRent float64 `json:"daily_rent" gorm:"type:numeric(12,2);not null;default:0"`
}
