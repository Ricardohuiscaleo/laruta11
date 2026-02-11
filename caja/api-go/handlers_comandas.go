package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== COMANDAS (POS Orders) ==========

// GET /api/comandas?status=pending&date=2026-02-10
func (s *Server) getComandas(c *gin.Context) {
	if s.DB == nil {
		c.JSON(200, gin.H{"success": true, "comandas": []interface{}{}, "total": 0})
		return
	}
	status := c.DefaultQuery("status", "all")
	date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	
	query := `
		SELECT o.id, o.order_number, o.customer_name, o.customer_phone,
			o.items_data, o.installment_amount, o.order_status, o.payment_status,
			o.payment_method, o.delivery_type, o.created_at
		FROM tuu_orders o
		WHERE DATE(o.created_at) = ?`
	
	params := []interface{}{date}
	
	if status != "all" {
		query += " AND o.order_status = ?"
		params = append(params, status)
	}
	
	query += " ORDER BY o.created_at DESC LIMIT 100"
	
	rows, err := s.DB.Query(query, params...)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": "DB query failed: " + err.Error()})
		return
	}
	defer rows.Close()
	
	comandas := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var orderNum, customerName, customerPhone, itemsData, orderStatus, paymentStatus, paymentMethod, deliveryType, createdAt string
		var amount float64
		
		if err := rows.Scan(&id, &orderNum, &customerName, &customerPhone, &itemsData, &amount, &orderStatus, &paymentStatus, &paymentMethod, &deliveryType, &createdAt); err != nil {
			c.JSON(500, gin.H{"success": false, "error": "Row scan failed: " + err.Error()})
			return
		}
		
		var items []interface{}
		if itemsData != "" && itemsData != "null" {
			if err := json.Unmarshal([]byte(itemsData), &items); err != nil {
				items = []interface{}{}
			}
		}
		
		comandas = append(comandas, map[string]interface{}{
			"id":             id,
			"order_number":   orderNum,
			"customer_name":  customerName,
			"customer_phone": customerPhone,
			"items":          items,
			"total":          amount,
			"order_status":   orderStatus,
			"payment_status": paymentStatus,
			"payment_method": paymentMethod,
			"delivery_type":  deliveryType,
			"created_at":     createdAt,
		})
	}
	
	c.JSON(200, gin.H{"success": true, "comandas": comandas, "total": len(comandas)})
}

// PUT /api/comandas/:id/status
func (s *Server) updateComandaStatus(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	c.BindJSON(&req)
	
	if req["order_status"] != nil {
		s.DB.Exec("UPDATE tuu_orders SET order_status = ?, updated_at = NOW() WHERE id = ?", req["order_status"], id)
	}
	
	c.JSON(200, gin.H{"success": true})
}

// ========== TRACKING & ANALYTICS ==========

// POST /api/track/visit
func (s *Server) trackVisit(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	pageURL := req["page_url"]
	if pageURL == nil {
		pageURL = "https://app.laruta11.cl"
	}
	referrer := c.GetHeader("Referer")
	sessionID := req["session_id"]
	
	// Detect device
	deviceType := "desktop"
	if strings.Contains(userAgent, "Mobile") || strings.Contains(userAgent, "Android") || strings.Contains(userAgent, "iPhone") {
		deviceType = "mobile"
	}
	if strings.Contains(userAgent, "iPad") {
		deviceType = "tablet"
	}
	
	// Detect browser
	browser := "Unknown"
	if strings.Contains(userAgent, "Chrome") {
		browser = "Chrome"
	} else if strings.Contains(userAgent, "Firefox") {
		browser = "Firefox"
	} else if strings.Contains(userAgent, "Safari") {
		browser = "Safari"
	} else if strings.Contains(userAgent, "Edge") {
		browser = "Edge"
	}
	
	s.DB.Exec(`
		INSERT INTO site_visits 
		(ip_address, user_agent, page_url, referrer, session_id, visit_date, device_type, browser, 
		 latitude, longitude, screen_resolution, viewport_size, timezone, language, platform)
		VALUES (?, ?, ?, ?, ?, CURDATE(), ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		ip, userAgent, pageURL, referrer, sessionID, deviceType, browser,
		req["latitude"], req["longitude"], req["screen_resolution"], 
		req["viewport_size"], req["timezone"], req["language"], req["platform"])
	
	c.JSON(200, gin.H{"success": true, "message": "Visit tracked"})
}

// POST /api/track/interaction
func (s *Server) trackInteraction(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	
	ip := c.ClientIP()
	actionType := req["action_type"]
	if actionType == nil {
		actionType = "click"
	}
	
	s.DB.Exec(`
		INSERT INTO user_interactions 
		(session_id, user_ip, action_type, element_type, element_id, element_text, product_id, category_id, page_url)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req["session_id"], ip, actionType, req["element_type"], req["element_id"],
		req["element_text"], req["product_id"], req["category_id"], req["page_url"])
	
	// Update product analytics
	if req["product_id"] != nil {
		field := ""
		switch actionType {
		case "view":
			field = "views_count"
		case "click":
			field = "clicks_count"
		case "add_to_cart":
			field = "cart_adds"
		case "remove_from_cart":
			field = "cart_removes"
		}
		
		if field != "" {
			s.DB.Exec(`
				INSERT INTO product_analytics (product_id, product_name, `+field+`)
				VALUES (?, ?, 1)
				ON DUPLICATE KEY UPDATE `+field+` = `+field+` + 1`,
				req["product_id"], req["product_name"])
		}
	}
	
	c.JSON(200, gin.H{"success": true})
}
