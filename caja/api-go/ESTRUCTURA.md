# Estructura de Carpetas - API Go

## 📁 Organización Lógica

```
api-go/
├── main.go                    # Entry point + rutas
│
├── auth/                      # 🔐 Autenticación
│   └── auth.go               # Login caja/inventario/comandas/admin
│
├── compras/                   # 🛒 Módulo Compras
│   ├── compras.go            # CRUD compras
│   ├── items.go              # Items disponibles
│   ├── proveedores.go        # Proveedores
│   └── saldo.go              # Saldo + historial
│
├── inventory/                 # 📦 Inventario
│   ├── ingredientes.go       # CRUD ingredientes
│   └── categories.go         # CRUD categorías
│
├── quality/                   # ✅ Calidad
│   └── checklist.go          # Checklists diarios
│
├── catalog/                   # 🍔 Catálogo
│   ├── products.go           # CRUD productos
│   └── combos.go             # CRUD combos
│
├── orders/                    # 📋 Órdenes
│   ├── orders.go             # CRUD órdenes
│   └── comandas.go           # Comandas cocina
│
├── analytics/                 # 📊 Analytics
│   ├── dashboard.go          # Dashboard cards
│   └── reports.go            # Reportes financieros
│
├── shared/                    # 🔧 Compartido
│   ├── db.go                 # Database connection
│   ├── middleware.go         # CORS, auth middleware
│   └── utils.go              # Helpers
│
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
```

## 🗂️ Agrupación por Dominio

### 1. **auth/** - Autenticación
- Login simple con env vars
- 4 tipos: caja, inventario, comandas, admin
- Sin JWT (por ahora)

### 2. **compras/** - Gestión de Compras
- Historial de compras
- Items disponibles (ingredientes + productos)
- Proveedores
- Saldo disponible (con lógica turnos)
- Historial de saldo
- Precio histórico
- Registrar compra
- Eliminar compra
- Upload respaldo (S3)

### 3. **inventory/** - Inventario
- CRUD ingredientes
- CRUD categorías
- Stock management

### 4. **quality/** - Control de Calidad
- Checklists diarios (apertura, cierre, limpieza)
- Templates de checklist

### 5. **catalog/** - Catálogo de Productos
- CRUD productos
- CRUD combos
- Recetas
- Precios

### 6. **orders/** - Gestión de Órdenes
- CRUD órdenes
- Comandas cocina
- Estados de orden
- Pagos TUU

### 7. **analytics/** - Analytics & Reportes
- Dashboard cards
- Reportes financieros
- Ventas por período
- Tracking de usuarios

### 8. **shared/** - Código Compartido
- Database connection pool
- Middleware (CORS, auth)
- Helpers (formatters, validators)

## 📊 Mapeo PHP → Go

| PHP | Go Module | Endpoints |
|-----|-----------|-----------|
| `/api/auth/*` | `auth/` | 3 |
| `/api/compras/*` | `compras/` | 9 |
| `/api/get_ingredientes.php` | `inventory/` | 3 |
| `/api/categories.php` | `inventory/` | 3 |
| `/api/checklist.php` | `quality/` | 3 |
| `/api/get_menu_products.php` | `catalog/` | 5 |
| `/api/create_order.php` | `orders/` | 5 |
| `/api/get_dashboard_*.php` | `analytics/` | 4 |

## 🎯 Ventajas

1. **Claridad**: Cada carpeta = 1 dominio de negocio
2. **Escalabilidad**: Fácil agregar nuevos módulos
3. **Mantenibilidad**: Código relacionado junto
4. **Testing**: Tests por módulo
5. **Documentación**: README por módulo

## 🚀 Implementación

Actualmente todo está en archivos planos:
- `auth.go`
- `compras.go`
- `resources.go`
- `handlers.go`

**Próximo paso**: Mover a carpetas cuando tengamos más endpoints.

Por ahora, **5 archivos son suficientes** para 25 endpoints.
