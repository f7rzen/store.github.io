package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    id          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
    Price       float64 `json:"price" binding:"gte=0"`
    ImageURL    string  `json:"image_url"`
    CategoryID  uint    `json:"category_id"`
}

