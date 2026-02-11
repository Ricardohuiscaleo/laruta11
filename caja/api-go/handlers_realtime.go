package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /api/comandas/realtime (SSE)
func (s *Server) realtimeComandas(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")
	
	flusher, ok := c.Writer.(Flusher)
	if !ok {
		c.JSON(500, gin.H{"error": "Streaming not supported"})
		return
	}
	
	clientGone := c.Request.Context().Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	
	// Enviar comandas iniciales
	sendComandas(s, c, flusher)
	
	for {
		select {
		case <-clientGone:
			return
		case <-ticker.C:
			sendComandas(s, c, flusher)
		}
	}
}

type Flusher interface {
	Flush()
}

func sendComandas(s *Server, c *gin.Context, flusher Flusher) {
	date := time.Now().Format("2006-01-02")
	
	rows, err := s.DB.Query(`
		SELECT o.id, o.order_number, o.customer_name, o.customer_phone,
			o.items_data, o.installment_amount, o.order_status, o.payment_status,
			o.payment_method, o.delivery_type, o.created_at
		FROM tuu_orders o
		WHERE DATE(o.created_at) = ? AND o.order_status IN ('pending', 'preparing')
		ORDER BY o.created_at DESC`, date)
	
	if err != nil {
		return
	}
	defer rows.Close()
	
	comandas := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var orderNum, customerName, customerPhone, itemsData, orderStatus, paymentStatus, paymentMethod, deliveryType, createdAt string
		var amount float64
		
		rows.Scan(&id, &orderNum, &customerName, &customerPhone, &itemsData, &amount, &orderStatus, &paymentStatus, &paymentMethod, &deliveryType, &createdAt)
		
		comandas = append(comandas, map[string]interface{}{
			"id":             id,
			"order_number":   orderNum,
			"customer_name":  customerName,
			"customer_phone": customerPhone,
			"items":          json.RawMessage(itemsData),
			"total":          amount,
			"order_status":   orderStatus,
			"payment_status": paymentStatus,
			"payment_method": paymentMethod,
			"delivery_type":  deliveryType,
			"created_at":     createdAt,
		})
	}
	
	data, _ := json.Marshal(map[string]interface{}{
		"success":  true,
		"comandas": comandas,
		"total":    len(comandas),
		"timestamp": time.Now().Unix(),
	})
	
	fmt.Fprintf(c.Writer, "data: %s\n\n", data)
	flusher.Flush()
}
