package main

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== COMANDAS (POS Orders) ==========

// GET /api/comandas?status=pending&date=2026-02-10
func (s *Server) getComandas(c *gin.Context) {
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
	
	query += " ORDER BY o.created_at DESC"
	
	rows, err := s.DB.Query(query, params...)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()
	
	comandas := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var orderNum, customerName, customerPhone, itemsData, orderStatus, paymentStatus, paymentMethod, deliveryType, createdAt string
		var amount float64
		
		if err := rows.Scan(&id, &orderNum, &customerName, &customerPhone, &itemsData, &amount, &orderStatus, &paymentStatus, &paymentMethod, &deliveryType, &createdAt); err != nil {
			continue
		}
		
		var items []interface{}
		if itemsData != "" {
			json.Unmarshal([]byte(itemsData), &items)
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
	
	s.DB.Exec(`
		INSERT INTO site_visits (ip_address, user_agent, page_url, visit_date, metadata)
		VALUES (?, ?, ?, CURDATE(), ?)
		ON DUPLICATE KEY UPDATE visit_count = visit_count + 1`,
		ip, userAgent, req["page"], req["metadata"])
	
	c.JSON(200, gin.H{"success": true})
}

// POST /api/track/interaction
func (s *Server) trackInteraction(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	
	s.DB.Exec(`
		INSERT INTO user_interactions (user_id, action_type, product_id, metadata, timestamp)
		VALUES (?, ?, ?, ?, NOW())`,
		req["user_id"], req["action_type"], req["product_id"], req["metadata"])
	
	// Actualizar analytics de producto si aplica
	if req["product_id"] != nil && req["action_type"] == "view" {
		s.DB.Exec(`
			INSERT INTO product_analytics (product_id, product_name, views_count)
			VALUES (?, ?, 1)
			ON DUPLICATE KEY UPDATE views_count = views_count + 1`,
			req["product_id"], req["product_name"])
	}
	
	if req["product_id"] != nil && req["action_type"] == "click" {
		s.DB.Exec(`
			UPDATE product_analytics SET clicks_count = clicks_count + 1
			WHERE product_id = ?`, req["product_id"])
	}
	
	c.JSON(200, gin.H{"success": true})
}
