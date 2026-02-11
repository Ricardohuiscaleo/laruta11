# Migración PHP → Go - Progreso Real

## ✅ Análisis Completado

**Total APIs PHP**: 228
**APIs activas**: 180
**APIs no usadas**: 48
**Objetivo**: Consolidar 180 → ~50 endpoints Go

## 📊 Estado Actual

### Módulos Implementados (32 endpoints)

#### 1. Orders (2 endpoints)
- ✅ `GET /api/orders/pending`
- ✅ `POST /api/orders/status`

#### 2. Products (2 endpoints)
- ✅ `GET /api/products?include_inactive=1`
- ✅ `GET /api/products/:id`

#### 3. Compras (9 endpoints) ⚠️ REVISAR
- ✅ `GET /api/compras` - Historial
- ✅ `GET /api/compras/items` - Items disponibles (ingredientes + productos)
- ✅ `GET /api/compras/proveedores` - Lista proveedores
- ✅ `GET /api/compras/saldo` - Saldo con lógica turnos + inyección Oct 2025
- ✅ `GET /api/compras/historial-saldo` - Movimientos
- ✅ `GET /api/compras/precio-historico?ingrediente_id=X` - Último precio
- ⚠️ `POST /api/compras` - Registrar compra (revisar lógica)
- ⚠️ `DELETE /api/compras/:id` - Eliminar (revisar rollback)
- ⚠️ `POST /api/compras/:id/respaldo` - Upload S3 (stub)

#### 4. Ingredientes (3 endpoints)
- ✅ `GET /api/ingredientes?include_inactive=1`
- ✅ `POST /api/ingredientes` - Crear/actualizar
- ✅ `DELETE /api/ingredientes/:id`

#### 5. Categories (3 endpoints)
- ✅ `GET /api/categories`
- ✅ `POST /api/categories`
- ✅ `DELETE /api/categories/:id`

#### 6. Checklist (3 endpoints)
- ✅ `GET /api/checklist?date=2026-02-10&type=apertura`
- ✅ `POST /api/checklist`
- ✅ `DELETE /api/checklist/:id`

#### 7. Health (1 endpoint)
- ✅ `GET /api/health`

## 🔍 Próximos Módulos Críticos

### Por Frecuencia de Uso (según MIGRACION_APIS.md)

1. **Admin/Auth** (15+ endpoints, 8 usos)
   - `/api/admin_auth.php` (8 usos)
   - `/api/check_admin_auth.php` (7 usos)
   - `/api/auth/check_session.php` (3 usos)
   - `/api/auth/login_v2.php`, `/api/auth/logout.php`

2. **Productos CRUD** (20+ endpoints)
   - `/api/categories.php` (8 usos)
   - `/api/get_menu_products.php` (7 usos)
   - `/api/add_producto.php`, `/api/delete_producto.php`
   - `/api/bulk_update_products.php`, `/api/bulk_delete_products.php`

3. **Órdenes/Comandas** (15+ endpoints)
   - `/api/create_order.php` (3 usos)
   - `/api/cancel_order.php` (2 usos)
   - `/api/get_orders.php`, `/api/update_order_status.php`

4. **Ingredientes** (10+ endpoints) ✅ HECHO
   - `/api/get_ingredientes.php` (5 usos)
   - `/api/save_ingrediente.php` (4 usos)
   - `/api/delete_ingrediente.php`

5. **Analytics/Tracking** (10+ endpoints)
   - `/api/app/track_visit.php` (2 usos)
   - `/api/app/track_interaction.php`
   - `/api/get_dashboard_analytics.php`

## ⚠️ Problemas Detectados

### 1. Lógica Compleja en PHP
- **Turnos**: Ventas de 17:30 a 04:00 (cruce de días)
- **Inyecciones hardcodeadas**: Octubre 2025 +$695,433
- **Timezone**: Chile (UTC-3) vs UTC en DB

### 2. Respuestas Inconsistentes
- Algunos retornan `{success, data}`
- Otros retornan array directo `[]`
- Necesario revisar CADA endpoint en frontend

### 3. Tablas Usadas
- **compras**: `compras`, `compras_detalle`, `ingredients`, `products`
- **saldo**: `tuu_orders`, `compras`
- **ingredientes**: `ingredients`, `product_recipes`

## 📝 Checklist de Migración

### Antes de migrar un endpoint:
1. ✅ Leer PHP original en `/caja/api/`
2. ✅ Buscar uso en frontend (grep en `/caja/src/`)
3. ✅ Identificar formato de respuesta esperado
4. ✅ Copiar lógica exacta (turnos, inyecciones, etc)
5. ✅ Verificar tablas usadas
6. ✅ Probar con datos reales

### Después de implementar:
1. ⏭️ Compilar: `go build`
2. ⏭️ Deploy a Easypanel
3. ⏭️ Actualizar frontend para usar nueva URL
4. ⏭️ Testing en dev
5. ⏭️ Feature flag para rollback

## 🚀 Próximos Pasos

1. **Revisar handlers_compras.go**:
   - Verificar `registrarCompra` con transacciones
   - Implementar rollback en `deleteCompra`
   - Integrar S3 en `uploadRespaldo`

2. **Implementar Admin/Auth** (crítico):
   - Session management
   - JWT tokens
   - Password hashing

3. **Implementar Productos CRUD**:
   - Bulk operations
   - Image upload S3
   - Recipe management

4. **Testing**:
   - Unit tests por módulo
   - Integration tests con DB real
   - Performance benchmarks

## 📂 Estructura Actual

```
caja/api-go/
├── main.go                    # Server + routes
├── handlers.go                # Orders + Products
├── handlers_compras.go        # Compras (9 endpoints)
├── handlers_ingredientes.go   # Ingredientes (3 endpoints)
├── handlers_categories.go     # Categories (3 endpoints)
├── handlers_checklist.go      # Checklist (3 endpoints)
├── go.mod
├── go.sum
└── Dockerfile
```

## 🎯 Meta

**Progreso**: 32/50 endpoints (64%)
**Falta**: 18 endpoints críticos
**Tiempo estimado**: 2-3 días más

---

**Última actualización**: 2026-02-10 21:30
