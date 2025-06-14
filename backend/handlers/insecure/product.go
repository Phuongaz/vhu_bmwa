package insecure

import (
	"fmt"
	"net/http"
	"vhu_bmwa/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InsecureListProductsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query models.ProductQuery
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// VULNERABILITY: Direct SQL injection through string concatenation
		sqlQuery := "SELECT * FROM products WHERE 1=1"

		if query.Name != "" {
			// VULNERABILITY: SQL injection in LIKE clause
			sqlQuery += fmt.Sprintf(" AND name LIKE '%%%s%%'", query.Name)
		}

		if query.Description != "" {
			// VULNERABILITY: SQL injection in LIKE clause
			sqlQuery += fmt.Sprintf(" AND description LIKE '%%%s%%'", query.Description)
		}

		if query.Price > 0 {
			// VULNERABILITY: SQL injection in numeric comparison
			sqlQuery += fmt.Sprintf(" AND price = %f", query.Price)
		}

		if query.Sort != "" {
			// VULNERABILITY: SQL injection in ORDER BY clause
			sqlQuery += fmt.Sprintf(" ORDER BY %s", query.Sort)
		}

		if query.Filter != "" {
			// VULNERABILITY: Direct SQL injection through WHERE clause
			sqlQuery += " AND " + query.Filter
		}

		var products []models.InsecureProduct
		result := db.Raw(sqlQuery).Scan(&products)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error: " + result.Error.Error(),
				"query": sqlQuery, // VULNERABILITY: Exposing SQL query in error
			})
			return
		}

		// VULNERABILITY: No pagination, potential for large data exposure
		c.JSON(http.StatusOK, gin.H{
			"message":  "Products retrieved",
			"count":    result.RowsAffected,
			"products": products,
			"query_details": gin.H{
				"sql":     sqlQuery, // VULNERABILITY: Exposing SQL query
				"filters": query,
			},
		})
	}
}

// CreateProductHandler demonstrates multiple security vulnerabilities:
// 1. No authentication check
// 2. No input validation
// 3. XSS vulnerability
// 4. HTML injection
// 5. SQL injection
func InsecureCreateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.InsecureProduct
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		role := c.GetHeader("X-Role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to create products"})
			return
		}

		// VULNERABILITY: No price validation (allows negative prices)
		// VULNERABILITY: No XSS protection in text fields
		// VULNERABILITY: Allows HTML injection

		// VULNERABILITY: SQL injection through string concatenation
		query := "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"

		result := db.Exec(query, product.Name, product.Description, product.Price)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Failed to create product",
				"sql_error": result.Error.Error(), // VULNERABILITY: Exposing SQL error
				"query":     query,                // VULNERABILITY: Exposing SQL query
			})
			return
		}

		// Get the created product
		var createdProduct models.InsecureProduct
		db.Raw(fmt.Sprintf("SELECT * FROM products WHERE id = %d", result.RowsAffected)).Scan(&createdProduct)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product created",
		})
	}
}

// UpdateProductHandler demonstrates multiple security vulnerabilities:
// 1. No authentication check
// 2. No ownership validation
// 3. Mass assignment
// 4. SQL injection
// 5. No input validation
func InsecureUpdateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//check X-Role header
		role := c.GetHeader("X-Role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update products"})
			return
		}

		id := c.Param("id")

		var product models.InsecureProduct
		// VULNERABILITY: SQL injection in WHERE clause
		findQuery := fmt.Sprintf("SELECT * FROM products WHERE id = %s", id)
		result := db.Raw(findQuery).Scan(&product)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
				"query": findQuery, // VULNERABILITY: Exposing SQL query
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
				"id":    id,
				"query": findQuery, // VULNERABILITY: Exposing SQL query
			})
			return
		}

		// VULNERABILITY: Mass assignment vulnerability
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// VULNERABILITY: SQL injection through string concatenation
		updateQuery := fmt.Sprintf(
			"UPDATE products SET name = '%s', description = '%s', price = %f WHERE id = %s",
			product.Name,
			product.Description,
			product.Price,
			id,
		)

		updateResult := db.Exec(updateQuery)
		if updateResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Failed to update product",
				"sql_error": updateResult.Error.Error(), // VULNERABILITY: Exposing SQL error
				"query":     updateQuery,                // VULNERABILITY: Exposing SQL query
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product updated",
			"product": product,
		})
	}
}

// DeleteProductHandler demonstrates multiple security vulnerabilities:
// 1. No authentication check
// 2. No ownership validation
// 3. SQL injection
// 4. No soft delete
// 5. No record existence check
func InsecureDeleteProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//check X-Role header
		role := c.GetHeader("X-Role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete products"})
			return
		}

		id := c.Param("id")

		// VULNERABILITY: SQL injection through string concatenation
		query := fmt.Sprintf("DELETE FROM products WHERE id = %s", id)
		result := db.Exec(query)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Failed to delete product",
				"sql_error": result.Error.Error(), // VULNERABILITY: Exposing SQL error
				"query":     query,                // VULNERABILITY: Exposing SQL query
			})
			return
		}

		// VULNERABILITY: Information disclosure about record existence
		c.JSON(http.StatusOK, gin.H{
			"message":       "Product deletion attempted",
			"affected_rows": result.RowsAffected,
			"sql_details": gin.H{ // VULNERABILITY: Exposing implementation details
				"query":   query,
				"success": result.RowsAffected > 0,
			},
		})
	}
}
