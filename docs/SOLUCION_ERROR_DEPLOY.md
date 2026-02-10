# 🚀 Guía de Deploy - La Ruta 11 Monorepo

## 📦 Estructura del Proyecto

```
laruta11/
├── app/              # App clientes (Astro estático)
├── caja/             # Sistema caja (Astro estático)
├── landing/          # Landing page (Astro estático)
│   └── api-go/       # API Go para S3
```

---

## 🎯 Apps Astro Estáticas (/app, /caja, /landing)

### ❌ Problema Original
Easypanel intentaba ejecutar las apps como servidores Node.js SSR, causando:
```
Cannot find module '/app/dist/server/entry.mjs'
```

### ✅ Solución

**Archivos creados:**
- `/app/nixpacks.toml`
- `/caja/nixpacks.toml`
- `/landing/nixpacks.toml`

**Configuración en Easypanel:**

1. **Build Method:** Nixpacks
2. **Build Path:** `/app` (o `/caja`, `/landing`)
3. **Port:** 3000
4. Easypanel detecta automáticamente el `nixpacks.toml`

**Contenido de nixpacks.toml:**
```toml
[phases.install]
cmds = ["npm ci"]

[phases.build]
cmds = ["npm run build"]

[start]
cmd = "npx serve dist -l 3000"
```

**Variables de entorno:**
```env
PUBLIC_SUPABASE_URL=https://uznvakpuuxnpdhoejrog.supabase.co
PUBLIC_SUPABASE_ANON_KEY=eyJ...
```

---

## 🔥 API Go (/landing/api-go) - CRÍTICO

### ⚠️ IMPORTANTE: NO USAR NIXPACKS PARA GO

**Problema con Nixpacks + Go:**
- ❌ No maneja bien `go.sum` en monorepos
- ❌ Errores de checksum imposibles de resolver
- ❌ Caché corrupto sin forma de limpiar
- ❌ Proceso de build opaco y no debuggeable

### ✅ Solución: Dockerfile Clásico

**Archivo:** `/landing/api-go/Dockerfile`

```dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download && go mod tidy

COPY . .
RUN go build -o api main.go

EXPOSE 3001

CMD ["./api"]
```

### 📋 Pasos para Deploy de API Go

#### 1. Preparación Local (REQUERIDO)

**Instalar Go en macOS:**
```bash
brew install go
```

**Generar go.sum:**
```bash
cd landing/api-go
go mod tidy
```

**Verificar que go.sum existe:**
```bash
ls -la go.sum
```

#### 2. Configuración en Easypanel

1. **Build Method:** Dockerfile
2. **Build Path:** `landing/api-go`
3. **Dockerfile:** `Dockerfile`
4. **Port:** 3001

#### 3. Variables de Entorno

```env
AWS_ACCESS_KEY_ID=tu_access_key
AWS_SECRET_ACCESS_KEY=tu_secret_key
S3_REGION=us-east-1
S3_BUCKET=laruta11-images
PORT=3001
```

#### 4. Endpoints Disponibles

- `POST /api/s3` - Operaciones S3 (list, upload, delete, test)
- `GET /api/health` - Health check

**Ejemplo de uso:**
```bash
# Health check
curl https://api.laruta11.cl/api/health

# Test S3 connection
curl -X POST https://api.laruta11.cl/api/s3 -d "action=test"
```

---

## 🚨 Errores Comunes y Soluciones

### Error: "missing go.sum entry"

**Causa:** No existe `go.sum` o está corrupto

**Solución:**
```bash
cd landing/api-go
rm go.sum  # Si existe
go mod tidy
git add go.sum
git commit -m "regenerate go.sum"
git push
```

### Error: "checksum mismatch"

**Causa:** `go.sum` desactualizado o caché de Easypanel

**Solución:**
1. Regenerar `go.sum` localmente (ver arriba)
2. En Easypanel: Settings → Destroy Service
3. Crear servicio nuevo desde cero

### Error: "imported and not used"

**Causa:** Import no utilizado en el código Go

**Solución:**
```bash
# Eliminar el import del archivo .go
# Ejemplo: quitar "net/http" si no se usa
git add .
git commit -m "fix: remove unused import"
git push
```

---

## 📈 Escalabilidad: De 5 a 200+ APIs

### Estructura Recomendada

```
laruta11/
├── apps/
│   ├── app/
│   ├── caja/
│   └── landing/
├── apis/
│   ├── api-s3/          # API actual
│   ├── api-payments/    # Nueva API
│   ├── api-orders/      # Nueva API
│   └── api-users/       # Nueva API
└── shared/
    ├── types/
    └── utils/
```

### Template para Nuevas APIs Go

**1. Crear carpeta:**
```bash
mkdir -p apis/api-nombre
cd apis/api-nombre
```

**2. Inicializar Go:**
```bash
go mod init laruta11-api-nombre
```

**3. Copiar Dockerfile base:**
```bash
cp ../api-s3/Dockerfile .
```

**4. Crear main.go:**
```go
package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    r.GET("/api/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    
    log.Printf("🚀 API running on port %s", port)
    r.Run(":" + port)
}
```

**5. Generar go.sum:**
```bash
go mod tidy
```

**6. Commit y push:**
```bash
git add .
git commit -m "add: api-nombre"
git push
```

**7. Deploy en Easypanel:**
- Build Method: Dockerfile
- Build Path: `apis/api-nombre`
- Port: 3000 (o el que uses)

### Checklist para Cada Nueva API

- [ ] Carpeta creada en `/apis/`
- [ ] `go.mod` inicializado
- [ ] `Dockerfile` copiado y ajustado
- [ ] `main.go` con health check
- [ ] `go mod tidy` ejecutado
- [ ] `go.sum` commiteado
- [ ] Variables de entorno documentadas
- [ ] Servicio creado en Easypanel
- [ ] Health check funcionando
- [ ] Dominio configurado

---

## 🎯 Verificación Final

### Apps Astro
- ✅ `laruta11.cl` → landing
- ✅ `app.laruta11.cl` → app
- ✅ `caja.laruta11.cl` → caja

### APIs Go
- ✅ `api.laruta11.cl/api/health` → API S3

---

## 💡 Lecciones Aprendidas

1. **Dockerfile > Nixpacks para Go** - Siempre usar Dockerfile para APIs Go
2. **go.sum es obligatorio** - Generar localmente con `go mod tidy`
3. **Caché de Easypanel** - Si falla, destruir y recrear servicio
4. **Monorepo** - Usar Build Path específico para cada servicio
5. **Health checks** - Siempre incluir endpoint `/api/health`
