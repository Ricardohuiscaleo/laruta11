# 🎓 Lecciones Aprendidas - La Ruta 11

## 🔴 ERRORES CRÍTICOS Y SOLUCIONES

### 1. API Response Structure Mismatch
**Problema**: Frontend esperaba `data.images` pero Go API devolvía `data.data.images`

**Solución**: 
```javascript
// Soportar ambas estructuras
const images = data.data?.images || data.images || [];
```

**Lección**: Siempre verificar estructura de respuesta entre frontend y backend antes de deploy.

---

### 2. Go S3 Upload - Binary Data Corruption
**Problema**: `NotImplemented: A header you provided implies functionality that is not implemented`

**Causa Raíz**:
- ❌ Leer archivo dos veces (`io.ReadAll` + reusar `file`)
- ❌ Usar `strings.NewReader(string(body))` corrompe imágenes binarias
- ❌ Falta prefijo `menu/` en key

**Solución Correcta**:
```go
import "bytes"  // Agregar import

body, err := io.ReadAll(file)  // Leer UNA vez
// ...
Body: bytes.NewReader(body),   // bytes.NewReader NO strings.NewReader
Key: aws.String("menu/" + filename),  // Agregar prefijo
```

**Lección**: 
- `strings.NewReader(string([]byte))` convierte bytes a string y corrompe datos binarios
- Siempre usar `bytes.NewReader()` para archivos binarios (imágenes, PDFs, etc.)
- En Go, los `io.Reader` se consumen. No se pueden leer dos veces sin reset.

---

### 3. Easypanel Auto-Deploy No Funciona
**Problema**: Push a GitHub no actualiza servicios automáticamente

**Solución**: 
1. Push a GitHub
2. Ir a Easypanel → Servicio específico → **Rebuild manual**
3. Esperar ~1 segundo

**Lección**: Easypanel requiere rebuild manual por servicio. No confiar en auto-deploy.

---

### 4. Browser Cache en Admin
**Problema**: Cambios en frontend no se reflejan después de deploy

**Solución**: Hard refresh `Cmd+Shift+R` (Mac) o `Ctrl+Shift+R` (Windows)

**Lección**: Siempre hacer hard refresh después de deploy de frontend.

---

### 5. S3 Bucket Structure
**Estructura Correcta**:
```
s3://laruta11-images/
├── menu/          ← 36 imágenes (filtradas por API)
├── checklist/     ← 74 imágenes
└── compras/       ← Otras imágenes
```

**Lección**: 
- API Go filtra por `prefix: "menu/"` en ListObjects
- Upload debe agregar `menu/` al key
- PHP legacy no agregaba prefijo (error histórico)

---

### 6. Go Upload - Missing File Extension
**Problema**: Archivo sube a S3 pero no aparece en frontend. Key: `menu/barcodecc` sin extensión

**Causa**: `custom_name` del frontend no preserva extensión original del archivo

**Solución**:
```go
if filename := c.PostForm("custom_name"); filename != "" {
    // Extract extension from original file
    if idx := strings.LastIndex(header.Filename, "."); idx >= 0 {
        ext := header.Filename[idx:]
        if !strings.HasSuffix(strings.ToLower(filename), strings.ToLower(ext)) {
            filename += ext  // Append .jpeg, .png, etc.
        }
    }
}
```

**Lección**: Siempre preservar extensión de archivo original. `isImageFile()` filtra por extensión.

---

### 7. Go Rename - CopySource URL Encoding
**Problema**: Rename falla con archivos que tienen espacios o caracteres especiales

**Causa**: `CopySource` en S3 debe estar URL encoded

**Solución**:
```go
import "net/url"

copySource := url.PathEscape(bucket + "/" + oldKey)
CopySource: aws.String(copySource)
```

**Lección**: S3 CopySource requiere URL encoding. Usar `url.PathEscape()` no `url.QueryEscape()`.

---

## ✅ CHECKLIST DE DEPLOY

### Antes de Deploy
- [ ] Verificar estructura de respuesta API coincide con frontend
- [ ] Probar endpoints en local con Postman/curl
- [ ] Verificar prefijos de carpetas S3 (`menu/`, etc.)
- [ ] Commit con mensaje descriptivo

### Durante Deploy
- [ ] Push a GitHub
- [ ] Rebuild manual del servicio específico en Easypanel
- [ ] Verificar logs del servicio (no confiar en "success")

### Después de Deploy
- [ ] Hard refresh en navegador (Cmd+Shift+R)
- [ ] Probar funcionalidad completa (list, upload, delete)
- [ ] Verificar en consola del navegador (F12) por errores

---

## 🏗️ ARQUITECTURA

### Servicios en Easypanel
1. **landing-r11** → `/landing` (Astro) → laruta11.cl
2. **app-r11** → `/app` (React/Vite) → app.laruta11.cl
3. **caja-r11** → `/caja` (PHP) → caja.laruta11.cl
4. **api-go-landing-r11** → `/landing/api-go` (Go) → API S3

### S3 Bucket Compartido
- Bucket: `laruta11-images`
- Región: `us-east-1`
- Todos los servicios acceden al mismo bucket
- ~~PHP usa AWS Signature V2 (manual)~~ **DEPRECADO**
- Go usa aws-sdk-go-v2 (oficial) ✅

### Legacy Code
- `/landing/api/` → **NO SUBIR A GIT** (agregado a `.gitignore`)
- Solo referencia histórica local
- Usar `/landing/api-go/` en producción

---

## 🐛 DEBUGGING TIPS

### Error 500 en Upload
1. Verificar logs del servicio api-go en Easypanel
2. Buscar `NotImplemented` → problema con headers S3
3. Buscar `EOF` → archivo leído dos veces
4. Verificar ContentType está definido

### Imágenes No Aparecen
1. Verificar prefijo `menu/` en S3
2. Console: `fetch(API_URL, {method:'POST', body:'action=list'}).then(r=>r.json())`
3. Verificar estructura: `data.data.images` vs `data.images`
4. Hard refresh navegador

### Upload Exitoso pero No Visible
1. Verificar key tiene prefijo `menu/`
2. Verificar extensión es válida (jpg, jpeg, png, gif, webp)
3. Refrescar galería con botón 🔄

---

## 📝 COMANDOS ÚTILES

```bash
# Commit y push
git add -A && git commit -m "fix: descripción" && git push

# Ver logs Go API (en Easypanel terminal)
docker logs -f <container-id>

# Test API desde consola navegador
fetch('https://websites-api-go-landing-r11.dj3bvg.easypanel.host/api/s3', {
  method: 'POST',
  headers: {'Content-Type': 'application/x-www-form-urlencoded'},
  body: 'action=list'
}).then(r => r.json()).then(console.log)
```

---

## 🔐 SEGURIDAD

### ✅ Hecho
- Google Maps API Key removida de `food-trucks.astro` (revocada)
- AWS credentials en variables de entorno (no en código)

### ⚠️ Pendiente
- Implementar rate limiting en API Go
- Agregar autenticación en endpoints admin
- CORS más restrictivo (actualmente `*`)

---

**Última actualización**: 2024 - Migración PHP → Go API
