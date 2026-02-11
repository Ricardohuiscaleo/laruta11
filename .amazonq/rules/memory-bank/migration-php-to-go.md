# Migración PHP → Go APIs - Plan de Ejecución

## Estado Actual
- **PHP APIs**: ~600 archivos en `/app/api/` y `/caja/api/`
- **APIs en uso real**: **180 endpoints activos** (48 no usadas)
- **Total APIs detectadas**: 228
- **Clases/Helpers**: 2 (S3Manager, youtube_helper)
- **Webhooks**: 15
- **Tablas BD**: 73
- **Go APIs**: 2 servicios básicos (`api-go-landing`, `api-go-caja`)
- **Objetivo**: Consolidar 180 endpoints → ~50 endpoints Go optimizados

---

## Herramienta de Análisis

**Ubicación**: `/caja/src/pages/final.astro`

**Uso**:
```bash
cd caja
npm run dev
# Visitar: http://localhost:4321/final
```

**Detecta**:
- ✅ Total de APIs PHP: 228
- ✅ APIs usadas: 180
- ✅ APIs no usadas: 48
- ✅ Clases PHP (managers, helpers): 2
- ✅ Webhooks: 15
- ✅ Includes PHP entre archivos
- ✅ Tablas de base de datos usadas: 73
- ✅ Archivos frontend que usan cada API

**Resultado del análisis**: Ver `/caja/MIGRACION_APIS.md`

---

## Checklist de Migración

### ✅ Fase 0: Preparación (COMPLETADO)
- [x] Auditoría de endpoints PHP
- [x] Identificación de endpoints críticos
- [x] Definición de arquitectura Go
- [x] Plan de consolidación

### ✅ Fase 1: Infraestructura Base (COMPLETADO)
- [x] **1.1 Estructura de proyecto**
  - [x] Crear estructura base en raíz
  - [x] Setup `main.go` con Gin
  - [x] Configurar `go.mod` con dependencias
  
- [x] **1.2 Database Layer**
  - [x] Connection pool MySQL (max 25 conns)
  - [x] Helper functions (success, error)
  - [x] Error handling centralizado
  
- [ ] **1.3 Caching Layer**
  - [ ] Redis client setup
  - [ ] Cache middleware
  - [ ] TTL strategies (productos: 5min, órdenes: 30s)
  
- [x] **1.4 Middleware**
  - [x] CORS handler
  - [ ] JWT authentication (token simple implementado)
  - [ ] Rate limiting (100 req/min)
  - [x] Request logging (Gin default)

### 📦 Fase 2: Módulos Core (4 semanas)

#### Semana 1: Productos & Inventario
- [x] **2.1 Productos API**
  - [x] `GET /api/products` (con filtros: menu, stock, recipes)
  - [x] `GET /api/products/:id`
  - [x] `POST /api/products`
  - [x] `PUT /api/products/:id`
  - [x] `DELETE /api/products/:id`
  - [ ] Cache: 5 minutos
  - [ ] Tests: 10 casos

- [ ] **2.2 Inventario API**
  - [ ] `GET /api/inventory` (ingredientes + stock)
  - [ ] `PUT /api/inventory/adjust` (batch updates)
  - [ ] Cálculo automático de stock productos
  - [ ] Alertas de stock bajo
  - [ ] Tests: 8 casos

- [ ] **2.3 Frontend Update**
  - [ ] Cambiar URLs en componentes React
  - [ ] Feature flag `USE_GO_API=true`
  - [ ] Testing en dev

#### Semana 2: Órdenes
- [x] **2.4 Órdenes API**
  - [x] `GET /api/orders` (filtros: status, user, date)
  - [x] `GET /api/orders/:id`
  - [x] `POST /api/orders` (crear orden + items)
  - [x] `PUT /api/orders/:id/status`
  - [x] `DELETE /api/orders/:id`
  - [ ] Webhook notifications
  - [ ] Tests: 12 casos

- [ ] **2.5 Integración Inventario**
  - [ ] Deducción automática de stock
  - [ ] Rollback en cancelación
  - [ ] Audit trail
  - [ ] Tests: 6 casos

- [ ] **2.6 Frontend Update**
  - [ ] Actualizar componentes de órdenes
  - [ ] Testing flujo completo

#### Semana 3: Autenticación & Usuarios
- [x] **2.7 Auth API**
  - [x] `POST /api/auth/login` (email + password SHA256)
  - [x] `POST /api/auth/register`
  - [x] `POST /api/auth/refresh` (token simple)
  - [x] Session management básico
  - [ ] Google OAuth
  - [ ] Tests: 10 casos

- [ ] **2.8 Usuarios API**
  - [ ] `GET /api/users` (admin only)
  - [ ] `GET /api/users/:id`
  - [ ] `PUT /api/users/:id`
  - [ ] `DELETE /api/users/:id`
  - [ ] Tests: 8 casos

- [ ] **2.9 Frontend Update**
  - [ ] Migrar login/register
  - [ ] JWT storage
  - [ ] Auto-refresh token

#### Semana 4: Pagos & Analytics
- [ ] **2.10 Pagos TUU API**
  - [ ] `POST /api/payments` (crear pago TUU)
  - [ ] `POST /api/payments/webhook` (callback)
  - [ ] Validación de pagos
  - [ ] Retry logic
  - [ ] Tests: 8 casos

- [ ] **2.11 Analytics API**
  - [ ] `GET /api/analytics/dashboard`
  - [ ] `GET /api/analytics/sales`
  - [ ] `GET /api/analytics/financial`
  - [ ] Cache: 10 minutos
  - [ ] Tests: 6 casos

- [ ] **2.12 Frontend Update**
  - [ ] Migrar flujo de pagos
  - [ ] Dashboards analytics

### 🧪 Fase 3: Testing & Validación (1 semana)
- [ ] **3.1 Tests Paralelos**
  - [ ] Comparar respuestas PHP vs Go
  - [ ] Validar tiempos de respuesta
  - [ ] Edge cases testing

- [ ] **3.2 Performance Benchmarks**
  - [ ] Apache Bench: 1000 requests
  - [ ] Comparar PHP vs Go
  - [ ] Optimizar queries lentas

- [ ] **3.3 Rollout Gradual**
  - [ ] 10% tráfico → Go (feature flag)
  - [ ] Monitorear errores 24h
  - [ ] 50% tráfico → Go
  - [ ] Monitorear errores 48h
  - [ ] 100% tráfico → Go

### 🗑️ Fase 4: Cleanup (3 días)
- [ ] **4.1 Deprecar PHP**
  - [ ] Mover `/app/api/` → `/app/api-legacy-php/`
  - [ ] Mover `/caja/api/` → `/caja/api-legacy-php/`
  - [ ] Actualizar `.gitignore`

- [ ] **4.2 Documentación**
  - [ ] Actualizar README con nuevas rutas
  - [ ] Documentar endpoints Go (Swagger)
  - [ ] Actualizar memory bank

- [ ] **4.3 Optimizaciones Finales**
  - [ ] Query optimization
  - [ ] Index database
  - [ ] Monitoring setup (logs, metrics)

---

## Endpoints Consolidados

### Productos (50+ PHP → 5 Go)
```
GET    /api/products?menu=true&stock=true&recipes=true
GET    /api/products/:id
POST   /api/products
PUT    /api/products/:id
DELETE /api/products/:id
```

### Órdenes (30+ PHP → 5 Go)
```
GET    /api/orders?status=pending&user_id=123&date=2026-02-10
GET    /api/orders/:id
POST   /api/orders
PUT    /api/orders/:id/status
DELETE /api/orders/:id
```

### Compras (40+ PHP → 8 Go)
```
GET    /api/compras                    # Historial compras
GET    /api/compras/items              # Items disponibles
GET    /api/compras/proveedores        # Lista proveedores
GET    /api/compras/saldo              # Saldo disponible
GET    /api/compras/historial          # Historial saldo
POST   /api/compras                    # Registrar compra
DELETE /api/compras/:id                # Eliminar compra
POST   /api/compras/:id/respaldo       # Subir respaldo (S3)
```

### Ingredientes (3 PHP → 3 Go)
```
GET    /api/ingredientes               # Listar ingredientes
POST   /api/ingredientes               # Crear/actualizar
GET    /api/ingredientes/:id/precio    # Precio histórico
```

### Food Trucks (10+ PHP → 6 Go)
```
GET    /api/trucks                     # Listar trucks
GET    /api/trucks/:id                 # Detalle truck
PUT    /api/trucks/:id                 # Actualizar config
PUT    /api/trucks/:id/status          # Cambiar estado
GET    /api/trucks/:id/schedules       # Horarios
PUT    /api/trucks/:id/schedules       # Actualizar horarios
POST   /api/trucks/nearby              # Trucks cercanos
```

### Location (5 PHP → 5 Go)
```
POST   /api/location/geocode           # Geocodificar coords
POST   /api/location/save              # Guardar ubicación
POST   /api/location/delivery          # Verificar zona delivery
POST   /api/location/products          # Productos cercanos
POST   /api/location/time              # Calcular tiempo delivery
```

### Notificaciones (2 PHP → 2 Go)
```
GET    /api/notifications              # Listar notificaciones
POST   /api/notifications/admin        # Notificar admin
```

### Autenticación (20+ PHP → 3 Go)
```
POST /api/auth/login
POST /api/auth/register
POST /api/auth/refresh
```

### Usuarios (10+ PHP → 4 Go)
```
GET    /api/users
GET    /api/users/:id
PUT    /api/users/:id
DELETE /api/users/:id
```

### Pagos (50+ PHP → 2 Go)
```
POST /api/payments
POST /api/payments/webhook
```

### Analytics (30+ PHP → 3 Go)
```
GET /api/analytics/dashboard?period=week
GET /api/analytics/sales?group_by=product
GET /api/analytics/financial?report=monthly
```

---

## Métricas de Éxito

### Performance
- [ ] Tiempo respuesta promedio < 50ms (vs ~200ms PHP)
- [ ] Throughput > 1000 req/s (vs ~100 req/s PHP)
- [ ] Memory usage < 100MB (vs ~500MB PHP)

### Calidad
- [ ] 100% endpoints críticos migrados
- [ ] 0 errores en producción (primera semana)
- [ ] Test coverage > 80%

### Consolidación
- [ ] **168 APIs PHP activas** → ~50 endpoints Go
- [ ] Reducción 70% en endpoints
- [ ] Reducción 95% en archivos de código (600 archivos → 1 main.go)
- [ ] Mantenibilidad mejorada
- [ ] 24 APIs no usadas pueden eliminarse

## Módulos Identificados del Análisis Real

**APIs más críticas por uso:**

1. **Admin/Auth** - 15+ endpoints
   - admin_auth.php (8 usos)
   - check_admin_auth.php (7 usos)
   - auth/check_session.php (3 usos)
   - auth/login_v2.php, logout, etc.

2. **Productos** - 20+ endpoints
   - categories.php (8 usos)
   - get_menu_products.php (7 usos)
   - add_producto.php, delete_producto.php
   - bulk_update_products.php, bulk_delete_products.php

3. **Órdenes/Comandas** - 15+ endpoints
   - create_order.php (3 usos)
   - cancel_order.php (2 usos)
   - get_orders.php, update_order_status.php

4. **Compras** - 8 endpoints (✅ identificados)
   - compras/get_compras.php
   - compras/registrar_compra.php
   - compras/get_saldo_disponible.php
   - compras/get_items_compra.php

5. **Ingredientes/Inventario** - 10+ endpoints
   - get_ingredientes.php (5 usos)
   - save_ingrediente.php (4 usos)
   - delete_ingrediente.php

6. **Analytics/Tracking** - 10+ endpoints
   - app/track_visit.php (2 usos)
   - app/track_interaction.php
   - get_dashboard_analytics.php

7. **Pagos TUU** - 8+ endpoints
   - tuu_payment_gateway.php
   - confirm_transfer_payment.php
   - get_tuu_transactions.php

8. **Checklist/Calidad** - 4 endpoints
   - checklist.php (4 usos)

9. **Combos** - 3 endpoints
   - get_combos.php (3 usos)
   - save_combo.php, delete_combo.php

10. **Reviews** - 3 endpoints
    - add_review.php, get_reviews.php

**Total consolidado: ~50 endpoints Go** (vs 168 PHP activos)

---

## Riesgos & Mitigaciones

| Riesgo | Probabilidad | Impacto | Mitigación |
|--------|--------------|---------|------------|
| Endpoints PHP no documentados | Alta | Alto | Auditoría completa + logs de acceso |
| Diferencias en respuestas | Media | Alto | Tests de comparación automatizados |
| Downtime durante migración | Baja | Crítico | Feature flags + rollback rápido |
| Bugs en producción | Media | Alto | Rollout gradual 10%→50%→100% |
| Performance degradation | Baja | Medio | Benchmarks continuos + monitoring |

---

## APIs Identificadas en Frontend (NUEVO)

### MenuApp.jsx - Endpoints Críticos
```
✅ /api/get_menu_products.php → GET /api/products
⏭️ /api/toggle_product_status.php → PUT /api/products/:id/status
⏭️ /api/toggle_like.php → POST /api/products/:id/like
✅ /api/create_order.php → POST /api/orders
⏭️ /api/get_user_orders.php → GET /api/orders?user_id=X
⏭️ /api/get_order_notifications.php → GET /api/notifications
⏭️ /api/update_cashier_profile.php → PUT /api/users/:id
⏭️ /api/update_truck_status.php → PUT /api/trucks/:id/status
⏭️ /api/get_truck_status.php → GET /api/trucks/:id
⏭️ /api/get_nearby_trucks.php → POST /api/trucks/nearby
⏭️ /api/location/* → POST /api/location/*
✅ /api/auth/check_session.php → GET /api/auth/session
```

### ComprasApp.jsx - Endpoints Compras
```
⏭️ /api/compras/get_items_compra.php → GET /api/compras/items
⏭️ /api/compras/get_compras.php → GET /api/compras
⏭️ /api/compras/get_proveedores.php → GET /api/compras/proveedores
⏭️ /api/compras/get_saldo_disponible.php → GET /api/compras/saldo
⏭️ /api/compras/registrar_compra.php → POST /api/compras
⏭️ /api/compras/delete_compra.php → DELETE /api/compras/:id
⏭️ /api/compras/upload_respaldo.php → POST /api/compras/:id/respaldo
⏭️ /api/save_ingrediente.php → POST /api/ingredientes
```

### OrderPOSApp.jsx - Endpoints POS
```
✅ /api/create_order.php → POST /api/orders (duplicado)
⏭️ /api/tuu_payment_gateway.php → POST /api/payments/tuu
```

## Próximos Pasos Inmediatos

1. ✅ Auditoría completa de endpoints frontend
2. ⏭️ Implementar endpoints faltantes críticos:
   - PUT /api/products/:id/status
   - POST /api/products/:id/like
   - GET /api/notifications
   - PUT /api/users/:id
   - GET /api/trucks/:id + PUT status
3. ⏭️ Implementar módulo Compras completo (8 endpoints)
4. ⏭️ Implementar módulo Location (5 endpoints)
5. ⏭️ Testing + Frontend integration

---

## Notas de Implementación

### Dependencias Go
```go
github.com/gin-gonic/gin           // HTTP framework
github.com/go-sql-driver/mysql     // MySQL driver
github.com/redis/go-redis/v9       // Redis client
github.com/golang-jwt/jwt/v5       // JWT auth
github.com/joho/godotenv           // Env vars
```

### Variables de Entorno
```bash
APP_DB_HOST=websites_mysql-laruta11
APP_DB_NAME=laruta11
APP_DB_USER=laruta11_user
APP_DB_PASS=CCoonn22kk11@
REDIS_HOST=localhost:6379
JWT_SECRET=<generate-random>
PORT=3000
GIN_MODE=release
```

### Estructura de Respuesta Estándar
```json
{
  "success": true,
  "data": {...},
  "error": null,
  "meta": {
    "timestamp": "2026-02-10T20:00:00Z",
    "cached": false
  }
}
```

---

**Última actualización**: 2026-02-10
**Estado**: Fase 2 - Implementación Core
**Progreso**: 35% completado

## Archivos Creados
- `main.go` - Servidor principal con Gin + DB pooling
- `handlers_products.go` - CRUD productos completo
- `handlers_orders.go` - CRUD órdenes completo
- `handlers_auth.go` - Login/Register/Refresh
- `handlers_stubs.go` - Stubs para endpoints pendientes
- `go.mod` - Dependencias
- `Dockerfile` - Deploy Easypanel
