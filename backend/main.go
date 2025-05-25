package main

import (
	"log"
	"os"
	"time"

	"vhu_bmwa/handlers"
	"vhu_bmwa/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&handlers.User{}, &handlers.Product{})

	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost", "https://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RateLimitMiddleware(10, time.Second))

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.RegHandler(db))
			auth.POST("/login", handlers.LoginHandler(db))
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", handlers.ProfHandler(db))
			protected.POST("/profile/change-password", handlers.ChgPwdHandler(db))

			products := protected.Group("/products")
			{
				products.GET("", handlers.ListProductsHandler(db))

				admin := products.Group("")
				admin.Use(middleware.RequireRole("admin"))
				{
					admin.POST("", handlers.CreateProductHandler(db))
					admin.PUT("/:id", handlers.UpdateProductHandler(db))
					admin.DELETE("/:id", handlers.DeleteProductHandler(db))
				}
			}
		}
	}

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `json:"-"`
	Role     string `gorm:"default:user"`
}

type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

func RateLimitMiddleware(maxRequests int, per time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
		c.Next()
	}
}

func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

func ProfileHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

func ChangePasswordHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

func ListProductsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

func CreateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

func UpdateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}

func DeleteProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
	}
}
