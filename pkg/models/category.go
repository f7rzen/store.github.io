package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
