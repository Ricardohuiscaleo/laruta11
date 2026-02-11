package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Server struct{ DB *sql.DB }

func main() {
	godotenv.Load()
	dsn := os.Getenv("APP_DB_USER") + ":" + os.Getenv("APP_DB_PASS") + "@tcp(" + os.Getenv("APP_DB_HOST") + ")/" + os.Getenv("APP_DB_NAME") + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("⚠️  DB open error:", err, "- usando modo mock")
	}
	if db != nil && db.Ping() != nil {
		log.Println("⚠️  DB ping error - usando modo mock para desarrollo local")
		db = nil
	}
	if db != nil {
		defer db.Close()
		db.SetMaxOpenConns(25)
		log.Println("✅ DB conectada")
	}

	s := &Server{DB: db}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" {
			origin = "*"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-User,Authorization")
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
	r.GET("/api/auth/session", s.checkSession)
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

	// Dashboard (consolidado - 8 endpoints PHP → 1 endpoint Go con goroutines)
	r.GET("/api/dashboard", s.getDashboard)

	// TUU Payments
	r.GET("/api/tuu/transactions", s.getTUUTransactions)
	r.GET("/api/tuu/stream", s.streamTUUTransactions)
	r.GET("/api/dashboard/analytics", s.getDashboardAnalytics)
	r.GET("/api/dashboard/cards", s.getDashboardCards)
	r.GET("/api/sales/analytics", s.getSalesAnalytics)
	r.GET("/api/sales/month-comparison", s.getMonthComparison)
	r.GET("/api/sales/previous-month", s.getPreviousMonthSummary)
	r.GET("/api/financial/reports", s.getFinancialReports)

	// MenuApp
	r.GET("/api/menu", s.getMenu)
	r.POST("/api/products/:id/like", s.toggleLike)
	r.PUT("/api/products/:id/status", s.toggleProductStatus)
	r.POST("/api/orders", s.createOrderFull)
	r.GET("/api/orders/user/:user_id", s.getUserOrders)
	r.GET("/api/notifications", s.getNotifications)
	r.POST("/api/notifications/admin", s.notifyAdmin)
	r.GET("/api/trucks", s.getTrucks)
	r.POST("/api/trucks", s.getTrucks)
	r.GET("/api/trucks/status", s.getTruckStatus)
	r.GET("/api/trucks/schedules", s.getTruckSchedules)
	r.PUT("/api/trucks/:id", s.updateTruck)
	r.POST("/api/trucks/status", s.updateTruck)
	r.POST("/api/trucks/config", s.updateTruck)
	r.POST("/api/trucks/schedule", s.updateTruckSchedule)
	r.POST("/api/location/geocode", s.handleLocation)
	r.POST("/api/location/save", s.handleLocation)
	r.POST("/api/location/delivery-zone", s.handleLocation)
	r.POST("/api/location/products", s.handleLocation)
	r.POST("/api/location/delivery-time", s.handleLocation)
	r.PUT("/api/users/:id", s.updateUser)
	r.POST("/api/users/profile", s.updateUser)
	r.DELETE("/api/users/:id", s.deleteUser)
	r.POST("/api/track/usage", s.trackUsage)

	// Comandas & Tracking
	r.GET("/api/comandas", s.getComandas)
	r.GET("/api/comandas/realtime", s.realtimeComandas)
	r.PUT("/api/comandas/:id/status", s.updateComandaStatus)
	r.POST("/api/track/visit", s.trackVisit)
	r.POST("/api/track/interaction", s.trackInteraction)

	// Arqueo
	r.GET("/api/get_sales_summary.php", s.getSalesSummary)
	r.GET("/api/get_saldo_caja.php", s.getSaldoCaja)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Printf("🚀 API on :%s", port)
	r.Run(":" + port)
}
