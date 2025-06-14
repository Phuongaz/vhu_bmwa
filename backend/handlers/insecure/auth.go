package insecure

import (
	"fmt"
	"net/http"
	"vhu_bmwa/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterHandler demonstrates multiple security vulnerabilities:
// 1. No input validation
// 2. Plain text password storage
// 3. SQL injection vulnerability
// 4. Role manipulation
func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Set default role if not provided
		if req.Role == "" {
			req.Role = "user"
		}

		// VULNERABILITY: SQL Injection through string concatenation
		// Instead of using parameterized queries
		createTableQuery := `
		CREATE TABLE IF NOT EXISTS insecure_users (
			id bigint unsigned AUTO_INCREMENT PRIMARY KEY,
			username varchar(255) NOT NULL,
			password varchar(255) NOT NULL,
			role varchar(50) NOT NULL DEFAULT 'user'
		)`

		if err := db.Exec(createTableQuery).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to ensure table exists",
				"details": err.Error(),
			})
			return
		}

		// VULNERABILITY: SQL Injection through string concatenation
		insertQuery := fmt.Sprintf(
			"INSERT INTO insecure_users (username, password, role) VALUES ('%s', '%s', '%s')",
			req.Username,
			req.Password, // VULNERABILITY: Storing password in plain text
			req.Role,     // VULNERABILITY: Allowing role selection
		)

		result := db.Exec(insertQuery)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Failed to create user",
				"sql_error": result.Error.Error(), // VULNERABILITY: Exposing SQL error
				"query":     insertQuery,          // VULNERABILITY: Exposing SQL query
			})
			return
		}

		// VULNERABILITY: Returning sensitive data
		c.JSON(http.StatusOK, gin.H{
			"ID":       result.RowsAffected,
			"username": req.Username,
			"password": req.Password, // VULNERABILITY: Returning password
			"role":     req.Role,
		})
	}
}

// LoginHandler demonstrates multiple security vulnerabilities:
// 1. SQL injection
// 2. Plain text password comparison
// 3. Weak session management
func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// VULNERABILITY: SQL Injection through string concatenation
		query := fmt.Sprintf(
			"SELECT * FROM insecure_users WHERE username = '%s' AND password = '%s'",
			req.Username,
			req.Password, // VULNERABILITY: Plain text password comparison
		)

		var user models.InsecureUser
		result := db.Raw(query).Scan(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Login failed",
				"sql_error": result.Error.Error(), // VULNERABILITY: Exposing SQL error
				"query":     query,                // VULNERABILITY: Exposing SQL query
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// VULNERABILITY: Using a weak secret and exposing token
		// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 	"user_id": user.ID,
		// 	"role":    user.Role,
		// })

		// // VULNERABILITY: Using a weak secret
		// tokenString, err := token.SignedString([]byte("weak_secret"))
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		// 	return
		// }

		// VULNERABILITY: Returning sensitive data and token in response
		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"ID":       user.ID,
				"username": user.Username,
				"password": user.Password, // VULNERABILITY: Returning password
				"role":     user.Role,
			},
		})
	}
}

// LogoutHandler demonstrates security vulnerability:
// 1. No token invalidation
func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// VULNERABILITY: No token invalidation, just returning success
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})
	}
}

// ChangePasswordHandler demonstrates multiple security vulnerabilities:
// 1. SQL injection
// 2. Plain text password storage
// 3. No authentication check
func InsecureChangePasswordHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//check X-User header
		username := c.GetHeader("X-User")
		if username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		var req models.ChangePasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// VULNERABILITY: SQL Injection through string concatenation
		query := fmt.Sprintf(
			"UPDATE insecure_users SET password = '%s' WHERE username = '%s'",
			req.NewPassword, // VULNERABILITY: Storing new password in plain text
			req.Username,    // VULNERABILITY: No authentication check
		)

		result := db.Exec(query)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Failed to change password",
				"sql_error": result.Error.Error(), // VULNERABILITY: Exposing SQL error
				"query":     query,                // VULNERABILITY: Exposing SQL query
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or current password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Password changed successfully",
			"debug": gin.H{ // VULNERABILITY: Exposing sensitive debug information
				"old_password": req.CurrentPassword,
				"new_password": req.NewPassword,
				"username":     req.Username,
			},
		})
	}
}

// ProfileHandler demonstrates security vulnerability:
// 1. No authentication check
// 2. SQL injection
func ProfileHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//check X-User header
		username := c.GetHeader("X-User")
		if username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// VULNERABILITY: SQL Injection through string concatenation
		query := fmt.Sprintf(
			"SELECT * FROM insecure_users WHERE username = '%s'",
			username,
		)

		var user models.InsecureUser
		result := db.Raw(query).Scan(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":     "Failed to get profile",
				"sql_error": result.Error.Error(), // VULNERABILITY: Exposing SQL error
				"query":     query,                // VULNERABILITY: Exposing SQL query
			})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// VULNERABILITY: Returning sensitive data
		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"ID":       user.ID,
				"username": user.Username,
				"password": user.Password, // VULNERABILITY: Returning password
				"role":     user.Role,
			},
		})
	}
}
