package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id           string `gorm:"primary_key" json:"id"`
	Email        string `gorm:"type:varchar(40)" json:"email"`
	PhoneNumber  string `gorm:"type:varchar(40)" json:"phone_number"`
	ParcelWeight string `gorm:"type:varchar(40)" json:"parcel_weight"`
	Country      string `gorm:"type:varchar(40); index; not null" json:"country"`
}
