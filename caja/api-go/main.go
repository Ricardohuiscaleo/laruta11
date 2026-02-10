package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db *sql.DB
}

func main() {
	dbHost := os.Getenv("APP_DB_HOST")
	dbName := os.Getenv("APP_DB_NAME")
	dbUser := os.Getenv("APP_DB_USER")
	dbPass := os.Getenv("APP_DB_PASS")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	server := &Server{db: db}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "caja-api"})
	})

	// Orders
	r.GET("/api/orders/pending", server.getPendingOrders)
	r.POST("/api/orders/status", server.updateOrderStatus)

	// Products
	r.GET("/api/products", server.getProducts)
	r.GET("/api/products/:id", server.getProductByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}

	log.Printf("🚀 Caja API running on port %s", port)
	r.Run(":" + port)
}
