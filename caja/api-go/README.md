# API Go - Caja Laruta11

API optimizada en Go para sistema de caja, inventario y compras.

## 📁 Estructura

```
api-go/
├── main.go          # Server + rutas (80 líneas)
├── auth.go          # Login (65 líneas)
├── compras.go       # Compras (150 líneas)
├── resources.go     # Ingredientes + Categories + Checklist (120 líneas)
├── handlers.go      # Products + Orders (75 líneas)
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
```

**Total: 490 líneas** (vs 800+ PHP)

## 🚀 Endpoints (25)

### Auth (3)
- `POST /api/auth/login` - Caja/Inventario/Comandas/Admin
- `GET /api/auth/check`
- `POST /api/auth/logout`

### Compras (9)
- `GET /api/compras`
- `GET /api/compras/items`
- `GET /api/compras/proveedores`
- `GET /api/compras/saldo`
- `GET /api/compras/historial-saldo`
- `GET /api/compras/precio-historico?ingrediente_id=X`
- `POST /api/compras`
- `DELETE /api/compras/:id`
- `POST /api/compras/:id/respaldo`

### Ingredientes (3)
- `GET /api/ingredientes`
- `POST /api/ingredientes`
- `DELETE /api/ingredientes/:id`

### Categories (3)
- `GET /api/categories`
- `POST /api/categories`
- `DELETE /api/categories/:id`

### Checklist (3)
- `GET /api/checklist?date=2026-02-10`
- `POST /api/checklist`
- `DELETE /api/checklist/:id`

### Products (2)
- `GET /api/products?include_inactive=1`
- `GET /api/products/:id`

### Orders (2)
- `GET /api/orders/pending`
- `POST /api/orders/status`

## 🔐 Auth

```json
POST /api/auth/login
{
  "user": "ruta11caja",
  "pass": "***",
  "type": "caja"
}
```

**Tipos**: `caja`, `inventario`, `comandas`, `admin`

## ⚙️ Env Vars

```bash
APP_DB_HOST=websites_mysql-laruta11
APP_DB_NAME=laruta11
APP_DB_USER=laruta11_user
APP_DB_PASS=***
PORT=3002
CAJA_USER_CAJERA=***
INVENTARIO_USER=***
INVENTARIO_PASSWORD=***
ADMIN_USER_ADMIN=***
```

## 🏗️ Deploy

```bash
go mod tidy
go build
```

Easypanel: Dockerfile, Port 3002
