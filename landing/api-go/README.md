# 🐹 La Ruta 11 - Go API

1 API que reemplaza 6 APIs PHP.

## 🚀 Endpoints

**POST /api/s3**
- `action=list` - Listar imágenes
- `action=upload` - Subir imagen
- `action=delete` - Eliminar imagen
- `action=test` - Test conexión S3

**GET /api/health** - Health check

## 📦 Deploy en Easypanel

1. Crear servicio: **"api-laruta11"**
2. Repository: Tu repo GitHub
3. Branch: `main`
4. **Build Path**: `/landing/api-go`
5. **Port**: `3001`
6. **Domain**: `api.laruta11.cl`

### Variables de entorno:
```env
AWS_ACCESS_KEY_ID=<tu_access_key>
AWS_SECRET_ACCESS_KEY=<tu_secret_key>
S3_REGION=us-east-1
S3_BUCKET=laruta11-images
PORT=3001
```

**Nota:** Obtén las credenciales reales de `SECRETS.txt`

## 🔄 Actualizar frontend

En `landing/src/pages/admin.astro` cambiar:

```javascript
// Antes
const API_URL = '/api/s3-manager.php';

// Después  
const API_URL = 'https://api.laruta11.cl/api/s3';
```

## ✅ Verificar

```bash
# Local
curl http://localhost:3001/api/health

# Producción
curl https://api.laruta11.cl/api/health
```

## 📊 Ventajas vs PHP

- ⚡ 50x más rápido
- 📦 1 binario (no dependencias)
- 🔒 Más seguro
- 💰 Menos recursos
- 🚀 Async nativo
