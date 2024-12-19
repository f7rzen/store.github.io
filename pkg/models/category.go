package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID       uint      `gorm:"id" gorm:"primaryKey;autoIncrement"`
	Name     string    `gorm:"name"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
