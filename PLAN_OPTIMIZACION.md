# Plan de Optimización - Arquitectura Go

## Premisa: Eficiencia y Velocidad Máxima
- Menos código, menos endpoints
- Consolidar requests múltiples en uno solo
- Goroutines para queries paralelas
- Arquitectura sólida y rápida

---

## APIs Detectadas en admin/index.astro (41 endpoints)

### ✅ Ya Migradas a Go
1. `/api/products` - CRUD productos
2. `/api/orders` - CRUD órdenes
3. `/api/auth/login` - Autenticación
4. `/api/dashboard` - Dashboard consolidado

### 🔥 Alta Prioridad (Consolidar)

#### Grupo 1: Dashboard (8 endpoints → 1 endpoint)
```
❌ /api/get_dashboard_analytics.php
❌ /api/get_dashboard_cards.php
❌ /api/get_sales_analytics.php
❌ /api/get_month_comparison.php
❌ /api/get_previous_month_summary.php
❌ /api/get_smart_projection.php
❌ /api/get_quality_score.php
❌ /api/get_technical_report.php

✅ GET /api/dashboard?include=analytics,cards,sales,comparison,projection,quality,report
```

**Beneficio**: 8 requests → 1 request con goroutines paralelas

#### Grupo 2: Productos (7 endpoints → 1 endpoint)
```
❌ /api/get_productos.php
❌ /api/add_producto.php
❌ /api/delete_producto.php
❌ /api/bulk_update_products.php
❌ /api/bulk_delete_products.php
❌ /api/bulk_edit_products.php
❌ /api/bulk_adjust_price.php

✅ GET /api/products?include_inactive=1&status=all
✅ POST /api/products (crear)
✅ PUT /api/products/:id (editar)
✅ DELETE /api/products/:id (eliminar)
✅ POST /api/products/bulk (operaciones masivas)
```

**Beneficio**: Operaciones batch en 1 request

#### Grupo 3: Usuarios (2 endpoints → 1 endpoint)
```
❌ /api/users/get_users.php
❌ /api/get_user_details.php

✅ GET /api/users?include=orders,stats
✅ GET /api/users/:id?include=orders,stats
```

#### Grupo 4: Militares RL6 (2 endpoints → 1 endpoint)
```
❌ /api/get_militares_rl6.php
❌ /api/approve_militar_rl6.php

✅ GET /api/militares?status=pending
✅ POST /api/militares/:id/approve
```

#### Grupo 5: Combos (3 endpoints → 1 endpoint)
```
❌ /api/get_combos.php
❌ /api/save_combo.php
❌ /api/delete_combo.php

✅ GET /api/combos
✅ POST /api/combos
✅ DELETE /api/combos/:id
```

#### Grupo 6: Concurso (3 endpoints → 1 endpoint)
```
❌ /api/get_concurso_stats.php
❌ /api/get_participantes_concurso.php
❌ /api/delete_concursante.php

✅ GET /api/concurso?include=stats,participantes
✅ DELETE /api/concurso/participantes/:id
```

#### Grupo 7: Utilidades (4 endpoints → mantener)
```
✅ /api/upload_image.php → Mantener (S3)
✅ /api/admin_logout.php → Mantener (sesión)
✅ /api/check_tracking_data.php → Mantener (robots)
✅ /api/cleanup_fake_data.php → Mantener (testing)
```

---

## Estructura de Carpetas Propuesta

```
caja/api-go/
├── main.go                    # Server + routes
├── handlers/
│   ├── dashboard.go          # Dashboard consolidado (goroutines)
│   ├── products.go           # CRUD + bulk operations
│   ├── orders.go             # CRUD órdenes
│   ├── users.go              # CRUD usuarios
│   ├── militares.go          # Militares RL6
│   ├── combos.go             # Combos
│   ├── concurso.go           # Concurso
│   └── auth.go               # Login/logout
├── models/
│   ├── product.go
│   ├── order.go
│   ├── user.go
│   └── militar.go
├── middleware/
│   └── auth.go               # JWT validation
├── utils/
│   ├── db.go                 # Connection pool
│   ├── response.go           # JSON helpers
│   └── cache.go              # Redis (futuro)
├── Dockerfile
├── go.mod
└── go.sum
```

---

## Implementación por Fases

### Fase 1: Dashboard Consolidado (1 día)
**Objetivo**: 8 requests → 1 request

```go
// handlers/dashboard.go
func GetDashboard(c *gin.Context) {
    includes := c.Query("include") // analytics,cards,sales,comparison
    
    var wg sync.WaitGroup
    results := make(map[string]interface{})
    
    if strings.Contains(includes, "analytics") {
        wg.Add(1)
        go func() {
            defer wg.Done()
            results["analytics"] = fetchAnalytics()
        }()
    }
    
    if strings.Contains(includes, "cards") {
        wg.Add(1)
        go func() {
            defer wg.Done()
            results["cards"] = fetchCards()
        }()
    }
    
    // ... más goroutines
    
    wg.Wait()
    c.JSON(200, gin.H{"success": true, "data": results})
}
```

**Impacto**: Reducción de 80% en tiempo de carga del dashboard

### Fase 2: Productos Bulk (1 día)
**Objetivo**: Operaciones masivas en 1 request

```go
// handlers/products.go
func BulkProducts(c *gin.Context) {
    var req struct {
        Action string   `json:"action"` // activate, deactivate, delete, adjust_price
        IDs    []int    `json:"ids"`
        Data   map[string]interface{} `json:"data"`
    }
    
    c.BindJSON(&req)
    
    tx, _ := db.Begin()
    defer tx.Rollback()
    
    switch req.Action {
    case "activate":
        _, err := tx.Exec("UPDATE productos SET is_active=1 WHERE id IN (?)", req.IDs)
    case "adjust_price":
        amount := req.Data["amount"].(float64)
        _, err := tx.Exec("UPDATE productos SET price=price+? WHERE id IN (?)", amount, req.IDs)
    }
    
    tx.Commit()
    c.JSON(200, gin.H{"success": true})
}
```

### Fase 3: Usuarios + Militares (1 día)
**Objetivo**: Consolidar endpoints de usuarios

### Fase 4: Combos + Concurso (1 día)
**Objetivo**: Endpoints RESTful estándar

---

## Métricas de Éxito

### Antes (PHP)
- **41 endpoints** en admin/index.astro
- **8 requests** para cargar dashboard
- **~200ms** tiempo promedio por request
- **~1.6s** tiempo total dashboard

### Después (Go)
- **~15 endpoints** consolidados
- **1 request** para cargar dashboard
- **~50ms** tiempo promedio por request
- **~50ms** tiempo total dashboard (goroutines)

**Mejora**: 97% más rápido

---

## Próximos Pasos

1. ✅ Analizar admin/index.astro (HECHO)
2. ⏭️ Implementar dashboard consolidado
3. ⏭️ Implementar productos bulk
4. ⏭️ Migrar usuarios + militares
5. ⏭️ Migrar combos + concurso
6. ⏭️ Testing con API producción
7. ⏭️ Deploy

---

**Fecha**: 2026-02-10
**Estado**: Análisis completado, listo para implementar
