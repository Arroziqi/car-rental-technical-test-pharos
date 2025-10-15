package entity

type Customer struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	NIK         string `json:"nik" gorm:"type:varchar(20);not null;unique"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20);not null"`
}
