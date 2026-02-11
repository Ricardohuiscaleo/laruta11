package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct{ DB *sql.DB }

func main() {
	dsn := os.Getenv("APP_DB_USER") + ":" + os.Getenv("APP_DB_PASS") + "@tcp(" + os.Getenv("APP_DB_HOST") + ")/" + os.Getenv("APP_DB_NAME") + "?parseTime=true"
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()
	db.SetMaxOpenConns(25)

	s := &Server{DB: db}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-User")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/api/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	// Auth
	r.POST("/api/auth/login", s.authLogin)
	r.GET("/api/auth/check", s.authCheck)
	r.POST("/api/auth/logout", s.authLogout)

	// Compras
	r.GET("/api/compras", s.getCompras)
	r.GET("/api/compras/items", s.getComprasItems)
	r.GET("/api/compras/proveedores", s.getProveedores)
	r.GET("/api/compras/saldo", s.getSaldoDisponible)
	r.GET("/api/compras/historial-saldo", s.getHistorialSaldo)
	r.GET("/api/compras/precio-historico", s.getPrecioHistorico)
	r.POST("/api/compras", s.registrarCompra)
	r.DELETE("/api/compras/:id", s.deleteCompra)
	r.POST("/api/compras/:id/respaldo", s.uploadRespaldo)

	// Inventory
	r.GET("/api/ingredientes", s.getIngredientes)
	r.POST("/api/ingredientes", s.saveIngrediente)
	r.DELETE("/api/ingredientes/:id", s.deleteIngrediente)
	r.GET("/api/categories", s.getCategories)
	r.POST("/api/categories", s.saveCategory)
	r.DELETE("/api/categories/:id", s.deleteCategory)

	// Quality
	r.GET("/api/checklist", s.getChecklists)
	r.POST("/api/checklist", s.saveChecklist)
	r.DELETE("/api/checklist/:id", s.deleteChecklist)

	// Catalog
	r.GET("/api/products", s.getProducts)
	r.GET("/api/products/:id", s.getProductByID)

	// Orders
	r.GET("/api/orders/pending", s.getPendingOrders)
	r.POST("/api/orders/status", s.updateOrderStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Printf("🚀 API on :%s", port)
	r.Run(":" + port)
}
