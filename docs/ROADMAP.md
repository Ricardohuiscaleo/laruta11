# 📊 Estado del Proyecto La Ruta 11 - Roadmap

## 🎯 Contexto General

**Proyecto:** Sistema completo de restaurante con 3 apps frontend + APIs backend
**Stack:** Astro (frontend) + Go (APIs) + Supabase (DB)
**Deploy:** Easypanel (VPS)
**Monorepo:** GitHub

---

## ✅ Completado (Fase 1)

### Frontend Desplegado
- ✅ `/app` - App clientes (Astro estático) → `app.laruta11.cl`
- ✅ `/caja` - Sistema caja (Astro estático) → `caja.laruta11.cl`
- ✅ `/landing` - Landing page (Astro estático) → `laruta11.cl`

**Método:** Nixpacks + `nixpacks.toml`

### API Desplegada
- ✅ `/landing/api-go` - API S3 para imágenes (Go) → `api.laruta11.cl`

**Método:** Dockerfile
**Endpoints:**
- `POST /api/s3` (list, upload, delete, test)
- `GET /api/health`

---

## 🚧 En Progreso (Fase 2)

### APIs PHP a Migrar a Go

#### 1. `/app/api/` - APIs del App de Clientes

**Archivos PHP actuales:**
```
app/api/
├── get-categories.php
├── get-products.php
├── create-order.php
├── get-orders.php
└── update-order-status.php
```

**Funcionalidad:**
- Obtener categorías de productos
- Obtener productos por categoría
- Crear pedidos de clientes
- Consultar pedidos
- Actualizar estado de pedidos

**Migración a Go:**
- [ ] Crear `/apis/api-app/`
- [ ] Endpoints REST con Gin
- [ ] Conexión a Supabase
- [ ] CORS configurado
- [ ] Validación de datos
- [ ] Manejo de errores

#### 2. `/caja/api/` - APIs del Sistema de Caja

**Archivos PHP actuales:**
```
caja/api/
├── get-pending-orders.php
├── update-order-status.php
├── get-sales-report.php
├── get-inventory.php
└── update-inventory.php
```

**Funcionalidad:**
- Ver pedidos pendientes
- Actualizar estado de pedidos
- Reportes de ventas
- Gestión de inventario
- Actualizar stock

**Migración a Go:**
- [ ] Crear `/apis/api-caja/`
- [ ] Endpoints REST con Gin
- [ ] Conexión a Supabase
- [ ] Autenticación/autorización
- [ ] Reportes en tiempo real
- [ ] WebSockets para notificaciones

---

## 📋 Plan de Migración PHP → Go

### Paso 1: Análisis de APIs PHP

**Para cada archivo PHP:**
1. Documentar endpoints (método, ruta, params)
2. Identificar queries a Supabase
3. Listar validaciones necesarias
4. Mapear respuestas JSON

### Paso 2: Crear Estructura Go

```
apis/
├── api-app/
│   ├── main.go
│   ├── handlers/
│   │   ├── categories.go
│   │   ├── products.go
│   │   └── orders.go
│   ├── models/
│   │   ├── category.go
│   │   ├── product.go
│   │   └── order.go
│   ├── db/
│   │   └── supabase.go
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
└── api-caja/
    ├── main.go
    ├── handlers/
    │   ├── orders.go
    │   ├── reports.go
    │   └── inventory.go
    ├── models/
    ├── db/
    ├── Dockerfile
    ├── go.mod
    └── go.sum
```

### Paso 3: Implementación por Prioridad

**Alta prioridad (crítico para operación):**
1. `get-products.php` → Mostrar menú
2. `create-order.php` → Crear pedidos
3. `get-pending-orders.php` → Ver pedidos en caja

**Media prioridad:**
4. `get-categories.php`
5. `update-order-status.php`
6. `get-orders.php`

**Baja prioridad:**
7. `get-sales-report.php`
8. `get-inventory.php`
9. `update-inventory.php`

### Paso 4: Testing y Deploy

**Por cada API:**
1. Desarrollo local
2. `go mod tidy`
3. Testing con Postman/curl
4. Commit y push
5. Deploy en Easypanel
6. Configurar variables de entorno
7. Probar en producción
8. Actualizar frontend para usar nueva API

---

## 🔧 Tecnologías y Dependencias

### Go Packages Necesarios

```go
// HTTP Framework
github.com/gin-gonic/gin

// Supabase Client
github.com/supabase-community/supabase-go

// CORS
github.com/gin-contrib/cors

// Validación
github.com/go-playground/validator/v10

// Variables de entorno
github.com/joho/godotenv

// UUID
github.com/google/uuid
```

### Variables de Entorno Comunes

```env
# Supabase
SUPABASE_URL=https://uznvakpuuxnpdhoejrog.supabase.co
SUPABASE_KEY=eyJ...
SUPABASE_SERVICE_KEY=eyJ...

# Server
PORT=3000
GIN_MODE=release

# CORS
ALLOWED_ORIGINS=https://app.laruta11.cl,https://caja.laruta11.cl
```

---

## 📝 Checklist de Migración

### API App (Clientes)

- [ ] Analizar `/app/api/*.php`
- [ ] Crear `/apis/api-app/`
- [ ] Implementar handlers
- [ ] Conectar Supabase
- [ ] Testing local
- [ ] Deploy en Easypanel
- [ ] Actualizar frontend `/app`
- [ ] Eliminar archivos PHP

### API Caja (Admin)

- [ ] Analizar `/caja/api/*.php`
- [ ] Crear `/apis/api-caja/`
- [ ] Implementar handlers
- [ ] Conectar Supabase
- [ ] Testing local
- [ ] Deploy en Easypanel
- [ ] Actualizar frontend `/caja`
- [ ] Eliminar archivos PHP

---

## 🎯 Objetivos Finales

### Arquitectura Target

```
Frontend (Astro Estático)
├── app.laruta11.cl → /app
├── caja.laruta11.cl → /caja
└── laruta11.cl → /landing

APIs (Go + Gin)
├── api-app.laruta11.cl → /apis/api-app
├── api-caja.laruta11.cl → /apis/api-caja
└── api-s3.laruta11.cl → /apis/api-s3

Database
└── Supabase (PostgreSQL)
```

### Beneficios de la Migración

1. **Performance:** Go es 10-50x más rápido que PHP
2. **Escalabilidad:** Mejor manejo de concurrencia
3. **Mantenibilidad:** Código tipado y estructurado
4. **Deploy:** Binario único, sin dependencias
5. **Costo:** Menor uso de recursos del servidor

---

## 📚 Documentos de Referencia

- `SOLUCION_ERROR_DEPLOY.md` - Guía completa de deploy
- `README.md` - Estructura del monorepo
- `SECRETS.txt` - Variables de entorno (NO COMMITEAR)

---

## 🚀 Próximos Pasos Inmediatos

1. **Analizar APIs PHP existentes**
   - Listar todos los endpoints
   - Documentar parámetros y respuestas
   - Identificar lógica de negocio

2. **Crear primera API Go (api-app)**
   - Empezar con `get-products.php`
   - Implementar en Go
   - Deploy y testing

3. **Migración gradual**
   - Una API a la vez
   - Mantener PHP funcionando en paralelo
   - Cambiar frontend cuando Go esté listo

---

## 💡 Información Crítica para Continuidad

### Si pierdes contexto, recuerda:

1. **Estructura:** Monorepo con apps Astro + APIs Go
2. **Deploy:** Easypanel con Dockerfile para Go, Nixpacks para Astro
3. **Problema resuelto:** go.sum se genera con `go mod tidy` localmente
4. **Nunca usar:** Nixpacks para Go (solo Dockerfile)
5. **Siguiente tarea:** Migrar APIs PHP a Go
6. **Prioridad:** Empezar con `/app/api/get-products.php`

### Comandos esenciales:

```bash
# Generar go.sum
cd apis/api-nombre
go mod tidy

# Deploy
git add .
git commit -m "mensaje"
git push

# En Easypanel: Dockerfile, Build Path: apis/api-nombre
```

---

**Última actualización:** 10 Feb 2026
**Estado:** Fase 1 completa, iniciando Fase 2 (migración PHP → Go)
