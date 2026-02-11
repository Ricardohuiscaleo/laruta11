# Reorganización en Carpetas

## 🎯 Objetivo

Pasar de estructura plana a carpetas lógicas:

```
ANTES:                    DESPUÉS:
api-go/                   api-go/
├── main.go              ├── main.go
├── auth.go              ├── auth/
├── compras.go           │   └── handler.go
├── resources.go         ├── compras/
├── handlers.go          │   └── handler.go
└── ...                  ├── inventory/
                         │   └── handler.go
                         ├── quality/
                         │   └── handler.go
                         ├── catalog/
                         │   └── handler.go
                         ├── orders/
                         │   └── handler.go
                         └── shared/
                             └── db.go
```

## 🚀 Pasos Manuales

### 1. Crear carpetas
```bash
cd caja/api-go
mkdir -p auth compras inventory quality catalog orders shared
```

### 2. Mover auth.go → auth/handler.go
```bash
# Cambiar package main → package auth
# Cambiar func (s *Server) → func (h *Handler)
# Agregar type Handler struct{ DB *sql.DB }
mv auth.go auth/handler.go
```

### 3. Mover compras.go → compras/handler.go
```bash
# Mismo proceso
mv compras.go compras/handler.go
```

### 4. Mover resources.go → inventory/handler.go + quality/handler.go
```bash
# Separar ingredientes/categories → inventory/
# Separar checklist → quality/
```

### 5. Mover handlers.go → catalog/handler.go + orders/handler.go
```bash
# Separar products → catalog/
# Separar orders → orders/
```

### 6. Actualizar main.go
```go
import (
    "api-go/auth"
    "api-go/compras"
    "api-go/inventory"
    "api-go/quality"
    "api-go/catalog"
    "api-go/orders"
)

authH := &auth.Handler{DB: db}
comprasH := &compras.Handler{DB: db}
// ...

r.POST("/api/auth/login", authH.Login)
r.GET("/api/compras", comprasH.GetCompras)
```

### 7. Actualizar go.mod
```bash
go mod tidy
```

### 8. Compilar
```bash
go build
```

## ⚡ Opción Rápida (Script)

```bash
chmod +x reorganize.sh
./reorganize.sh
go mod tidy
go build
```

## 📊 Resultado

- **7 carpetas** organizadas por dominio
- **Código más mantenible**
- **Imports claros**
- **Fácil testing por módulo**

## ⚠️ Importante

Por ahora, **mantener estructura plana** es más simple para 25 endpoints.

Reorganizar cuando tengamos **50+ endpoints**.
