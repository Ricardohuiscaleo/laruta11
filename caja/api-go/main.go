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

	// Dashboard (consolidado - 8 endpoints PHP → 1 endpoint Go con goroutines)
	r.GET("/api/dashboard", s.getDashboard)

	// TUU Payments
	r.GET("/api/tuu/get_from_mysql.php", s.getTUUTransactions)
	r.GET("/api/tuu/stream", s.streamTUUTransactions)
	r.GET("/api/get_dashboard_analytics.php", s.getDashboardAnalytics)
	r.GET("/api/get_dashboard_cards.php", s.getDashboardCards)
	r.GET("/api/get_sales_analytics.php", s.getSalesAnalytics)
	r.GET("/api/get_month_comparison.php", s.getMonthComparison)
	r.GET("/api/get_previous_month_summary.php", s.getPreviousMonthSummary)
	r.GET("/api/get_financial_reports.php", s.getFinancialReports)

	// MenuApp (22 PHP → 8 Go consolidados)
	r.GET("/api/menu", s.getMenu)
	r.GET("/api/get_menu_products.php", s.getMenu) // Alias legacy
	r.POST("/api/products/:id/like", s.toggleLike)
	r.POST("/api/toggle_like.php", s.toggleLike) // Alias legacy
	r.PUT("/api/products/:id/status", s.toggleProductStatus)
	r.POST("/api/toggle_product_status.php", s.toggleProductStatus) // Alias legacy
	r.POST("/api/orders", s.createOrderFull)
	r.POST("/api/create_order.php", s.createOrderFull) // Alias legacy
	r.GET("/api/orders/user/:user_id", s.getUserOrders)
	r.GET("/api/get_user_orders.php", s.getUserOrders) // Alias legacy
	r.GET("/api/notifications", s.getNotifications)
	r.GET("/api/get_order_notifications.php", s.getNotifications) // Alias legacy
	r.POST("/api/notifications/admin", s.notifyAdmin)
	r.POST("/api/notify_admin_payment.php", s.notifyAdmin) // Alias legacy
	r.GET("/api/trucks", s.getTrucks)
	r.GET("/api/get_nearby_trucks.php", s.getTrucks) // Alias legacy
	r.GET("/api/get_truck_status.php", s.getTrucks) // Alias legacy
	r.GET("/api/get_truck_schedules.php", s.getTrucks) // Alias legacy
	r.PUT("/api/trucks/:id", s.updateTruck)
	r.POST("/api/update_truck_status.php", s.updateTruck) // Alias legacy
	r.POST("/api/update_truck_config.php", s.updateTruck) // Alias legacy
	r.POST("/api/update_truck_schedule.php", s.updateTruck) // Alias legacy
	r.POST("/api/location", s.handleLocation)
	r.POST("/api/location/geocode.php", s.handleLocation) // Alias legacy
	r.POST("/api/location/save_location.php", s.handleLocation) // Alias legacy
	r.POST("/api/location/check_delivery_zone.php", s.handleLocation) // Alias legacy
	r.POST("/api/location/get_nearby_products.php", s.handleLocation) // Alias legacy
	r.POST("/api/location/calculate_delivery_time.php", s.handleLocation) // Alias legacy
	r.GET("/api/auth/session", s.checkSession)
	r.GET("/api/auth/check_session.php", s.checkSession) // Alias legacy
	r.PUT("/api/users/:id", s.updateUser)
	r.POST("/api/update_cashier_profile.php", s.updateUser) // Alias legacy
	r.DELETE("/api/users/:id", s.deleteUser)
	r.POST("/api/auth/delete_account.php", s.deleteUser) // Alias legacy
	r.POST("/api/track", s.trackUsage)
	r.POST("/api/track_usage.php", s.trackUsage) // Alias legacy

	// Comandas & Tracking
	r.GET("/api/comandas", s.getComandas)
	r.GET("/api/get_comandas_v2.php", s.getComandas) // Alias legacy
	r.GET("/api/comandas/realtime", s.realtimeComandas) // SSE realtime
	r.PUT("/api/comandas/:id/status", s.updateComandaStatus)
	r.POST("/api/track/visit", s.trackVisit)
	r.POST("/api/track_visit.php", s.trackVisit) // Alias legacy
	r.POST("/api/track/interaction", s.trackInteraction)
	r.POST("/api/track_interaction.php", s.trackInteraction) // Alias legacy

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Printf("🚀 API on :%s", port)
	r.Run(":" + port)
}
