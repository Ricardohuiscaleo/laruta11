package main

import (
	"database/sql"
	"time"
	"github.com/gin-gonic/gin"
)

// GET /api/compras
func (s *Server) getCompras(c *gin.Context) {
	rows, _ := s.db.Query(`SELECT id, fecha_compra, proveedor, monto_total, metodo_pago, notas FROM compras ORDER BY fecha_compra DESC LIMIT 50`)
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

// GET /api/compras/items
func (s *Server) getComprasItems(c *gin.Context) {
	rows, _ := s.db.Query(`
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

// GET /api/compras/proveedores
func (s *Server) getProveedores(c *gin.Context) {
	rows, _ := s.db.Query(`SELECT proveedor, COUNT(*) FROM compras GROUP BY proveedor ORDER BY COUNT(*) DESC`)
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

// GET /api/compras/saldo
func (s *Server) getSaldoDisponible(c *gin.Context) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	start1 := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 17, 30, 0, 0, time.UTC)
	end1 := time.Date(lastMonth.Year(), lastMonth.Month()+1, 1, 4, 0, 0, 0, time.UTC)
	start2 := time.Date(now.Year(), now.Month(), 1, 17, 30, 0, 0, time.UTC)

	var v1, v2, comp float64
	s.db.QueryRow(`SELECT COALESCE(SUM(installment_amount - COALESCE(delivery_fee, 0)), 0) FROM tuu_orders WHERE payment_status = 'paid' AND created_at >= ? AND created_at < ?`, start1, end1).Scan(&v1)
	s.db.QueryRow(`SELECT COALESCE(SUM(installment_amount - COALESCE(delivery_fee, 0)), 0) FROM tuu_orders WHERE payment_status = 'paid' AND created_at >= ?`, start2).Scan(&v2)
	s.db.QueryRow(`SELECT COALESCE(SUM(monto_total), 0) FROM compras WHERE DATE_FORMAT(fecha_compra, '%Y-%m') = ?`, now.Format("2006-01")).Scan(&comp)

	if lastMonth.Format("2006-01") == "2025-10" {
		v1 += 695433
	}

	saldo := v1 + v2 - 1590000 - comp
	c.JSON(200, gin.H{"success": true, "saldo_disponible": saldo, "ventas_mes_anterior": v1, "ventas_mes_actual": v2, "sueldos": 1590000, "compras_mes": comp})
}

// GET /api/compras/historial-saldo
func (s *Server) getHistorialSaldo(c *gin.Context) {
	rows, _ := s.db.Query(`
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

// GET /api/compras/precio-historico?ingrediente_id=X
func (s *Server) getPrecioHistorico(c *gin.Context) {
	id := c.Query("ingrediente_id")
	var precio, cant, sub float64
	var unidad, fecha, prov string
	err := s.db.QueryRow(`
		SELECT cd.precio_unitario, cd.cantidad, cd.unidad, cd.subtotal, c.fecha_compra, c.proveedor
		FROM compras_detalle cd JOIN compras c ON cd.compra_id = c.id
		WHERE cd.ingrediente_id = ? ORDER BY c.fecha_compra DESC LIMIT 1`, id).Scan(&precio, &cant, &unidad, &sub, &fecha, &prov)

	if err == sql.ErrNoRows {
		c.JSON(200, gin.H{"success": false, "error": "Sin historial"})
		return
	}
	c.JSON(200, gin.H{"success": true, "precio_unitario": precio, "unidad": unidad, "ultima_cantidad": cant, "ultimo_subtotal": sub, "fecha_compra": fecha, "proveedor": prov})
}

// POST /api/compras
func (s *Server) registrarCompra(c *gin.Context) {
	var req map[string]interface{}
	c.BindJSON(&req)

	tx, _ := s.db.Begin()
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

// DELETE /api/compras/:id
func (s *Server) deleteCompra(c *gin.Context) {
	id := c.Param("id")
	tx, _ := s.db.Begin()
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

// POST /api/compras/:id/respaldo
func (s *Server) uploadRespaldo(c *gin.Context) {
	id := c.Param("id")
	file, _ := c.FormFile("image")
	// TODO: S3 upload
	s.db.Exec(`UPDATE compras SET imagen_respaldo = ? WHERE id = ?`, file.Filename, id)
	c.JSON(200, gin.H{"success": true})
}
