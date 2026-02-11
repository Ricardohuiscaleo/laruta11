package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== AUTH ==========

func (s *Server) authLogin(c *gin.Context) {
	var req struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		Type string `json:"type"`
	}
	c.BindJSON(&req)
	valid, role := false, ""
	switch req.Type {
	case "caja":
		if (req.User == os.Getenv("CAJA_USER_CAJERA") && req.Pass == os.Getenv("CAJA_PASSWORD_CAJERA")) || (req.User == os.Getenv("CAJA_USER_ADMIN") && req.Pass == os.Getenv("CAJA_PASSWORD_ADMIN")) {
			valid, role = true, "caja"
		}
	case "inventario":
		if req.User == os.Getenv("INVENTARIO_USER") && req.Pass == os.Getenv("INVENTARIO_PASSWORD") {
			valid, role = true, "inventario"
		}
	case "comandas":
		if req.User == os.Getenv("COMANDAS_USER") && req.Pass == os.Getenv("COMANDAS_PASSWORD") {
			valid, role = true, "comandas"
		}
	case "admin":
		// Validar contra múltiples usuarios admin desde variables de entorno
		adminUsers := map[string]string{
			os.Getenv("ADMIN_USER_ADMIN"):   os.Getenv("ADMIN_PASSWORD_ADMIN"),
			os.Getenv("ADMIN_USER_RICARDO"): os.Getenv("ADMIN_PASSWORD_RICARDO"),
			os.Getenv("ADMIN_USER_MANAGER"): os.Getenv("ADMIN_PASSWORD_MANAGER"),
			os.Getenv("ADMIN_USER_RUTA11"):  os.Getenv("ADMIN_PASSWORD_RUTA11"),
		}
		if expectedPass, exists := adminUsers[req.User]; exists && expectedPass != "" && req.Pass == expectedPass {
			valid, role = true, "admin"
		}
	}
	if valid {
		c.JSON(200, gin.H{"success": true, "role": role, "user": req.User})
	} else {
		c.JSON(401, gin.H{"success": false, "error": "Credenciales inválidas"})
	}
}

func (s *Server) authCheck(c *gin.Context) {
	if user := c.GetHeader("X-User"); user != "" {
		c.JSON(200, gin.H{"success": true, "authenticated": true})
	} else {
		c.JSON(401, gin.H{"success": false, "authenticated": false})
	}
}

func (s *Server) authLogout(c *gin.Context) {
	c.JSON(200, gin.H{"success": true})
}

// ========== COMPRAS ==========

func (s *Server) getCompras(c *gin.Context) {
	rows, _ := s.DB.Query(`SELECT id, fecha_compra, proveedor, monto_total, metodo_pago, notas FROM compras ORDER BY fecha_compra DESC LIMIT 50`)
	defer rows.Close()

	compras := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var fecha, prov, metodo, notas string
		var monto float64
		rows.Scan(&id, &fecha, &prov, &monto, &metodo, &notas)
		compras = append(compras, map[string]interface{}{"id": id, "fecha_compra": fecha, "proveedor": prov, "monto_total": monto, "metodo_pago": metodo, "notas": notas})
	}
	c.JSON(200, gin.H{"success": true, "compras": compras})
}

func (s *Server) getComprasItems(c *gin.Context) {
	rows, _ := s.DB.Query(`
		SELECT id, name, category, unit, current_stock, 'ingredient' as type FROM ingredients WHERE is_active = 1
		UNION ALL
		SELECT p.id, p.name, COALESCE(c.name, 'Sin categoría'), 'unidad', p.stock_quantity, 'product' 
		FROM products p LEFT JOIN categories c ON p.category_id = c.id WHERE p.is_active = 1
		ORDER BY name`)
	defer rows.Close()

	items := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name, cat, unit, typ string
		var stock float64
		rows.Scan(&id, &name, &cat, &unit, &stock, &typ)
		items = append(items, map[string]interface{}{"id": id, "name": name, "category": cat, "unit": unit, "current_stock": stock, "type": typ})
	}
	c.JSON(200, items)
}

func (s *Server) getProveedores(c *gin.Context) {
	rows, _ := s.DB.Query(`SELECT proveedor, COUNT(*) FROM compras GROUP BY proveedor ORDER BY COUNT(*) DESC`)
	defer rows.Close()

	provs := []map[string]interface{}{}
	for rows.Next() {
		var name string
		var count int
		rows.Scan(&name, &count)
		provs = append(provs, map[string]interface{}{"name": name, "count": count})
	}
	c.JSON(200, gin.H{"success": true, "proveedores": provs})
}

func (s *Server) getSaldoDisponible(c *gin.Context) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	start1 := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 17, 30, 0, 0, time.UTC)
	end1 := time.Date(lastMonth.Year(), lastMonth.Month()+1, 1, 4, 0, 0, 0, time.UTC)
	start2 := time.Date(now.Year(), now.Month(), 1, 17, 30, 0, 0, time.UTC)

	var v1, v2, comp float64
	s.DB.QueryRow(`SELECT COALESCE(SUM(installment_amount - COALESCE(delivery_fee, 0)), 0) FROM tuu_orders WHERE payment_status = 'paid' AND created_at >= ? AND created_at < ?`, start1, end1).Scan(&v1)
	s.DB.QueryRow(`SELECT COALESCE(SUM(installment_amount - COALESCE(delivery_fee, 0)), 0) FROM tuu_orders WHERE payment_status = 'paid' AND created_at >= ?`, start2).Scan(&v2)
	s.DB.QueryRow(`SELECT COALESCE(SUM(monto_total), 0) FROM compras WHERE DATE_FORMAT(fecha_compra, '%Y-%m') = ?`, now.Format("2006-01")).Scan(&comp)

	if lastMonth.Format("2006-01") == "2025-10" {
		v1 += 695433
	}

	saldo := v1 + v2 - 1590000 - comp
	c.JSON(200, gin.H{"success": true, "saldo_disponible": saldo, "ventas_mes_anterior": v1, "ventas_mes_actual": v2, "sueldos": 1590000, "compras_mes": comp})
}

func (s *Server) getHistorialSaldo(c *gin.Context) {
	rows, _ := s.DB.Query(`
		SELECT fecha_compra, 'egreso', monto_total, CONCAT('Compra - ', proveedor) FROM compras
		UNION ALL
		SELECT created_at, 'ingreso', installment_amount, order_number FROM tuu_orders WHERE payment_status = 'paid'
		ORDER BY 1 DESC LIMIT 100`)
	defer rows.Close()

	hist := []map[string]interface{}{}
	for rows.Next() {
		var fecha, tipo, desc string
		var monto float64
		rows.Scan(&fecha, &tipo, &monto, &desc)
		hist = append(hist, map[string]interface{}{"fecha": fecha, "tipo": tipo, "monto": monto, "concepto": desc})
	}
	c.JSON(200, gin.H{"success": true, "movimientos": hist})
}

func (s *Server) getPrecioHistorico(c *gin.Context) {
	id := c.Query("ingrediente_id")
	var precio, cant, sub float64
	var unidad, fecha, prov string
	err := s.DB.QueryRow(`
		SELECT cd.precio_unitario, cd.cantidad, cd.unidad, cd.subtotal, c.fecha_compra, c.proveedor
		FROM compras_detalle cd JOIN compras c ON cd.compra_id = c.id
		WHERE cd.ingrediente_id = ? ORDER BY c.fecha_compra DESC LIMIT 1`, id).Scan(&precio, &cant, &unidad, &sub, &fecha, &prov)

	if err == sql.ErrNoRows {
		c.JSON(200, gin.H{"success": false, "error": "Sin historial"})
		return
	}
	c.JSON(200, gin.H{"success": true, "precio_unitario": precio, "unidad": unidad, "ultima_cantidad": cant, "ultimo_subtotal": sub, "fecha_compra": fecha, "proveedor": prov})
}

func (s *Server) registrarCompra(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)

	tx, _ := s.DB.Begin()
	defer tx.Rollback()

	res, _ := tx.Exec(`INSERT INTO compras (fecha_compra, proveedor, monto_total, metodo_pago, notas, usuario) VALUES (?, ?, ?, ?, ?, ?)`,
		req["fecha_compra"], req["proveedor"], req["monto_total"], req["metodo_pago"], req["notas"], req["usuario"])
	id, _ := res.LastInsertId()

	items := req["items"].([]interface{})
	for _, item := range items {
		it := item.(map[string]interface{})
		tx.Exec(`INSERT INTO compras_detalle (compra_id, ingrediente_id, product_id, item_type, cantidad, unidad, precio_unitario, subtotal) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			id, it["ingrediente_id"], it["product_id"], it["item_type"], it["cantidad"], it["unidad"], it["precio_unitario"], it["subtotal"])

		if it["item_type"] == "ingredient" {
			tx.Exec(`UPDATE ingredients SET current_stock = current_stock + ? WHERE id = ?`, it["cantidad"], it["ingrediente_id"])
		}
	}

	tx.Commit()
	c.JSON(200, gin.H{"success": true, "compra_id": id})
}

func (s *Server) deleteCompra(c *gin.Context) {
	id := c.Param("id")
	tx, _ := s.DB.Begin()
	defer tx.Rollback()

	rows, _ := tx.Query(`SELECT ingrediente_id, cantidad FROM compras_detalle WHERE compra_id = ? AND item_type = 'ingredient'`, id)
	for rows.Next() {
		var ingID int
		var cant float64
		rows.Scan(&ingID, &cant)
		tx.Exec(`UPDATE ingredients SET current_stock = current_stock - ? WHERE id = ?`, cant, ingID)
	}
	rows.Close()

	tx.Exec(`DELETE FROM compras_detalle WHERE compra_id = ?`, id)
	tx.Exec(`DELETE FROM compras WHERE id = ?`, id)
	tx.Commit()

	c.JSON(200, gin.H{"success": true})
}

func (s *Server) uploadRespaldo(c *gin.Context) {
	id := c.Param("id")
	file, _ := c.FormFile("image")
	// TODO: S3 upload
	s.DB.Exec(`UPDATE compras SET imagen_respaldo = ? WHERE id = ?`, file.Filename, id)
	c.JSON(200, gin.H{"success": true})
}

// ========== INVENTORY ==========

// GET /api/ingredientes (5 usos)
func (s *Server) getIngredientes(c *gin.Context) {
	rows, _ := s.DB.Query(`SELECT id, name, unit, current_stock, min_stock_level, category, is_active FROM ingredients ORDER BY name`)
	defer rows.Close()

	ings := []map[string]interface{}{}
	for rows.Next() {
		var id, active int
		var name, unit, cat string
		var stock, min float64
		rows.Scan(&id, &name, &unit, &stock, &min, &cat, &active)
		ings = append(ings, map[string]interface{}{"id": id, "name": name, "unit": unit, "current_stock": stock, "min_stock_level": min, "category": cat, "is_active": active == 1})
	}
	c.JSON(200, gin.H{"success": true, "ingredientes": ings})
}

// POST /api/ingredientes
func (s *Server) saveIngrediente(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)

	if req["id"] != nil && req["id"].(float64) > 0 {
		s.DB.Exec(`UPDATE ingredients SET name = ?, unit = ?, current_stock = ?, min_stock_level = ?, category = ?, is_active = ? WHERE id = ?`,
			req["name"], req["unit"], req["current_stock"], req["min_stock_level"], req["category"], req["is_active"], int(req["id"].(float64)))
		c.JSON(200, gin.H{"success": true, "id": int(req["id"].(float64))})
	} else {
		res, _ := s.DB.Exec(`INSERT INTO ingredients (name, unit, current_stock, min_stock_level, category, is_active) VALUES (?, ?, ?, ?, ?, ?)`,
			req["name"], req["unit"], req["current_stock"], req["min_stock_level"], req["category"], req["is_active"])
		id, _ := res.LastInsertId()
		c.JSON(200, gin.H{"success": true, "id": id})
	}
}

// DELETE /api/ingredientes/:id (2 usos)
func (s *Server) deleteIngrediente(c *gin.Context) {
	s.DB.Exec(`DELETE FROM ingredients WHERE id = ?`, c.Param("id"))
	c.JSON(200, gin.H{"success": true})
}

// GET /api/categories (8 usos + 4 usos = 12 usos)
func (s *Server) getCategories(c *gin.Context) {
	rows, err := s.DB.Query(`SELECT id, name, is_active, display_order FROM categories ORDER BY display_order, name`)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()

	cats := []map[string]interface{}{}
	for rows.Next() {
		var id, order, active int
		var name string
		rows.Scan(&id, &name, &active, &order)
		cats = append(cats, map[string]interface{}{"id": id, "name": name, "is_active": active == 1, "display_order": order})
	}
	c.JSON(200, gin.H{"success": true, "categories": cats})
}

// POST /api/categories
func (s *Server) saveCategory(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)

	if req["id"] != nil && req["id"].(float64) > 0 {
		s.DB.Exec(`UPDATE categories SET name = ?, is_active = ?, display_order = ? WHERE id = ?`,
			req["name"], req["is_active"], req["display_order"], int(req["id"].(float64)))
		c.JSON(200, gin.H{"success": true})
	} else {
		res, _ := s.DB.Exec(`INSERT INTO categories (name, is_active, display_order) VALUES (?, ?, ?)`,
			req["name"], req["is_active"], req["display_order"])
		id, _ := res.LastInsertId()
		c.JSON(200, gin.H{"success": true, "id": id})
	}
}

// DELETE /api/categories/:id
func (s *Server) deleteCategory(c *gin.Context) {
	s.DB.Exec(`DELETE FROM categories WHERE id = ?`, c.Param("id"))
	c.JSON(200, gin.H{"success": true})
}

// GET /api/checklist (4 usos)
func (s *Server) getChecklists(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	rows, err := s.DB.Query(`SELECT id, date, type, items, completed FROM checklists WHERE date = ? ORDER BY type`, date)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()

	chs := []map[string]interface{}{}
	for rows.Next() {
		var id, comp int
		var date, typ, items string
		rows.Scan(&id, &date, &typ, &items, &comp)
		chs = append(chs, map[string]interface{}{"id": id, "date": date, "type": typ, "items": json.RawMessage(items), "completed": comp == 1})
	}
	c.JSON(200, gin.H{"success": true, "checklists": chs})
}

// POST /api/checklist
func (s *Server) saveChecklist(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)

	if req["id"] != nil && req["id"].(float64) > 0 {
		s.DB.Exec(`UPDATE checklists SET items = ?, completed = ? WHERE id = ?`, req["items"], req["completed"], int(req["id"].(float64)))
	} else {
		s.DB.Exec(`INSERT INTO checklists (date, type, items, completed) VALUES (?, ?, ?, ?)`,
			req["date"], req["type"], req["items"], req["completed"])
	}
	c.JSON(200, gin.H{"success": true})
}

// DELETE /api/checklist/:id
func (s *Server) deleteChecklist(c *gin.Context) {
	s.DB.Exec(`DELETE FROM checklists WHERE id = ?`, c.Param("id"))
	c.JSON(200, gin.H{"success": true})
}

// POST /api/create_order (3 usos)
func (s *Server) createOrder(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)

	tx, _ := s.DB.Begin()
	defer tx.Rollback()

	res, _ := tx.Exec(`INSERT INTO tuu_orders (order_number, customer_data, items_data, total_amount, payment_method, order_status, payment_status) VALUES (?, ?, ?, ?, ?, 'pending', 'pending')`,
		req["order_number"], req["customer_data"], req["items_data"], req["total_amount"], req["payment_method"])
	orderID, _ := res.LastInsertId()

	// Deducir inventario
	items := req["items"].([]interface{})
	for _, item := range items {
		it := item.(map[string]interface{})
		if it["product_id"] != nil {
			tx.Exec(`UPDATE products SET stock_quantity = stock_quantity - ? WHERE id = ?`, it["quantity"], it["product_id"])
		}
	}

	tx.Commit()
	c.JSON(200, gin.H{"success": true, "order_id": orderID})
}

// GET /api/products
func (s *Server) getProducts(c *gin.Context) {
	inactive := c.Query("include_inactive") == "1"
	query := "SELECT id, name, price, category_id, is_active FROM products"
	if !inactive {
		query += " WHERE is_active = 1"
	}
	query += " ORDER BY name"

	rows, _ := s.DB.Query(query)
	defer rows.Close()

	products := []map[string]interface{}{}
	for rows.Next() {
		var id, catID, active int
		var name string
		var price float64
		rows.Scan(&id, &name, &price, &catID, &active)
		products = append(products, map[string]interface{}{"id": id, "name": name, "price": price, "category_id": catID, "is_active": active == 1})
	}
	c.JSON(200, gin.H{"success": true, "products": products})
}

// GET /api/products/:id
func (s *Server) getProductByID(c *gin.Context) {
	var id, catID, active int
	var name string
	var price float64
	err := s.DB.QueryRow("SELECT id, name, price, category_id, is_active FROM products WHERE id = ?", c.Param("id")).Scan(&id, &name, &price, &catID, &active)

	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"success": false, "error": "Not found"})
		return
	}
	c.JSON(200, gin.H{"success": true, "product": map[string]interface{}{"id": id, "name": name, "price": price, "category_id": catID, "is_active": active == 1}})
}

// GET /api/orders/pending
func (s *Server) getPendingOrders(c *gin.Context) {
	rows, err := s.DB.Query(`SELECT id, order_number, customer_data, items_data, total_amount, order_status FROM tuu_orders WHERE order_status IN ('pending', 'preparing') ORDER BY created_at DESC`)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	defer rows.Close()

	orders := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var orderNum, customer, items, status string
		var total float64
		rows.Scan(&id, &orderNum, &customer, &items, &total, &status)
		orders = append(orders, map[string]interface{}{"id": id, "order_number": orderNum, "customer": json.RawMessage(customer), "items": json.RawMessage(items), "total": total, "status": status})
	}
	c.JSON(200, gin.H{"success": true, "orders": orders})
}

// POST /api/orders/status
func (s *Server) updateOrderStatus(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)

	if req["order_status"] != nil {
		s.DB.Exec("UPDATE tuu_orders SET order_status = ? WHERE id = ?", req["order_status"], req["order_id"])
	}
	if req["payment_status"] != nil {
		s.DB.Exec("UPDATE tuu_orders SET payment_status = ? WHERE id = ?", req["payment_status"], req["order_id"])
	}
	c.JSON(200, gin.H{"success": true})
}

// ========== DASHBOARD (consolidado - reemplaza 3 PHP endpoints) ==========

// GET /api/dashboard?date=YYYY-MM-DD
func (s *Server) getDashboard(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))

	var wg sync.WaitGroup
	var salesData, behaviorData, analyticsData gin.H

	wg.Add(3)
	go func() { defer wg.Done(); salesData = s.queryTUUSales(date) }()
	go func() { defer wg.Done(); behaviorData = s.queryUserBehavior() }()
	go func() { defer wg.Done(); analyticsData = s.queryAnalytics() }()
	wg.Wait()

	c.JSON(200, gin.H{
		"success":   true,
		"sales":     salesData,
		"behavior":  behaviorData,
		"analytics": analyticsData,
	})
}

func (s *Server) queryTUUSales(date string) gin.H {
	type Tx struct {
		ID            int64   `json:"id"`
		OrderRef      string  `json:"order_reference"`
		Amount        float64 `json:"amount"`
		CustomerName  string  `json:"customer_name"`
		ProductName   string  `json:"product_name"`
		CreatedAt     string  `json:"created_at"`
		PaymentMethod string  `json:"payment_method"`
		DeliveryFee   float64 `json:"delivery_fee"`
		OrderCost     float64 `json:"order_cost"`
	}

	rows, err := s.DB.Query(`
		SELECT o.id, COALESCE(o.order_number,''), 
			COALESCE(o.tuu_amount, o.installment_amount, o.product_price, 0),
			COALESCE(o.customer_name,''), COALESCE(o.product_name,''),
			o.created_at, COALESCE(o.payment_method,'cash'),
			COALESCE(o.delivery_fee,0),
			COALESCE((SELECT SUM(item_cost*quantity) FROM tuu_order_items WHERE order_reference=o.order_number),0)
		FROM tuu_orders o
		WHERE DATE(o.created_at) >= ? AND DATE(o.created_at) <= ? AND o.payment_status='paid'
		ORDER BY o.created_at DESC`, date, date)
	if err != nil {
		return gin.H{"all_transactions": []any{}, "combined_stats": gin.H{}}
	}
	defer rows.Close()

	var txs []Tx
	loc, _ := time.LoadLocation("America/Santiago")
	for rows.Next() {
		var t Tx
		var ca time.Time
		if rows.Scan(&t.ID, &t.OrderRef, &t.Amount, &t.CustomerName, &t.ProductName, &ca, &t.PaymentMethod, &t.DeliveryFee, &t.OrderCost) != nil {
			continue
		}
		chile := ca.In(loc)
		if chile.Hour() < 4 {
			chile = chile.AddDate(0, 0, -1)
		}
		t.CreatedAt = chile.Format("2006-01-02 15:04:05")
		txs = append(txs, t)
	}
	if txs == nil {
		txs = []Tx{}
	}

	byMethod := map[string]gin.H{}
	var totalRev, totalDel float64
	for _, t := range txs {
		totalRev += t.Amount
		totalDel += t.DeliveryFee
		if _, ok := byMethod[t.PaymentMethod]; !ok {
			byMethod[t.PaymentMethod] = gin.H{"sales": 0.0, "orders": 0}
		}
		byMethod[t.PaymentMethod]["sales"] = byMethod[t.PaymentMethod]["sales"].(float64) + t.Amount
		byMethod[t.PaymentMethod]["orders"] = byMethod[t.PaymentMethod]["orders"].(int) + 1
	}

	return gin.H{
		"all_transactions": txs,
		"combined_stats": gin.H{
			"total_revenue": totalRev, "total_transactions": len(txs),
			"total_delivery_fee": totalDel, "by_method": byMethod,
		},
	}
}

func (s *Server) queryUserBehavior() gin.H {
	result := gin.H{"top_products": []any{}, "interactions_today": []any{}, "engagement": gin.H{"avg_time": 0.0, "avg_scroll": 0.0}}

	type P struct {
		Name   string `json:"product_name"`
		Views  int    `json:"views_count"`
		Clicks int    `json:"clicks_count"`
	}
	if rows, err := s.DB.Query(`SELECT product_name, views_count, clicks_count FROM product_analytics ORDER BY views_count DESC LIMIT 5`); err == nil {
		defer rows.Close()
		var ps []P
		for rows.Next() {
			var p P
			rows.Scan(&p.Name, &p.Views, &p.Clicks)
			ps = append(ps, p)
		}
		if ps != nil {
			result["top_products"] = ps
		}
	}

	type I struct {
		Type  string `json:"action_type"`
		Count int    `json:"count"`
	}
	if rows, err := s.DB.Query(`SELECT action_type, COUNT(*) FROM user_interactions WHERE DATE(timestamp)=CURDATE() GROUP BY action_type`); err == nil {
		defer rows.Close()
		var is []I
		for rows.Next() {
			var i I
			rows.Scan(&i.Type, &i.Count)
			is = append(is, i)
		}
		if is != nil {
			result["interactions_today"] = is
		}
	}

	var avgT, avgS sql.NullFloat64
	s.DB.QueryRow(`SELECT AVG(time_spent), AVG(scroll_depth) FROM user_journey WHERE DATE(timestamp)=CURDATE() AND time_spent>0`).Scan(&avgT, &avgS)
	result["engagement"] = gin.H{"avg_time": avgT.Float64, "avg_scroll": avgS.Float64}

	return result
}

func (s *Server) queryAnalytics() gin.H {
	today := time.Now().Format("2006-01-02")
	weekAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	monthAgo := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	var vToday, vWeek, vMonth, totalUsers, newUsers, totalProducts int
	s.DB.QueryRow(`SELECT COUNT(DISTINCT ip_address) FROM site_visits WHERE visit_date=?`, today).Scan(&vToday)
	s.DB.QueryRow(`SELECT COUNT(DISTINCT ip_address) FROM site_visits WHERE visit_date>=?`, weekAgo).Scan(&vWeek)
	s.DB.QueryRow(`SELECT COUNT(DISTINCT ip_address) FROM site_visits WHERE visit_date>=?`, monthAgo).Scan(&vMonth)
	s.DB.QueryRow(`SELECT COUNT(*) FROM app_users`).Scan(&totalUsers)
	if s.DB.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM app_users WHERE DATE(created_at)='%s'`, today)).Scan(&newUsers) != nil {
		s.DB.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM app_users WHERE DATE(registration_date)='%s'`, today)).Scan(&newUsers)
	}
	s.DB.QueryRow(`SELECT COUNT(*) FROM products`).Scan(&totalProducts)

	return gin.H{
		"visitors": gin.H{"today": vToday, "week": vWeek, "month": vMonth},
		"users":    gin.H{"total": totalUsers, "new_today": newUsers},
		"products": gin.H{"total": totalProducts},
	}
}

// ========== DASHBOARD ENDPOINTS (5 PHP → 5 Go) ==========

// GET /api/get_dashboard_analytics.php
func (s *Server) getDashboardAnalytics(c *gin.Context) {
	var totalOrders, totalUsers, totalProducts int
	var totalRevenue float64
	s.DB.QueryRow(`SELECT COUNT(*) FROM tuu_orders WHERE payment_status='paid'`).Scan(&totalOrders)
	s.DB.QueryRow(`SELECT COALESCE(SUM(installment_amount),0) FROM tuu_orders WHERE payment_status='paid'`).Scan(&totalRevenue)
	s.DB.QueryRow(`SELECT COUNT(*) FROM app_users`).Scan(&totalUsers)
	s.DB.QueryRow(`SELECT COUNT(*) FROM products WHERE is_active=1`).Scan(&totalProducts)
	
	c.JSON(200, gin.H{"success": true, "data": gin.H{"total_orders": totalOrders, "total_revenue": totalRevenue, "total_users": totalUsers, "total_products": totalProducts}})
}

// GET /api/get_dashboard_cards.php
func (s *Server) getDashboardCards(c *gin.Context) {
	var ordersToday, ordersMonth int
	var revenueToday, revenueMonth float64
	today := time.Now().Format("2006-01-02")
	s.DB.QueryRow(`SELECT COUNT(*), COALESCE(SUM(installment_amount),0) FROM tuu_orders WHERE DATE(created_at)=? AND payment_status='paid'`, today).Scan(&ordersToday, &revenueToday)
	s.DB.QueryRow(`SELECT COUNT(*), COALESCE(SUM(installment_amount),0) FROM tuu_orders WHERE MONTH(created_at)=MONTH(NOW()) AND payment_status='paid'`).Scan(&ordersMonth, &revenueMonth)
	
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"compras": gin.H{
				"total_mes":      0,
				"numero_compras":  0,
				"items_criticos":  0,
				"top_proveedor":   "-",
			},
			"inventarios": gin.H{
				"valor_total":            0,
				"items_activos":          0,
				"top_inventario":         "-",
				"top_inventario_valor":   0,
				"rotacion":               0,
				"mas_vendido":            "-",
				"mas_vendido_id":         0,
				"mas_vendido_ingresos":   0,
			},
			"plan_compras": gin.H{
				"items_reposicion": 0,
				"costo_estimado":   0,
				"items_urgentes":   0,
				"dias_stock":       "-",
			},
		},
	})
}

// GET /api/get_sales_analytics.php?period=month
func (s *Server) getSalesAnalytics(c *gin.Context) {
	period := c.DefaultQuery("period", "month")
	var dateFilter string
	switch period {
	case "day":
		dateFilter = "DATE(created_at)=CURDATE()"
	case "week":
		dateFilter = "YEARWEEK(created_at)=YEARWEEK(NOW())"
	default:
		dateFilter = "MONTH(created_at)=MONTH(NOW())"
	}
	
	var totalOrders int
	var totalRevenue, totalCost, totalProfit, totalDelivery, avgTicket float64
	s.DB.QueryRow(`SELECT COUNT(*) FROM tuu_orders WHERE `+dateFilter+` AND payment_status='paid'`).Scan(&totalOrders)
	s.DB.QueryRow(`SELECT COALESCE(SUM(installment_amount),0) FROM tuu_orders WHERE `+dateFilter+` AND payment_status='paid'`).Scan(&totalRevenue)
	s.DB.QueryRow(`SELECT COALESCE(SUM(delivery_fee),0) FROM tuu_orders WHERE `+dateFilter+` AND payment_status='paid'`).Scan(&totalDelivery)
	if totalOrders > 0 {
		avgTicket = totalRevenue / float64(totalOrders)
	}
	
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"summary_kpis": gin.H{
				"total_orders":   totalOrders,
				"total_revenue":  totalRevenue,
				"total_cost":     totalCost,
				"total_profit":   totalProfit,
				"total_delivery": totalDelivery,
				"avg_ticket":     avgTicket,
			},
			"payment_summary": []gin.H{},
		},
	})
}

// GET /api/get_month_comparison.php
func (s *Server) getMonthComparison(c *gin.Context) {
	var currentMonth, previousMonth float64
	s.DB.QueryRow(`SELECT COALESCE(SUM(installment_amount),0) FROM tuu_orders WHERE MONTH(created_at)=MONTH(NOW()) AND payment_status='paid'`).Scan(&currentMonth)
	s.DB.QueryRow(`SELECT COALESCE(SUM(installment_amount),0) FROM tuu_orders WHERE MONTH(created_at)=MONTH(NOW())-1 AND payment_status='paid'`).Scan(&previousMonth)
	
	_ = currentMonth
	_ = previousMonth
	
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"currentMonth": gin.H{
				"salesByWeekday": []float64{},
			},
			"previousMonth": gin.H{
				"salesByWeekday": []float64{},
			},
		},
	})
}

// GET /api/get_financial_reports.php
func (s *Server) getFinancialReports(c *gin.Context) {
	rows, _ := s.DB.Query(`
		SELECT DATE_FORMAT(created_at, '%Y-%m') as month, 
			COUNT(*) as orders, 
			COALESCE(SUM(installment_amount),0) as revenue,
			COALESCE(SUM(delivery_fee),0) as delivery_fees
		FROM tuu_orders 
		WHERE payment_status='paid' AND created_at >= DATE_SUB(NOW(), INTERVAL 12 MONTH)
		GROUP BY month ORDER BY month DESC`)
	defer rows.Close()
	
	reports := []gin.H{}
	for rows.Next() {
		var month string
		var orders int
		var revenue, fees float64
		rows.Scan(&month, &orders, &revenue, &fees)
		reports = append(reports, gin.H{"month": month, "orders": orders, "revenue": revenue, "delivery_fees": fees, "net_revenue": revenue - fees})
	}
	c.JSON(200, gin.H{"success": true, "reports": reports})
}
