# 🏗️ Arquitectura del Proyecto - La Ruta 11

Mapa completo de páginas, APIs y tablas para migración a VPS.

---

## 📱 PÁGINAS USUARIO (Frontend)

### 1. **index.astro** - Menú Principal APP
**Ruta:** `/`  
**Componente:** `MenuApp.jsx`

**APIs que usa:**
- `/api/app/track_visit.php` → Registra visitas
- `/api/app/track_interaction.php` → Registra clicks
- `/api/app/track_journey.php` → Registra navegación
- `/api/get_productos.php` → Obtiene menú

**Tablas BD:**
- `visits` (analytics)
- `interactions` (analytics)
- `journey` (analytics)
- `productos` (menú)
- `recetas` (ingredientes por producto)
- `ingredientes` (stock)

---

### 2. **checkout.astro** - Proceso de Pago
**Ruta:** `/checkout`  
**Componente:** `CheckoutApp.jsx`

**APIs que usa:**
- `/api/tuu/create_payment_working.php` → Crea pago TUU
- `/api/tuu/callback.php` → Confirma pago
- `/api/registrar_venta.php` → Guarda venta
- `/api/process_sale_inventory.php` → Descuenta inventario

**Tablas BD:**
- `tuu_orders` (pagos online)
- `ventas` (registro ventas)
- `ingredientes` (descuento stock)
- `recetas` (cálculo ingredientes)

---

## 👨‍💼 PÁGINAS ADMIN (Backend)

### 3. **admin/index.astro** - Panel Principal Admin
**Ruta:** `/admin`  
**Componente:** `AdminSPA.jsx`

**APIs que usa:**
- `/api/get_dashboard_kpis.php` → KPIs principales
- `/api/get_productos.php` → Lista productos
- `/api/users/get_users.php` → Lista usuarios
- `/api/get_sales_analytics.php` → Analytics ventas
- `/api/get_dashboard_cards.php` → Tarjetas operativas
- `/api/get_smart_projection.php` → Proyecciones

**Tablas BD:**
- `productos` (gestión menú)
- `usuarios` (clientes)
- `ventas` (historial)
- `tuu_orders` (pagos)
- `ingredientes` (inventario)
- `compras` (gastos)

---

### 4. **admin/dashboard.astro** - Dashboard Métricas
**Ruta:** `/admin/dashboard`

**APIs que usa:**
- `/api/tuu/get_from_mysql.php` → Transacciones TUU
- `/api/app/get_user_behavior.php` → Comportamiento usuarios
- `/api/app/get_analytics.php` → Analytics general

**Tablas BD:**
- `tuu_orders` (pagos)
- `visits` (visitas)
- `interactions` (interacciones)

---

### 5. **admin/ingredients.astro** - Gestión Ingredientes
**Ruta:** `/admin/ingredients`

**APIs que usa:**
- `/api/get_ingredientes.php` → Lista ingredientes
- `/api/save_ingrediente.php` → Crear ingrediente
- `/api/update_ingrediente.php` → Actualizar ingrediente
- `/api/delete_ingrediente.php` → Eliminar ingrediente

**Tablas BD:**
- `ingredientes` (stock)
- `recetas` (relación producto-ingrediente)

---

### 6. **admin/calidad.astro** - Control de Calidad
**Ruta:** `/admin/calidad`  
**Componente:** `ChecklistApp.jsx`

**APIs que usa:**
- `/api/get_questions.php` → Preguntas checklist
- `/api/save_checklist.php` → Guardar checklist
- `/api/get_quality_score.php` → Score calidad

**Tablas BD:**
- `quality_questions` (preguntas)
- `quality_checklists` (respuestas)

---

### 7. **compras.astro** - Sistema de Compras
**Ruta:** `/compras`  
**Componente:** `ComprasApp.jsx`

**APIs que usa:**
- `/api/compras/get_compras.php` → Historial compras
- `/api/compras/registrar_compra.php` → Nueva compra
- `/api/compras/get_proveedores.php` → Lista proveedores
- `/api/compras/get_saldo_disponible.php` → Saldo caja

**Tablas BD:**
- `compras` (registro compras)
- `compras_items` (detalle items)
- `proveedores` (proveedores)
- `saldo_caja` (flujo efectivo)
- `ingredientes` (actualiza stock)

---

### 8. **ventas-detalle.astro** - Detalle Ventas por Turno
**Ruta:** `/ventas-detalle`  
**Componente:** `VentasDetalle.jsx`

**APIs que usa:**
- `/api/get_sales_detail.php` → Detalle ventas
- `/api/get_ventas_turno.php` → Ventas por turno

**Tablas BD:**
- `ventas` (registro ventas)
- `tuu_orders` (pagos online)
- `productos` (info productos)

---

## 🔌 GRUPOS DE APIs POR FUNCIONALIDAD

### **Productos** (4 APIs)
```
/api/get_productos.php
/api/add_producto.php
/api/update_producto.php
/api/delete_producto.php
```
**Tablas:** `productos`, `recetas`, `ingredientes`

---

### **Ventas** (4 APIs)
```
/api/registrar_venta.php
/api/get_sales_analytics.php
/api/get_sales_detail.php
/api/process_sale_inventory.php
```
**Tablas:** `ventas`, `tuu_orders`, `ingredientes`, `recetas`

---

### **TUU Pagos** (4 APIs)
```
/api/tuu/create_payment_working.php
/api/tuu/callback.php
/api/tuu/get_from_mysql.php
/api/tuu/sync_haulmer_data.php
```
**Tablas:** `tuu_orders`

---

### **Analytics** (5 APIs)
```
/api/app/track_visit.php
/api/app/track_interaction.php
/api/app/track_journey.php
/api/app/get_analytics.php
/api/app/get_user_behavior.php
```
**Tablas:** `visits`, `interactions`, `journey`

---

### **Compras** (4 APIs)
```
/api/compras/get_compras.php
/api/compras/registrar_compra.php
/api/compras/get_proveedores.php
/api/compras/get_saldo_disponible.php
```
**Tablas:** `compras`, `compras_items`, `proveedores`, `saldo_caja`, `ingredientes`

---

### **Calidad** (3 APIs)
```
/api/get_questions.php
/api/save_checklist.php
/api/get_quality_score.php
```
**Tablas:** `quality_questions`, `quality_checklists`

---

### **Dashboard** (3 APIs)
```
/api/get_dashboard_kpis.php
/api/get_dashboard_cards.php
/api/get_smart_projection.php
```
**Tablas:** `ventas`, `tuu_orders`, `ingredientes`, `compras`

---

## 🗄️ TABLAS DE BASE DE DATOS

### **Tablas Principales (11)**
1. `productos` - Menú de productos
2. `ingredientes` - Inventario de ingredientes
3. `recetas` - Relación producto-ingrediente
4. `ventas` - Registro de ventas
5. `tuu_orders` - Pagos online TUU
6. `compras` - Registro de compras
7. `compras_items` - Detalle items comprados
8. `proveedores` - Proveedores
9. `usuarios` - Clientes registrados
10. `quality_questions` - Preguntas calidad
11. `quality_checklists` - Respuestas calidad

### **Tablas Analytics (3)**
1. `visits` - Visitas a la app
2. `interactions` - Clicks e interacciones
3. `journey` - Navegación usuarios

### **Tablas Auxiliares (3)**
1. `saldo_caja` - Flujo de efectivo
2. `categories` - Categorías productos
3. `subcategories` - Subcategorías

---

## 📋 CHECKLIST MIGRACIÓN VPS

### **Paso 1: Migrar Base de Datos**
```bash
# Exportar desde Hostinger
mysqldump -h srv1438.hstgr.io -u u958525313_app -p u958525313_app > ruta11.sql

# Importar en VPS
mysql -u usuario_vps -p ruta11_db < ruta11.sql
```

**Tablas críticas a verificar:**
- ✅ `productos` (menú completo)
- ✅ `ingredientes` (stock actual)
- ✅ `ventas` (historial)
- ✅ `tuu_orders` (pagos)

---

### **Paso 2: Actualizar config.php**
```php
// Cambiar credenciales BD
$host = 'localhost'; // o IP VPS
$dbname = 'ruta11_db';
$username = 'usuario_vps';
$password = 'password_vps';
```

---

### **Paso 3: Subir APIs PHP**
Copiar carpeta `/api/` completa al VPS:
```bash
scp -r api/ root@VPS_IP:/var/www/html/
```

---

### **Paso 4: Build y Deploy Frontend**
```bash
npm run build
# Subir carpeta dist/ a VPS
```

---

## 🔗 DEPENDENCIAS ENTRE SISTEMAS

```
index.astro (APP)
    ↓
get_productos.php
    ↓
productos + recetas + ingredientes
```

```
checkout.astro (PAGO)
    ↓
create_payment_working.php → registrar_venta.php → process_sale_inventory.php
    ↓
tuu_orders + ventas + ingredientes
```

```
admin/index.astro (ADMIN)
    ↓
get_dashboard_kpis.php + get_sales_analytics.php
    ↓
ventas + tuu_orders + ingredientes + compras
```

---

## ⚠️ APIS CRÍTICAS (NO PUEDEN FALLAR)

1. **get_productos.php** - Sin esto no hay menú
2. **create_payment_working.php** - Sin esto no hay pagos
3. **registrar_venta.php** - Sin esto no se registran ventas
4. **process_sale_inventory.php** - Sin esto no se descuenta stock

---

## 📊 RESUMEN COMPLETO

- **45+ páginas totales**
- **80+ APIs PHP**
- **20+ tablas de BD**
- **15+ componentes React**

### Páginas por Módulo:
- **Usuario (8):** index, checkout, login, success, payment-success, card-pending, cash-pending, transfer-pending
- **Admin (22):** dashboard, productos, usuarios, ventas, compras, ingredientes, calidad, combos, mermas, etc.
- **Inventario (3):** inventario, arqueo, arqueo-resumen
- **Comandas (1):** comandas/index
- **Jobs (5):** aplicaciones de trabajo
- **JobsTracker (7):** sistema de reclutamiento completo
- **Concurso (5):** sistema de concurso (deshabilitado)

**Base de datos:** `u958525313_app` (Hostinger) → `ruta11_db` (VPS)

---

**Última actualización:** 2025-01-XX  
**Para migración a:** VPS con Easypanel


---

## 📦 PÁGINAS ADICIONALES (Faltantes en diagrama)

### 9. **login.astro** - Login Usuario
**Ruta:** `/login`  
**APIs:** `/api/auth/login.php`, `/api/auth/check_session.php`  
**Tablas:** `usuarios`

### 10. **arqueo.astro** - Arqueo de Caja
**Ruta:** `/arqueo`  
**Componente:** `ArqueoApp.jsx`  
**APIs:** `/api/close_cash_register.php`, `/api/get_cash_register_status.php`  
**Tablas:** `cash_register`, `ventas`, `tuu_orders`

### 11. **mermas.astro** - Registro de Mermas
**Ruta:** `/mermas`  
**Componente:** `MermasApp.jsx`  
**APIs:** `/api/registrar_merma.php`, `/api/get_mermas.php`  
**Tablas:** `mermas`, `ingredientes`

### 12. **comandas/index.astro** - Sistema de Comandas
**Ruta:** `/comandas`  
**APIs:** `/api/get_pending_orders.php`, `/api/update_order_status.php`  
**Tablas:** `ventas`, `tuu_orders`

### 13. **inventario/index.astro** - Control de Inventario
**Ruta:** `/inventario`  
**APIs:** `/api/get_ingredientes.php`, `/api/update_ingredient_stock.php`  
**Tablas:** `ingredientes`, `recetas`

---

## 🏢 MÓDULO ADMIN (22 páginas)

### 14. **admin/analytics.astro** - Analytics Avanzado
**APIs:** `/api/app/get_analytics.php`, `/api/app/get_user_behavior.php`  
**Tablas:** `visits`, `interactions`, `journey`

### 15. **admin/combos.astro** - Gestión de Combos
**APIs:** `/api/get_combos.php`, `/api/save_combo.php`, `/api/delete_combo.php`  
**Tablas:** `combos`, `combo_items`, `combo_selections`

### 16. **admin/edit-product.astro** - Editor de Productos
**APIs:** `/api/update_producto.php`, `/api/get_product_recipe.php`, `/api/save_product_recipe.php`  
**Tablas:** `productos`, `recetas`, `ingredientes`

### 17. **admin/mermas.astro** - Gestión de Mermas
**APIs:** `/api/get_mermas.php`, `/api/registrar_merma.php`  
**Tablas:** `mermas`, `ingredientes`

### 18. **admin/pagos-tuu.astro** - Reportes TUU
**APIs:** `/api/tuu/get_from_mysql.php`, `/api/tuu/get_reports.php`  
**Tablas:** `tuu_orders`

### 19. **admin/reportes.astro** - Reportes Generales
**APIs:** `/api/get_financial_reports.php`, `/api/get_inventory_report.php`  
**Tablas:** `ventas`, `compras`, `ingredientes`

### 20. **admin/food-trucks.astro** - Gestión Food Trucks
**APIs:** `/api/food_trucks/get_all.php`, `/api/food_trucks/save.php`  
**Tablas:** `food_trucks`

### 21. **admin/concurso-stats.astro** - Estadísticas Concurso
**APIs:** `/api/get_concurso_stats.php`, `/api/track_concurso_visit.php`  
**Tablas:** `concurso_tracking`, `concurso_participantes`

### 22. **admin/keys.astro** - Gestión de API Keys
**APIs:** `/api/admin/get_keys.php`, `/api/admin/verify_keys_access.php`  
**Tablas:** `api_keys`

### 23. **admin/users.astro** - Gestión de Usuarios
**APIs:** `/api/users/get_users.php`, `/api/users/get_user_detail.php`  
**Tablas:** `usuarios`

### 24. **admin/test.astro** - Testing de APIs
**APIs:** Múltiples APIs de testing  
**Tablas:** Varias

### 25. **admin/technical-report.astro** - Informe Técnico
**APIs:** `/api/get_technical_report.php`  
**Tablas:** N/A (análisis de archivos)

---

## 💼 MÓDULO JOBS (Sistema de Aplicaciones de Trabajo)

### 26. **jobs/index.astro** - Portal de Trabajos
**APIs:** `/api/jobs/get_keywords.php`  
**Tablas:** `jobs_keywords`

### 27. **jobs/cajero.astro** - Aplicación Cajero
**APIs:** `/api/jobs/start_application.php`, `/api/jobs/submit_application.php`  
**Tablas:** `jobs_applications`, `jobs_keywords`

### 28. **jobs/maestro-sanguchero.astro** - Aplicación Maestro
**APIs:** `/api/jobs/start_application.php`, `/api/jobs/submit_application.php`  
**Tablas:** `jobs_applications`

---

## 🎯 MÓDULO JOBSTRACKER (Sistema de Reclutamiento)

### 29. **jobsTracker/dashboard/index.astro** - Dashboard Reclutamiento
**APIs:** `/api/tracker/get_dashboard_stats.php`, `/api/tracker/get_candidates.php`  
**Tablas:** `tracker_candidates`, `tracker_interviews`

### 30. **jobsTracker/kanban/index.astro** - Kanban de Candidatos
**APIs:** `/api/tracker/get_kanban.php`, `/api/tracker/move_kanban_card.php`  
**Tablas:** `tracker_candidates`, `tracker_kanban_status`

### 31. **jobsTracker/entrevista/index.astro** - Sistema de Entrevistas
**APIs:** `/api/tracker/get_interview.php`, `/api/tracker/save_interview.php`  
**Tablas:** `tracker_interviews`, `tracker_questions`

### 32. **jobsTracker/candidate/[id].astro** - Detalle Candidato
**APIs:** `/api/tracker/get_candidate_detail.php`  
**Tablas:** `tracker_candidates`, `tracker_interviews`

### 33. **jobsTracker/keywords/index.astro** - Gestión Keywords
**APIs:** `/api/tracker/get_keywords.php`, `/api/tracker/save_keywords.php`  
**Tablas:** `tracker_keywords`

### 34. **jobsTracker/qr/index.astro** - QR Codes
**APIs:** `/api/tracker/get_qr_locations.php`, `/api/tracker/track_qr_view.php`  
**Tablas:** `tracker_qr_locations`, `tracker_qr_views`

---

## 🎪 MÓDULO CONCURSO (Deshabilitado pero presente)

### 35. **concurso_disabled/index.astro** - Landing Concurso
**APIs:** `/api/track_concurso_visit.php`, `/api/get_participantes_concurso.php`  
**Tablas:** `concurso_tracking`, `concurso_participantes`

### 36. **concurso_disabled/admin.astro** - Admin Concurso
**APIs:** `/api/update_concurso_state.php`, `/api/add_concursante_manual.php`  
**Tablas:** `concurso_state`, `concurso_participantes`

### 37. **concurso_disabled/live.astro** - Vista EN VIVO
**APIs:** `/api/get_concurso_live.php`  
**Tablas:** `concurso_state`

---

## 🗄️ TABLAS ADICIONALES DE BASE DE DATOS

### **Tablas Combos (3)**
1. `combos` - Definición de combos
2. `combo_items` - Productos en cada combo
3. `combo_selections` - Opciones seleccionables

### **Tablas Jobs (3)**
1. `jobs_applications` - Aplicaciones de trabajo
2. `jobs_keywords` - Keywords para análisis
3. `jobs_questions` - Preguntas de aplicación

### **Tablas JobsTracker (8)**
1. `tracker_candidates` - Candidatos
2. `tracker_interviews` - Entrevistas
3. `tracker_questions` - Preguntas de entrevista
4. `tracker_kanban_status` - Estados kanban
5. `tracker_keywords` - Keywords de análisis
6. `tracker_qr_locations` - Ubicaciones QR
7. `tracker_qr_views` - Vistas de QR
8. `tracker_notifications` - Notificaciones

### **Tablas Concurso (3)**
1. `concurso_participantes` - Participantes
2. `concurso_state` - Estado del torneo
3. `concurso_tracking` - Tracking de visitas

### **Tablas Food Trucks (2)**
1. `food_trucks` - Datos de food trucks
2. `truck_schedules` - Horarios

### **Tablas Caja (2)**
1. `cash_register` - Registro de caja
2. `cashiers` - Cajeros

### **Tablas Mermas (1)**
1. `mermas` - Registro de mermas

### **Tablas API Keys (1)**
1. `api_keys` - Claves de API

---

## 📊 RESUMEN ACTUALIZADO COMPLETO

### **Total de Páginas: 45+**
- Usuario: 8 páginas
- Admin: 22 páginas
- Jobs: 5 páginas
- JobsTracker: 7 páginas
- Concurso: 5 páginas (deshabilitado)

### **Total de APIs: 80+**
- Productos: 5 APIs
- Ventas: 6 APIs
- TUU Pagos: 8 APIs
- Analytics: 8 APIs
- Compras: 6 APIs
- Calidad: 3 APIs
- Dashboard: 5 APIs
- Jobs: 8 APIs
- JobsTracker: 15 APIs
- Concurso: 6 APIs
- Food Trucks: 5 APIs
- Auth: 8 APIs
- Otros: 7+ APIs

### **Total de Tablas: 35+**
- Core: 11 tablas
- Analytics: 3 tablas
- Auxiliares: 3 tablas
- Combos: 3 tablas
- Jobs: 3 tablas
- JobsTracker: 8 tablas
- Concurso: 3 tablas
- Food Trucks: 2 tablas
- Caja: 2 tablas
- Mermas: 1 tabla
- API Keys: 1 tabla

---

## ⚠️ MÓDULOS CRÍTICOS PARA MIGRACIÓN

### **Prioridad 1 (CRÍTICO):**
1. ✅ index.astro + MenuApp
2. ✅ checkout.astro + TUU
3. ✅ admin/index.astro + Dashboard
4. ✅ get_productos.php
5. ✅ create_payment_working.php
6. ✅ registrar_venta.php
7. ✅ Tablas: productos, ingredientes, ventas, tuu_orders

### **Prioridad 2 (IMPORTANTE):**
1. compras.astro + APIs compras
2. ventas-detalle.astro
3. admin/ingredients.astro
4. comandas/index.astro
5. Tablas: compras, compras_items, proveedores

### **Prioridad 3 (SECUNDARIO):**
1. admin/calidad.astro
2. mermas.astro
3. arqueo.astro
4. admin/combos.astro
5. Tablas: quality_*, mermas, combos

### **Prioridad 4 (OPCIONAL):**
1. Jobs module (completo)
2. JobsTracker module (completo)
3. Concurso module (deshabilitado)
4. Food Trucks
5. Analytics avanzado

---

**Última actualización:** 2025-01-XX  
**Proyecto:** La Ruta 11 - Sistema Completo  
**Stack:** Astro + React + PHP + MySQL
