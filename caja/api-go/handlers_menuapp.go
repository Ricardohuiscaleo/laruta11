package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== MENU & PRODUCTS (consolidado: 3 PHP → 1 Go) ==========

// GET /api/menu?active_only=1
func (s *Server) getMenu(c *gin.Context) {
	activeOnly := c.DefaultQuery("active_only", "1") == "1"
	query := `
		SELECT p.id, p.name, COALESCE(p.description,''), p.price, COALESCE(p.image_url,''), 
			COALESCE(p.stock_quantity,0), p.is_active, COALESCE(p.likes,0), 
			COALESCE(c.name,'Sin categoría'), COALESCE(c.id,0)
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id`
	
	if activeOnly {
		query += " WHERE p.is_active = 1"
	}
	query += " ORDER BY c.display_order, p.name"
	
	rows, err := s.DB.Query(query)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()
	
	menu := map[string][]map[string]interface{}{}
	for rows.Next() {
		var id, catID, stock, likes, active int
		var name, desc, img, catName string
		var price float64
		if err := rows.Scan(&id, &name, &desc, &price, &img, &stock, &active, &likes, &catName, &catID); err != nil {
			continue
		}
		
		if _, ok := menu[catName]; !ok {
			menu[catName] = []map[string]interface{}{}
		}
		menu[catName] = append(menu[catName], map[string]interface{}{
			"id": id, "name": name, "description": desc, "price": price,
			"image_url": img, "stock": stock, "active": active, "likes": likes,
			"category": catName, "category_id": catID,
		})
	}
	c.JSON(200, gin.H{"success": true, "menu": menu})
}

// POST /api/products/:id/like
func (s *Server) toggleLike(c *gin.Context) {
	id := c.Param("id")
	s.DB.Exec("UPDATE products SET likes = likes + 1 WHERE id = ?", id)
	
	var likes int
	s.DB.QueryRow("SELECT likes FROM products WHERE id = ?", id).Scan(&likes)
	c.JSON(200, gin.H{"success": true, "likes": likes})
}

// PUT /api/products/:id/status
func (s *Server) toggleProductStatus(c *gin.Context) {
	id := c.Param("id")
	s.DB.Exec("UPDATE products SET is_active = NOT is_active WHERE id = ?", id)
	
	var active int
	s.DB.QueryRow("SELECT is_active FROM products WHERE id = ?", id).Scan(&active)
	c.JSON(200, gin.H{"success": true, "is_active": active == 1})
}

// ========== ORDERS (consolidado: 4 PHP → 2 Go) ==========

// POST /api/orders (reemplaza create_order.php)
func (s *Server) createOrderFull(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	
	tx, _ := s.DB.Begin()
	defer tx.Rollback()
	
	// Serializar datos
	customerJSON, _ := json.Marshal(req["customer"])
	itemsJSON, _ := json.Marshal(req["cart_items"])
	
	orderNum := fmt.Sprintf("R11-%d", time.Now().UnixMilli())
	if req["order_number"] != nil {
		orderNum = req["order_number"].(string)
	}
	
	res, _ := tx.Exec(`
		INSERT INTO tuu_orders (
			order_number, customer_name, customer_phone, customer_data,
			items_data, installment_amount, delivery_fee, customer_notes,
			delivery_type, delivery_address, payment_method, payment_status, order_status
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'pending', 'pending')`,
		orderNum,
		req["customer_name"],
		req["customer_phone"],
		string(customerJSON),
		string(itemsJSON),
		req["amount"],
		req["delivery_fee"],
		req["customer_notes"],
		req["delivery_type"],
		req["delivery_address"],
		req["payment_method"],
	)
	
	orderID, _ := res.LastInsertId()
	
	// Insertar items
	if items, ok := req["cart_items"].([]interface{}); ok {
		for _, item := range items {
			it := item.(map[string]interface{})
			tx.Exec(`
				INSERT INTO tuu_order_items (order_id, order_reference, product_id, product_name, quantity, item_cost, subtotal)
				VALUES (?, ?, ?, ?, ?, ?, ?)`,
				orderID, orderNum, it["id"], it["name"], it["quantity"], it["price"], 
				float64(it["quantity"].(float64)) * it["price"].(float64),
			)
		}
	}
	
	tx.Commit()
	c.JSON(200, gin.H{"success": true, "order_id": orderID, "order_number": orderNum})
}

// GET /api/orders/user/:user_id
func (s *Server) getUserOrders(c *gin.Context) {
	userID := c.Param("user_id")
	
	rows, _ := s.DB.Query(`
		SELECT id, order_number, installment_amount, order_status, payment_status, created_at
		FROM tuu_orders WHERE user_id = ? ORDER BY created_at DESC LIMIT 50`, userID)
	defer rows.Close()
	
	orders := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var orderNum, orderStatus, paymentStatus, createdAt string
		var amount float64
		rows.Scan(&id, &orderNum, &amount, &orderStatus, &paymentStatus, &createdAt)
		orders = append(orders, map[string]interface{}{
			"id": id, "order_number": orderNum, "amount": amount,
			"order_status": orderStatus, "payment_status": paymentStatus, "created_at": createdAt,
		})
	}
	c.JSON(200, gin.H{"success": true, "orders": orders})
}

// GET /api/notifications?user_id=X
func (s *Server) getNotifications(c *gin.Context) {
	userID := c.DefaultQuery("user_id", "0")
	
	rows, _ := s.DB.Query(`
		SELECT id, title, message, type, is_read, created_at
		FROM notifications WHERE user_id = ? OR user_id IS NULL
		ORDER BY created_at DESC LIMIT 20`, userID)
	defer rows.Close()
	
	notifs := []map[string]interface{}{}
	for rows.Next() {
		var id, isRead int
		var title, msg, typ, createdAt string
		rows.Scan(&id, &title, &msg, &typ, &isRead, &createdAt)
		notifs = append(notifs, map[string]interface{}{
			"id": id, "titulo": title, "mensaje": msg, "type": typ,
			"leida": isRead == 1, "created_at": createdAt,
		})
	}
	c.JSON(200, gin.H{"success": true, "notifications": notifs})
}

// POST /api/notifications/admin
func (s *Server) notifyAdmin(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	
	s.DB.Exec(`
		INSERT INTO notifications (title, message, type, user_id)
		VALUES (?, ?, 'admin', NULL)`,
		req["title"], req["message"])
	
	c.JSON(200, gin.H{"success": true})
}

// ========== TRUCKS & LOCATION (consolidado: 11 PHP → 3 Go) ==========

// GET /api/trucks?nearby=1&lat=X&lng=Y
func (s *Server) getTrucks(c *gin.Context) {
	nearby := c.Query("nearby") == "1"
	
	if nearby {
		lat := c.Query("lat")
		lng := c.Query("lng")
		
		// Calcular distancia con Haversine
		rows, _ := s.DB.Query(`
			SELECT id, name, latitude, longitude, is_active, tarifa_delivery,
				(6371 * acos(cos(radians(?)) * cos(radians(latitude)) * 
				cos(radians(longitude) - radians(?)) + sin(radians(?)) * 
				sin(radians(latitude)))) AS distance
			FROM food_trucks
			WHERE is_active = 1
			HAVING distance < 10
			ORDER BY distance LIMIT 5`, lat, lng, lat)
		defer rows.Close()
		
		trucks := []map[string]interface{}{}
		for rows.Next() {
			var id, active, tarifa int
			var name string
			var lat, lng, dist float64
			rows.Scan(&id, &name, &lat, &lng, &active, &tarifa, &dist)
			trucks = append(trucks, map[string]interface{}{
				"id": id, "name": name, "latitude": lat, "longitude": lng,
				"is_active": active == 1, "tarifa_delivery": tarifa, "distance": dist,
			})
		}
		c.JSON(200, gin.H{"success": true, "trucks": trucks})
	} else {
		rows, _ := s.DB.Query(`
			SELECT id, name, latitude, longitude, is_active, tarifa_delivery, schedule
			FROM food_trucks ORDER BY name`)
		defer rows.Close()
		
		trucks := []map[string]interface{}{}
		for rows.Next() {
			var id, active, tarifa int
			var name, schedule string
			var lat, lng float64
			rows.Scan(&id, &name, &lat, &lng, &active, &tarifa, &schedule)
			trucks = append(trucks, map[string]interface{}{
				"id": id, "name": name, "latitude": lat, "longitude": lng,
				"is_active": active == 1, "tarifa_delivery": tarifa, "schedule": schedule,
			})
		}
		c.JSON(200, gin.H{"success": true, "trucks": trucks})
	}
}

// PUT /api/trucks/:id
func (s *Server) updateTruck(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	c.BindJSON(&req)
	
	if req["status"] != nil {
		s.DB.Exec("UPDATE food_trucks SET is_active = ? WHERE id = ?", req["status"], id)
	}
	if req["schedule"] != nil {
		s.DB.Exec("UPDATE food_trucks SET schedule = ? WHERE id = ?", req["schedule"], id)
	}
	if req["config"] != nil {
		configJSON, _ := json.Marshal(req["config"])
		s.DB.Exec("UPDATE food_trucks SET config = ? WHERE id = ?", string(configJSON), id)
	}
	
	c.JSON(200, gin.H{"success": true})
}

// POST /api/location (consolidado: 5 endpoints location)
func (s *Server) handleLocation(c *gin.Context) {
	action := c.DefaultQuery("action", "save")
	var req map[string]interface{}
	c.BindJSON(&req)
	
	switch action {
	case "geocode":
		// Geocodificar coordenadas (mock - usar API real en prod)
		c.JSON(200, gin.H{"success": true, "address": "Dirección aproximada", "lat": req["lat"], "lng": req["lng"]})
		
	case "save":
		// Guardar ubicación del usuario
		s.DB.Exec(`
			INSERT INTO user_locations (user_id, latitude, longitude, address)
			VALUES (?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE latitude = ?, longitude = ?, address = ?`,
			req["user_id"], req["lat"], req["lng"], req["address"],
			req["lat"], req["lng"], req["address"])
		c.JSON(200, gin.H{"success": true})
		
	case "check_delivery":
		// Verificar si está en zona de delivery
		lat := req["lat"].(float64)
		lng := req["lng"].(float64)
		
		var count int
		s.DB.QueryRow(`
			SELECT COUNT(*) FROM food_trucks
			WHERE is_active = 1 AND
			(6371 * acos(cos(radians(?)) * cos(radians(latitude)) * 
			cos(radians(longitude) - radians(?)) + sin(radians(?)) * 
			sin(radians(latitude)))) < 10`, lat, lng, lat).Scan(&count)
		
		c.JSON(200, gin.H{"success": true, "in_zone": count > 0})
		
	case "nearby_products":
		// Productos disponibles cerca
		rows, _ := s.DB.Query(`
			SELECT DISTINCT p.id, p.name, p.price
			FROM products p
			JOIN food_trucks t ON t.is_active = 1
			WHERE p.is_active = 1 AND p.stock_quantity > 0
			LIMIT 20`)
		defer rows.Close()
		
		products := []map[string]interface{}{}
		for rows.Next() {
			var id int
			var name string
			var price float64
			rows.Scan(&id, &name, &price)
			products = append(products, map[string]interface{}{"id": id, "name": name, "price": price})
		}
		c.JSON(200, gin.H{"success": true, "products": products})
		
	case "delivery_time":
		// Calcular tiempo de delivery (mock)
		c.JSON(200, gin.H{"success": true, "estimated_time": 30})
		
	default:
		c.JSON(400, gin.H{"success": false, "error": "Invalid action"})
	}
}

// ========== USERS (consolidado: 3 PHP → 2 Go) ==========

// GET /api/auth/session
func (s *Server) checkSession(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(401, gin.H{"success": false, "authenticated": false})
		return
	}
	
	var id int
	var name, email string
	err := s.DB.QueryRow(`
		SELECT id, name, email FROM app_users WHERE id = ?`, userID).Scan(&id, &name, &email)
	
	if err == sql.ErrNoRows {
		c.JSON(401, gin.H{"success": false, "authenticated": false})
		return
	}
	
	c.JSON(200, gin.H{"success": true, "authenticated": true, "user": gin.H{
		"id": id, "nombre": name, "email": email,
	}})
}

// PUT /api/users/:id
func (s *Server) updateUser(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	c.BindJSON(&req)
	
	if req["name"] != nil {
		s.DB.Exec("UPDATE app_users SET name = ? WHERE id = ?", req["name"], id)
	}
	if req["phone"] != nil {
		s.DB.Exec("UPDATE app_users SET phone = ? WHERE id = ?", req["phone"], id)
	}
	if req["email"] != nil {
		s.DB.Exec("UPDATE app_users SET email = ? WHERE id = ?", req["email"], id)
	}
	
	c.JSON(200, gin.H{"success": true})
}

// DELETE /api/users/:id
func (s *Server) deleteUser(c *gin.Context) {
	id := c.Param("id")
	s.DB.Exec("UPDATE app_users SET is_active = 0, deleted_at = NOW() WHERE id = ?", id)
	c.JSON(200, gin.H{"success": true})
}

// ========== TRACKING ==========

// POST /api/track
func (s *Server) trackUsage(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	
	s.DB.Exec(`
		INSERT INTO user_interactions (user_id, action_type, product_id, metadata)
		VALUES (?, ?, ?, ?)`,
		req["user_id"], req["action"], req["product_id"], req["metadata"])
	
	c.JSON(200, gin.H{"success": true})
}
