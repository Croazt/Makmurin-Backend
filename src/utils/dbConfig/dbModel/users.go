package dbModel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Post     []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;"`
}
