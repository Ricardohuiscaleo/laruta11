package main

import (
	"database/sql"
	"encoding/json"
	"os"
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
		if (req.User == "ruta11caja" && req.Pass == os.Getenv("CAJA_USER_CAJERA")) || (req.User == "admin123" && req.Pass == os.Getenv("CAJA_USER_ADMIN")) {
			valid, role = true, "caja"
		}
	case "inventario":
		if req.User == os.Getenv("INVENTARIO_USER") && req.Pass == os.Getenv("INVENTARIO_PASSWORD") {
			valid, role = true, "inventario"
		}
	case "comandas":
		if req.User == "comandas" && req.Pass == os.Getenv("CAJA_USER_CAJERA") {
			valid, role = true, "comandas"
		}
	case "admin":
		if (req.User == "admin" && req.Pass == os.Getenv("ADMIN_USER_ADMIN")) || (req.User == "ricardo" && req.Pass == os.Getenv("ADMIN_USER_RICARDO")) || (req.User == "manager" && req.Pass == os.Getenv("ADMIN_USER_MANAGER")) || (req.User == "ruta11" && req.Pass == os.Getenv("ADMIN_USER_RUTA11")) {
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
	rows, _ := s.DB.Query(`SELECT id, name, is_active, display_order FROM categories ORDER BY display_order, name`)
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
	rows, _ := s.DB.Query(`SELECT id, date, type, items, completed FROM checklists WHERE date = ? ORDER BY type`, date)
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
	rows, _ := s.DB.Query(`SELECT id, order_number, customer_data, items_data, total_amount, order_status FROM tuu_orders WHERE order_status IN ('pending', 'preparing') ORDER BY created_at DESC`)
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
