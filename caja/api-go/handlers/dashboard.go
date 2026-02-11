package handlers

import (
	"database/sql"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// GetDashboard - Endpoint consolidado: 8 requests → 1 request
// GET /api/dashboard?include=analytics,cards,sales,comparison,projection
func GetDashboard(c *gin.Context, db *sql.DB) {
	includes := c.Query("include")
	if includes == "" {
		includes = "analytics,cards,sales" // Default
	}

	var wg sync.WaitGroup
	results := make(map[string]interface{})
	mu := &sync.Mutex{}

	// Analytics
	if strings.Contains(includes, "analytics") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := fetchAnalytics(db)
			mu.Lock()
			results["analytics"] = data
			mu.Unlock()
		}()
	}

	// Cards (compras, inventario, plan)
	if strings.Contains(includes, "cards") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := fetchCards(db)
			mu.Lock()
			results["cards"] = data
			mu.Unlock()
		}()
	}

	// Sales analytics
	if strings.Contains(includes, "sales") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := fetchSalesAnalytics(db)
			mu.Lock()
			results["sales"] = data
			mu.Unlock()
		}()
	}

	// Month comparison
	if strings.Contains(includes, "comparison") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := fetchMonthComparison(db)
			mu.Lock()
			results["comparison"] = data
			mu.Unlock()
		}()
	}

	wg.Wait()

	c.JSON(200, gin.H{
		"success": true,
		"data":    results,
	})
}

func fetchAnalytics(db *sql.DB) map[string]interface{} {
	var totalSales, avgTicket float64
	var totalOrders int
	
	query := `SELECT COALESCE(SUM(total_amount),0), COUNT(*), COALESCE(AVG(total_amount),0) FROM tuu_orders WHERE MONTH(created_at)=MONTH(NOW()) AND YEAR(created_at)=YEAR(NOW())`
	db.QueryRow(query).Scan(&totalSales, &totalOrders, &avgTicket)
	
	return map[string]interface{}{"total_sales": totalSales, "total_orders": totalOrders, "avg_ticket": avgTicket}
}

func fetchCards(db *sql.DB) map[string]interface{} {
	var comprasTotal float64
	var numCompras, itemsActivos int
	var valorInventario float64
	
	db.QueryRow(`SELECT COALESCE(SUM(total),0), COUNT(*) FROM compras WHERE MONTH(fecha)=MONTH(NOW())`).Scan(&comprasTotal, &numCompras)
	db.QueryRow(`SELECT COUNT(*), COALESCE(SUM(stock_quantity*cost_price),0) FROM ingredientes WHERE is_active=1`).Scan(&itemsActivos, &valorInventario)
	
	return map[string]interface{}{"compras": map[string]interface{}{"total_mes": comprasTotal, "numero_compras": numCompras}, "inventario": map[string]interface{}{"valor_total": valorInventario, "items_activos": itemsActivos}}
}

func fetchSalesAnalytics(db *sql.DB) map[string]interface{} {
	var revenue, cost float64
	db.QueryRow(`SELECT COALESCE(SUM(total_amount),0), COALESCE(SUM(cost_amount),0) FROM tuu_orders WHERE MONTH(created_at)=MONTH(NOW())`).Scan(&revenue, &cost)
	return map[string]interface{}{"period": "month", "total_revenue": revenue, "total_profit": revenue - cost}
}

func fetchMonthComparison(db *sql.DB) map[string]interface{} {
	var current, previous float64
	db.QueryRow(`SELECT COALESCE(SUM(total_amount),0) FROM tuu_orders WHERE MONTH(created_at)=MONTH(NOW())`).Scan(&current)
	db.QueryRow(`SELECT COALESCE(SUM(total_amount),0) FROM tuu_orders WHERE MONTH(created_at)=MONTH(NOW())-1`).Scan(&previous)
	growth := 0.0
	if previous > 0 { growth = ((current - previous) / previous) * 100 }
	return map[string]interface{}{"current_month": current, "previous_month": previous, "growth": growth}
}
