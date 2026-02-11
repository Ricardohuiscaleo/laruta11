package main

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) getSalesSummary(c *gin.Context) {
	if s.DB == nil {
		c.JSON(500, gin.H{"success": false, "error": "DB no disponible"})
		return
	}

	daysAgo := 0
	if d := c.Query("days_ago"); d != "" {
		if _, err := time.Parse("2006-01-02", d); err == nil {
			// parse days_ago
		}
	}

	now := time.Now().In(time.FixedZone("CLT", -3*3600))
	hour := now.Hour()
	
	shiftStart := now.Format("2006-01-02")
	if hour >= 0 && hour < 4 {
		shiftStart = now.AddDate(0, 0, -1).Format("2006-01-02")
	}
	if daysAgo > 0 {
		shiftStart = now.AddDate(0, 0, -daysAgo).Format("2006-01-02")
	}

	startUTC := shiftStart + " 20:30:00"
	endUTC := now.AddDate(0, 0, 1).Format("2006-01-02") + " 07:00:00"

	rows, err := s.DB.Query(`
		SELECT payment_method, COUNT(*) as count, SUM(installment_amount) as total
		FROM tuu_orders
		WHERE created_at >= ? AND created_at < ? AND payment_status = 'paid'
		GROUP BY payment_method
	`, startUTC, endUTC)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()

	summary := map[string]map[string]interface{}{
		"cash":       {"count": 0, "total": 0.0},
		"card":       {"count": 0, "total": 0.0},
		"transfer":   {"count": 0, "total": 0.0},
		"pedidosya":  {"count": 0, "total": 0.0},
		"webpay":     {"count": 0, "total": 0.0},
		"rl6_credit": {"count": 0, "total": 0.0},
	}

	for rows.Next() {
		var method string
		var count int
		var total float64
		rows.Scan(&method, &count, &total)
		if _, ok := summary[method]; ok {
			summary[method] = map[string]interface{}{"count": count, "total": total}
		}
	}

	var deliveryCount int
	var deliveryTotal, deliveryExtras float64
	s.DB.QueryRow(`
		SELECT COUNT(*), COALESCE(SUM(delivery_fee), 0), COALESCE(SUM(delivery_extras), 0)
		FROM tuu_orders
		WHERE created_at >= ? AND created_at < ? AND payment_status = 'paid' AND delivery_type = 'delivery' AND delivery_fee > 0
	`, startUTC, endUTC).Scan(&deliveryCount, &deliveryTotal, &deliveryExtras)

	totalGeneral := 0.0
	totalOrders := 0
	for _, v := range summary {
		totalGeneral += v["total"].(float64)
		totalOrders += v["count"].(int)
	}

	c.JSON(200, gin.H{
		"success":        true,
		"summary":        summary,
		"total_general":  totalGeneral,
		"total_orders":   totalOrders,
		"delivery_fees":  deliveryTotal,
		"delivery_count": deliveryCount,
		"delivery_extras": deliveryExtras,
		"shift_hours":    "17:30-04:00",
		"shift_date":     shiftStart,
		"period":         gin.H{"start": startUTC, "end": endUTC},
	})
}

func (s *Server) getSaldoCaja(c *gin.Context) {
	if s.DB == nil {
		c.JSON(500, gin.H{"success": false, "error": "DB no disponible"})
		return
	}

	var saldo float64
	err := s.DB.QueryRow("SELECT saldo_nuevo FROM caja_movimientos ORDER BY id DESC LIMIT 1").Scan(&saldo)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}

	var totalCash float64
	s.DB.QueryRow(`
		SELECT COALESCE(SUM(installment_amount), 0)
		FROM tuu_orders
		WHERE payment_method = 'cash' AND payment_status = 'paid' AND DATE(created_at) = CURDATE()
	`).Scan(&totalCash)

	var ingresosAuto float64
	s.DB.QueryRow(`
		SELECT COALESCE(SUM(monto), 0)
		FROM caja_movimientos
		WHERE tipo = 'ingreso' AND usuario = 'Sistema' AND DATE(fecha_movimiento) = CURDATE()
	`).Scan(&ingresosAuto)

	c.JSON(200, gin.H{
		"success":                   true,
		"saldo_actual":              saldo,
		"efectivo_dia":              totalCash,
		"ingresos_automaticos_dia":  ingresosAuto,
	})
}
