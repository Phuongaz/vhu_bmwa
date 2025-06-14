package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	TotalPages int         `json:"total_pages"`
	Limit      int         `json:"limit"`
}

func ListProductsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get pagination parameters
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		var products []Product
		var total int64

		// Get total count
		db.Model(&Product{}).Count(&total)

		// Calculate offset
		offset := (page - 1) * limit

		// Get paginated products
		if result := db.Offset(offset).Limit(limit).Find(&products); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}

		// Calculate total pages
		totalPages := int(math.Ceil(float64(total) / float64(limit)))

		response := PaginatedResponse{
			Data:       products,
			Total:      total,
			Page:       page,
			TotalPages: totalPages,
			Limit:      limit,
		}

		c.JSON(http.StatusOK, response)
	}
}

func CreateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create products"})
			return
		}

		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Create(&product); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

func UpdateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update products"})
			return
		}

		id := c.Param("id")
		var product Product
		if result := db.First(&product, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Save(&product); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func DeleteProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete products"})
			return
		}

		id := c.Param("id")
		var product Product
		if result := db.First(&product, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if result := db.Delete(&product); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}
