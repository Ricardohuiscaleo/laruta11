# MenuApp.jsx - Estado de Migración PHP → Go

## APIs Usadas en MenuApp.jsx (22 endpoints)

### ✅ YA MIGRADAS A GO (2/22 = 9%)

1. ✅ `/api/auth/check_session.php` → `GET /api/auth/check`
2. ✅ `/api/create_order.php` → `POST /api/orders` (existe en handlers_all.go)

### ❌ FALTAN MIGRAR (20/22 = 91%)

#### Auth & Users (2)
3. ❌ `/api/auth/delete_account.php` → `DELETE /api/auth/account`
4. ❌ `/api/update_cashier_profile.php` → `PUT /api/users/:id`

#### Products (3)
5. ❌ `/api/get_menu_products.php` → `GET /api/products?menu=true`
6. ❌ `/api/toggle_like.php` → `POST /api/products/:id/like`
7. ❌ `/api/toggle_product_status.php` → `PUT /api/products/:id/status`

#### Orders & Notifications (2)
8. ❌ `/api/get_user_orders.php` → `GET /api/orders?user_id=X`
9. ❌ `/api/get_order_notifications.php` → `GET /api/notifications`
10. ❌ `/api/notify_admin_payment.php` → `POST /api/notifications/admin`

#### Food Trucks (6)
11. ❌ `/api/get_nearby_trucks.php` → `POST /api/trucks/nearby`
12. ❌ `/api/get_truck_status.php` → `GET /api/trucks/:id/status`
13. ❌ `/api/get_truck_schedules.php` → `GET /api/trucks/:id/schedules`
14. ❌ `/api/update_truck_status.php` → `PUT /api/trucks/:id/status`
15. ❌ `/api/update_truck_config.php` → `PUT /api/trucks/:id/config`
16. ❌ `/api/update_truck_schedule.php` → `PUT /api/trucks/:id/schedules`

#### Location (5)
17. ❌ `/api/location/geocode.php` → `POST /api/location/geocode`
18. ❌ `/api/location/save_location.php` → `POST /api/location/save`
19. ❌ `/api/location/check_delivery_zone.php` → `POST /api/location/delivery`
20. ❌ `/api/location/get_nearby_products.php` → `POST /api/location/products`
21. ❌ `/api/location/calculate_delivery_time.php` → `POST /api/location/time`

#### Analytics (1)
22. ❌ `/api/track_usage.php` → `POST /api/analytics/track`

---

## Prioridad de Migración (Crítico → Bajo)

### 🔴 CRÍTICO (bloquea funcionalidad core)
- `/api/get_menu_products.php` - Sin esto no hay menú
- `/api/get_user_orders.php` - Historial de pedidos
- `/api/get_order_notifications.php` - Notificaciones en tiempo real

### 🟡 IMPORTANTE (afecta UX)
- `/api/toggle_product_status.php` - Admin toggle productos
- `/api/get_nearby_trucks.php` - Mostrar trucks cercanos
- `/api/location/check_delivery_zone.php` - Validar delivery
- `/api/location/geocode.php` - Convertir coords a dirección

### 🟢 OPCIONAL (features secundarias)
- `/api/toggle_like.php` - Likes de productos
- `/api/track_usage.php` - Analytics
- `/api/update_cashier_profile.php` - Perfil cajero
- Resto de endpoints trucks/location

---

## Plan de Consolidación Eficiente

### Módulo 1: Products (5 PHP → 3 Go)
```go
GET    /api/products?menu=true&active=true    // Reemplaza get_menu_products.php
POST   /api/products/:id/like                 // Reemplaza toggle_like.php
PUT    /api/products/:id/status               // Reemplaza toggle_product_status.php
```

### Módulo 2: Orders & Notifications (3 PHP → 2 Go)
```go
GET    /api/orders?user_id=X                  // Reemplaza get_user_orders.php
GET    /api/notifications?user_id=X           // Reemplaza get_order_notifications.php + notify_admin_payment.php
```

### Módulo 3: Trucks (6 PHP → 3 Go)
```go
GET    /api/trucks?nearby=true&lat=X&lng=Y    // Reemplaza get_nearby_trucks.php + get_truck_status.php
GET    /api/trucks/:id/schedules              // Reemplaza get_truck_schedules.php
PUT    /api/trucks/:id                        // Reemplaza update_truck_status.php + update_truck_config.php + update_truck_schedule.php
```

### Módulo 4: Location (5 PHP → 2 Go)
```go
POST   /api/location/geocode                  // Reemplaza geocode.php + save_location.php
POST   /api/location/delivery                 // Reemplaza check_delivery_zone.php + get_nearby_products.php + calculate_delivery_time.php
```

### Módulo 5: Users & Analytics (3 PHP → 2 Go)
```go
PUT    /api/users/:id                         // Reemplaza update_cashier_profile.php
DELETE /api/users/:id                         // Reemplaza auth/delete_account.php
POST   /api/analytics/track                   // Reemplaza track_usage.php
```

---

## Resultado Final

**Antes**: 22 endpoints PHP dispersos
**Después**: 12 endpoints Go consolidados
**Reducción**: 45% menos endpoints
**Código**: 22 archivos PHP → 1 archivo handlers_all.go

---

## Próximos Pasos

1. **Implementar Módulo 1 (Products)** - 3 endpoints críticos
2. **Implementar Módulo 2 (Orders/Notifications)** - 2 endpoints críticos
3. **Actualizar MenuApp.jsx** - Cambiar URLs a Go API
4. **Testing paralelo** - Comparar respuestas PHP vs Go
5. **Rollout gradual** - Feature flag `USE_GO_API=true`
6. **Deprecar PHP** - Mover a `/api-legacy-php/`

---

**Estado actual**: Solo 2/22 endpoints migrados (9%)
**Objetivo**: 12 endpoints consolidados (100% funcionalidad)
**Tiempo estimado**: 2-3 días de desarrollo
