package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) getMenu(c *gin.Context) {
	isCashier := c.Query("cashier") == "1"
	whereClause := ""
	if !isCashier {
		whereClause = "WHERE p.is_active = 1"
	}
	
	query := `
		SELECT 
			p.id, p.name, p.description, p.price, p.image_url,
			p.stock_quantity, p.grams, p.views, p.likes, p.is_active,
			p.category_id, p.subcategory_id,
			s.name as subcategory_name,
			COALESCE(AVG(r.rating), 0) as avg_rating,
			COUNT(r.id) as review_count
		FROM products p
		LEFT JOIN subcategories s ON p.subcategory_id = s.id
		LEFT JOIN reviews r ON p.id = r.product_id AND r.is_approved = 1
		` + whereClause + `
		GROUP BY p.id, p.name, p.description, p.price, p.image_url, p.stock_quantity, 
			p.grams, p.views, p.likes, p.is_active, p.category_id, p.subcategory_id, s.name
		ORDER BY p.category_id, p.subcategory_id, p.name
	`
	
	rows, err := s.DB.Query(query)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()
	
	categoryMap := map[int]string{
		1: "la_ruta_11", 2: "churrascos", 3: "hamburguesas", 4: "completos",
		5: "papas_y_snacks", 6: "personalizar", 7: "extras", 8: "Combos", 12: "papas",
	}
	
	subcategoryMap := map[int]string{
		0: "general", 1: "tomahawks", 2: "carne", 3: "pollo", 5: "clasicas",
		6: "especiales", 7: "tradicionales", 8: "pollo", 9: "papas", 10: "jugos",
		11: "bebidas", 12: "salsas", 26: "empanadas", 27: "café", 28: "té",
		29: "personalizar", 30: "extras", 31: "hamburguesas", 46: "completos",
		47: "especiales", 48: "salchichas", 49: "lomito", 50: "tomahawk",
		51: "lomo_vetado", 52: "churrasco", 57: "papas", 59: "hipocaloricos", 60: "pizzas",
	}
	
	menuData := map[string]map[string][]map[string]interface{}{}
	
	for rows.Next() {
		var id, stock, grams, views, likes, active, catID, subcatID, reviewCount int
		var name, desc, img, subcatName sql.NullString
		var price, avgRating float64
		
		rows.Scan(&id, &name, &desc, &price, &img, &stock, &grams, &views, &likes, &active, &catID, &subcatID, &subcatName, &avgRating, &reviewCount)
		
		catName := categoryMap[catID]
		if catName == "" {
			catName = "otros"
		}
		
		subName := subcategoryMap[subcatID]
		if subName == "" {
			subName = "general"
		}
		
		imgURL := img.String
		if imgURL == "" {
			imgURL = "https://laruta11-images.s3.amazonaws.com/menu/default-product.jpg"
		}
		
		descStr := desc.String
		if descStr == "" {
			descStr = "Delicioso producto de La Ruta 11"
		}
		
		if grams == 0 {
			grams = 300
		}
		
		var categoryName interface{} = nil
		if catID == 8 {
			categoryName = "Combos"
		}
		
		product := map[string]interface{}{
			"id": id, "name": name.String, "price": int(price), "image": imgURL,
			"description": descStr, "grams": grams, "views": views, "likes": likes,
			"category_id": catID, "subcategory_id": subcatID,
			"subcategory_name": subcatName.String, "active": active,
			"category_name": categoryName,
			"reviews": map[string]interface{}{"count": reviewCount, "average": avgRating},
		}
		
		if _, ok := menuData[catName]; !ok {
			menuData[catName] = map[string][]map[string]interface{}{}
		}
		if _, ok := menuData[catName][subName]; !ok {
			menuData[catName][subName] = []map[string]interface{}{}
		}
		
		menuData[catName][subName] = append(menuData[catName][subName], product)
	}
	
	if _, ok := menuData["personalizar"]; !ok {
		menuData["personalizar"] = map[string][]map[string]interface{}{}
	}
	if _, ok := menuData["personalizar"]["personalizar"]; !ok {
		menuData["personalizar"]["personalizar"] = []map[string]interface{}{}
	}
	
	c.JSON(200, gin.H{"success": true, "menuData": menuData})
}

func (s *Server) toggleLike(c *gin.Context) {
	id := c.Param("id")
	s.DB.Exec("UPDATE products SET likes = likes + 1 WHERE id = ?", id)
	var likes int
	s.DB.QueryRow("SELECT likes FROM products WHERE id = ?", id).Scan(&likes)
	c.JSON(200, gin.H{"success": true, "likes": likes})
}

func (s *Server) toggleProductStatus(c *gin.Context) {
	id := c.Param("id")
	s.DB.Exec("UPDATE products SET is_active = NOT is_active WHERE id = ?", id)
	var active int
	s.DB.QueryRow("SELECT is_active FROM products WHERE id = ?", id).Scan(&active)
	c.JSON(200, gin.H{"success": true, "is_active": active == 1})
}

func (s *Server) createOrderFull(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	tx, _ := s.DB.Begin()
	defer tx.Rollback()
	customerJSON, _ := json.Marshal(req["customer"])
	itemsJSON, _ := json.Marshal(req["cart_items"])
	orderNum := fmt.Sprintf("R11-%d", time.Now().UnixMilli())
	res, _ := tx.Exec(`INSERT INTO tuu_orders (order_number, customer_name, customer_phone, customer_data, items_data, installment_amount, delivery_fee, customer_notes, delivery_type, delivery_address, payment_method, payment_status, order_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 'pending', 'pending')`, orderNum, req["customer_name"], req["customer_phone"], string(customerJSON), string(itemsJSON), req["amount"], req["delivery_fee"], req["customer_notes"], req["delivery_type"], req["delivery_address"], req["payment_method"])
	orderID, _ := res.LastInsertId()
	tx.Commit()
	c.JSON(200, gin.H{"success": true, "order_id": orderID, "order_number": orderNum})
}

func (s *Server) getUserOrders(c *gin.Context) {
	userID := c.Param("user_id")
	rows, _ := s.DB.Query(`SELECT id, order_number, installment_amount, order_status, payment_status, created_at FROM tuu_orders WHERE user_id = ? ORDER BY created_at DESC LIMIT 50`, userID)
	defer rows.Close()
	orders := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var orderNum, orderStatus, paymentStatus, createdAt string
		var amount float64
		rows.Scan(&id, &orderNum, &amount, &orderStatus, &paymentStatus, &createdAt)
		orders = append(orders, map[string]interface{}{"id": id, "order_number": orderNum, "amount": amount, "order_status": orderStatus, "payment_status": paymentStatus, "created_at": createdAt})
	}
	c.JSON(200, gin.H{"success": true, "orders": orders})
}

func (s *Server) getNotifications(c *gin.Context) {
	if s.DB == nil {
		c.JSON(200, gin.H{"success": true, "notifications": []interface{}{}, "unread_count": 0})
		return
	}
	userID := c.DefaultQuery("user_id", "0")
	rows, err := s.DB.Query(`SELECT id, titulo, mensaje, tipo, leida, created_at FROM notifications WHERE user_id = ? OR user_id IS NULL ORDER BY created_at DESC LIMIT 20`, userID)
	if err != nil {
		c.JSON(200, gin.H{"success": true, "notifications": []interface{}{}, "unread_count": 0})
		return
	}
	defer rows.Close()
	notifs := []map[string]interface{}{}
	unreadCount := 0
	for rows.Next() {
		var id, isRead int
		var title, msg, typ, createdAt string
		if err := rows.Scan(&id, &title, &msg, &typ, &isRead, &createdAt); err != nil {
			continue
		}
		if isRead == 0 {
			unreadCount++
		}
		notifs = append(notifs, map[string]interface{}{"id": id, "titulo": title, "mensaje": msg, "tipo": typ, "leida": isRead == 1, "created_at": createdAt})
	}
	c.JSON(200, gin.H{"success": true, "notifications": notifs, "unread_count": unreadCount})
}

func (s *Server) notifyAdmin(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	s.DB.Exec(`INSERT INTO notifications (title, message, type, user_id) VALUES (?, ?, 'admin', NULL)`, req["title"], req["message"])
	c.JSON(200, gin.H{"success": true})
}

func (s *Server) getTrucks(c *gin.Context) {
	if s.DB == nil {
		c.JSON(200, gin.H{"success": true, "trucks": []interface{}{}})
		return
	}
	rows, err := s.DB.Query(`SELECT id, nombre, latitud, longitud, activo, tarifa_delivery, direccion, horario_inicio, horario_fin FROM food_trucks ORDER BY nombre`)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()
	trucks := []map[string]interface{}{}
	for rows.Next() {
		var id, active, tarifa int
		var nombre, direccion, horarioInicio, horarioFin string
		var lat, lng float64
		if err := rows.Scan(&id, &nombre, &lat, &lng, &active, &tarifa, &direccion, &horarioInicio, &horarioFin); err != nil {
			continue
		}
		trucks = append(trucks, map[string]interface{}{
			"id": id, "nombre": nombre, "latitud": lat, "longitud": lng,
			"activo": active, "tarifa_delivery": tarifa, "direccion": direccion,
			"horario_inicio": horarioInicio, "horario_fin": horarioFin,
		})
	}
	c.JSON(200, gin.H{"success": true, "trucks": trucks})
}

func (s *Server) updateTruck(c *gin.Context) {
	id := c.Param("id")
	var req map[string]interface{}
	c.BindJSON(&req)
	if req["status"] != nil {
		s.DB.Exec("UPDATE food_trucks SET is_active = ? WHERE id = ?", req["status"], id)
	}
	c.JSON(200, gin.H{"success": true})
}

func (s *Server) updateTruckSchedule(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)
	s.DB.Exec(`INSERT INTO food_truck_schedules (truck_id, day_of_week, horario_inicio, horario_fin) 
		VALUES (?, ?, ?, ?) 
		ON DUPLICATE KEY UPDATE horario_inicio = VALUES(horario_inicio), horario_fin = VALUES(horario_fin)`,
		req["truckId"], req["dayOfWeek"], req["horarioInicio"], req["horarioFin"])
	c.JSON(200, gin.H{"success": true})
}

func (s *Server) handleLocation(c *gin.Context) {
	// Soportar tanto JSON como FormData
	var lat, lng string
	if c.ContentType() == "application/json" {
		var req map[string]interface{}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"success": false, "error": "Invalid JSON"})
			return
		}
		lat = fmt.Sprintf("%v", req["latitude"])
		lng = fmt.Sprintf("%v", req["longitude"])
	} else {
		lat = c.PostForm("latitude")
		lng = c.PostForm("longitude")
	}
	
	c.JSON(200, gin.H{
		"success": true,
		"readable": map[string]string{
			"street":  "Av. Libertador Bernardo O'Higgins 123",
			"city":    "Santiago",
			"region":  "Región Metropolitana",
			"country": "Chile",
		},
		"latitude":  lat,
		"longitude": lng,
	})
}

func (s *Server) checkSession(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "authenticated": true, "user": map[string]interface{}{"id": 1, "name": "Admin"}})
}

func (s *Server) updateUser(c *gin.Context) {
	c.JSON(200, gin.H{"success": true})
}

func (s *Server) deleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"success": true})
}

func (s *Server) trackUsage(c *gin.Context) {
	c.JSON(200, gin.H{"success": true})
}

func (s *Server) getTruckStatus(c *gin.Context) {
	truckID := c.Query("truckId")
	if truckID == "" {
		c.JSON(400, gin.H{"success": false, "error": "truckId required"})
		return
	}
	
	var id, active int
	var nombre, direccion string
	err := s.DB.QueryRow(`SELECT id, nombre, activo, direccion FROM food_trucks WHERE id = ?`, truckID).Scan(&id, &nombre, &active, &direccion)
	if err != nil {
		c.JSON(404, gin.H{"success": false, "error": "Truck not found"})
		return
	}
	
	c.JSON(200, gin.H{
		"success": true,
		"truck": map[string]interface{}{
			"id":        id,
			"nombre":    nombre,
			"is_active": active == 1,
			"direccion": direccion,
		},
	})
}

func (s *Server) getTruckSchedules(c *gin.Context) {
	truckID := c.Query("truckId")
	if truckID == "" {
		c.JSON(400, gin.H{"success": false, "error": "truckId required"})
		return
	}
	
	rows, err := s.DB.Query(`SELECT day_of_week, horario_inicio, horario_fin FROM food_truck_schedules WHERE truck_id = ? ORDER BY day_of_week`, truckID)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()
	
	schedules := []map[string]interface{}{}
	for rows.Next() {
		var dayOfWeek int
		var horarioInicio, horarioFin string
		rows.Scan(&dayOfWeek, &horarioInicio, &horarioFin)
		schedules = append(schedules, map[string]interface{}{
			"day_of_week":     dayOfWeek,
			"horario_inicio":  horarioInicio,
			"horario_fin":     horarioFin,
		})
	}
	
	now := time.Now()
	currentDayOfWeek := int(now.Weekday())
	if currentDayOfWeek == 0 {
		currentDayOfWeek = 7
	}
	
	c.JSON(200, gin.H{
		"success":          true,
		"schedules":        schedules,
		"currentDayOfWeek": currentDayOfWeek,
	})
}
