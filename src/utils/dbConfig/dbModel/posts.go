package dbModel

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID      int    `json:"-"`
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Photo       string `json:"photo" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Description string `json:"desc" validate:"required"`
	Password    string `json:"password" `
}
