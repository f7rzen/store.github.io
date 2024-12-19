package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name        string  `gorm:"name" binding:"required" json:"name"`
    Description string  `gorm:"description" json:"description"`
    Price       float64 `gorm:"price" binding:"gte=0" json:"price"`
    ImageURL    string  `gorm:"image_url" json:"image_url"`
    CategoryID  uint    `gorm:"category_id" json:"category_id"`
}


