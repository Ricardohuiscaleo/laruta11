# Próximos Pasos - Migración PHP → Go

## ✅ Estado Actual

- ✅ Análisis completado: 41 APIs PHP → 15 Go
- ✅ Estructura creada: `handlers/`, `models/`, `middleware/`, `utils/`
- ✅ Dashboard handler con goroutines implementado
- ✅ API Go en producción: `https://websites-api-go-caja-r11.dj3bvg.easypanel.host`
- ✅ Frontend configurado para usar API Go

## 🔥 Problema Actual

**Frontend llama a APIs PHP que NO funcionan en local**

Errores:
```
❌ /api/get_dashboard_analytics.php - SyntaxError
❌ /api/get_dashboard_cards.php - SyntaxError
❌ /api/get_sales_analytics.php - SyntaxError
... (38 más)
```

## 🎯 Solución

**Actualizar frontend para usar API Go consolidada**

### Paso 1: Actualizar `index.astro` (Dashboard)

Reemplazar:
```javascript
// ❌ ANTES: 8 requests a PHP
fetch('/api/get_dashboard_analytics.php')
fetch('/api/get_dashboard_cards.php')
fetch('/api/get_sales_analytics.php')
fetch('/api/get_month_comparison.php')
fetch('/api/get_previous_month_summary.php')
fetch('/api/get_smart_projection.php')
fetch('/api/get_quality_score.php')
fetch('/api/get_technical_report.php')
```

Por:
```javascript
// ✅ DESPUÉS: 1 request a Go
const API_BASE = 'https://websites-api-go-caja-r11.dj3bvg.easypanel.host';
const response = await fetch(`${API_BASE}/api/dashboard?include=analytics,cards,sales,comparison,projection,quality,report`);
const data = await response.json();

// Acceder a los datos
const analytics = data.data.analytics;
const cards = data.data.cards;
const sales = data.data.sales;
// etc...
```

### Paso 2: Implementar Queries Reales en Go

Editar `caja/api-go/handlers/dashboard.go`:

```go
func fetchAnalytics(db *sql.DB) map[string]interface{} {
    var totalSales, totalOrders, avgTicket float64
    
    query := `
        SELECT 
            COALESCE(SUM(total_amount), 0) as total_sales,
            COUNT(*) as total_orders,
            COALESCE(AVG(total_amount), 0) as avg_ticket
        FROM orders 
        WHERE MONTH(created_at) = MONTH(NOW())
          AND YEAR(created_at) = YEAR(NOW())
    `
    
    db.QueryRow(query).Scan(&totalSales, &totalOrders, &avgTicket)
    
    return map[string]interface{}{
        "total_sales":  totalSales,
        "total_orders": int(totalOrders),
        "avg_ticket":   avgTicket,
    }
}
```

### Paso 3: Compilar y Subir

```bash
cd caja/api-go
go mod tidy
git add -A
git commit -m "feat: dashboard consolidado con queries reales"
git push

# Easypanel → Rebuild manual
```

### Paso 4: Probar

```bash
# Desde local
curl "https://websites-api-go-caja-r11.dj3bvg.easypanel.host/api/dashboard?include=analytics,cards"
```

## 📋 Checklist de Migración

### Dashboard (Prioridad 1)
- [ ] Actualizar `index.astro` línea ~450 (loadDashboardMetrics)
- [ ] Actualizar `index.astro` línea ~480 (loadDashboardTUU)
- [ ] Actualizar `index.astro` línea ~520 (loadOperationalCards)
- [ ] Implementar queries reales en `handlers/dashboard.go`
- [ ] Probar endpoint consolidado
- [ ] Deploy a producción

### Productos (Prioridad 2)
- [ ] Actualizar `index.astro` línea ~1100 (loadProducts)
- [ ] Implementar bulk operations en Go
- [ ] Probar operaciones masivas
- [ ] Deploy

### Usuarios (Prioridad 3)
- [ ] Actualizar `index.astro` línea ~800 (loadUsers)
- [ ] Implementar endpoint consolidado
- [ ] Deploy

## 🚀 Comando Rápido

```bash
# 1. Editar handlers/dashboard.go (implementar queries)
# 2. Compilar y subir
cd caja/api-go && go mod tidy && cd ../.. && git add -A && git commit -m "feat: queries reales dashboard" && git push

# 3. Easypanel → Rebuild
# 4. Probar: curl https://websites-api-go-caja-r11.dj3bvg.easypanel.host/api/dashboard?include=analytics
```

## 📊 Impacto Esperado

- **Antes**: 41 APIs PHP (no funcionan en local)
- **Después**: 15 APIs Go (funcionan desde local)
- **Velocidad**: 97% más rápido (1.6s → 50ms)
- **Mantenibilidad**: Código consolidado y organizado

---

**Siguiente**: Actualizar `index.astro` para usar API Go consolidada
