package models

import (
	"gorm.io/gorm"
)

// Product represents a secure product model with proper validation
type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0" gorm:"not null"`
}

// InsecureProduct demonstrates common security anti-patterns:
// 1. No input validation
// 2. No price validation
// 3. Exposing internal fields
type InsecureProduct struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ProductQuery represents search parameters with SQL injection risks
type ProductQuery struct {
	Name        string  `form:"name"`
	Description string  `form:"description"`
	Price       float64 `form:"price"`
	Sort        string  `form:"sort"`   // Potential SQL injection in ORDER BY
	Filter      string  `form:"filter"` // Raw SQL WHERE clause injection
}

// ProductCreate represents the data needed to create a product
type ProductCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

// ProductUpdate represents the data that can be updated
type ProductUpdate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
