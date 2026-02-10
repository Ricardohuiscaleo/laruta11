package main

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID          int             `json:"id"`
	OrderNumber string          `json:"order_number"`
	Customer    json.RawMessage `json:"customer"`
	Items       json.RawMessage `json:"items"`
	Total       int             `json:"total"`
	Status      string          `json:"status"`
	CreatedAt   string          `json:"created_at"`
}

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID int    `json:"category_id"`
	IsActive   bool   `json:"is_active"`
	Type       string `json:"type"`
}

func (s *Server) getPendingOrders(c *gin.Context) {
	rows, err := s.db.Query(`
		SELECT 
			id, order_number, customer_data, items_data, 
			total_amount, status, created_at
		FROM orders 
		WHERE status IN ('pending', 'paid') 
		ORDER BY created_at DESC
	`)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var o Order
		var createdAt time.Time
		var customerData, itemsData string

		err := rows.Scan(&o.ID, &o.OrderNumber, &customerData, &itemsData, &o.Total, &o.Status, &createdAt)
		if err != nil {
			continue
		}

		o.Customer = json.RawMessage(customerData)
		o.Items = json.RawMessage(itemsData)
		o.CreatedAt = createdAt.Format("02/01/2006 15:04")
		orders = append(orders, o)
	}

	c.JSON(200, gin.H{"success": true, "orders": orders})
}

func (s *Server) updateOrderStatus(c *gin.Context) {
	var req struct {
		OrderID       int    `json:"order_id"`
		OrderStatus   string `json:"order_status"`
		PaymentStatus string `json:"payment_status"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "Invalid request"})
		return
	}

	if req.OrderID == 0 || (req.OrderStatus == "" && req.PaymentStatus == "") {
		c.JSON(400, gin.H{"success": false, "error": "Missing required fields"})
		return
	}

	if req.OrderStatus != "" {
		_, err := s.db.Exec("UPDATE tuu_orders SET order_status = ? WHERE id = ?", req.OrderStatus, req.OrderID)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "error": "Failed to update order status"})
			return
		}
	}

	if req.PaymentStatus != "" {
		_, err := s.db.Exec("UPDATE tuu_orders SET payment_status = ? WHERE id = ?", req.PaymentStatus, req.OrderID)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "error": "Failed to update payment status"})
			return
		}
	}

	c.JSON(200, gin.H{"success": true})
}

func (s *Server) getProducts(c *gin.Context) {
	includeInactive := c.Query("include_inactive") == "1"

	var query string
	if includeInactive {
		query = "SELECT id, name, price, category_id, is_active, 'product' as type FROM products ORDER BY is_active DESC, name ASC"
	} else {
		query = "SELECT id, name, price, category_id, is_active, 'product' as type FROM products WHERE is_active = 1 ORDER BY name ASC"
	}

	rows, err := s.db.Query(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID, &p.IsActive, &p.Type); err != nil {
			continue
		}
		products = append(products, p)
	}

	// Get combos
	var comboQuery string
	if includeInactive {
		comboQuery = "SELECT id, name, price, 8 as category_id, active as is_active, 'combo' as type FROM combos ORDER BY active DESC, name ASC"
	} else {
		comboQuery = "SELECT id, name, price, 8 as category_id, active as is_active, 'combo' as type FROM combos WHERE active = 1 ORDER BY name ASC"
	}

	rows, err = s.db.Query(comboQuery)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID, &p.IsActive, &p.Type); err != nil {
				continue
			}
			products = append(products, p)
		}
	}

	c.JSON(200, products)
}

func (s *Server) getProductByID(c *gin.Context) {
	id := c.Param("id")

	var p Product
	err := s.db.QueryRow("SELECT id, name, price, category_id, is_active, 'product' as type FROM products WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID, &p.IsActive, &p.Type)

	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, p)
}
