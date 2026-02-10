# 🚀 MIGRACIÓN A VPS CON EASYPANEL

## 📋 ESTRUCTURA DE DOMINIOS

Tienes **3 subdominios** que necesitan **3 servicios separados** en Easypanel:

### 1. **laruta11.cl** (Landing/Web principal)
- **Tipo**: Sitio web estático o WordPress
- **Repo**: `laruta11-web` (nuevo)
- **Puerto**: 80/443
- **Contenido**: Página principal del negocio

### 2. **app.laruta11.cl** (App de clientes)
- **Tipo**: Astro + React (este proyecto actual)
- **Repo**: `ruta11caja` (este mismo, renombrado)
- **Puerto**: 4321 → 80/443
- **Contenido**: Menú, pedidos, checkout

### 3. **caja.laruta11.cl** (Sistema interno)
- **Tipo**: Astro + React (mismo código que app)
- **Repo**: `ruta11caja` (mismo repo, diferente servicio)
- **Puerto**: 4321 → 80/443
- **Contenido**: Admin, comandas, inventario, caja

---

## 🎯 ESTRATEGIA RECOMENDADA

### Opción A: **1 Repo, 2 Servicios** (RECOMENDADO)
```
GitHub:
└── ruta11-app (este proyecto)
    ├── src/pages/
    │   ├── index.astro          → app.laruta11.cl
    │   ├── admin/               → caja.laruta11.cl
    │   └── comandas/            → caja.laruta11.cl
    └── api/                     → compartido

Easypanel:
├── Service 1: ruta11-app
│   └── Domain: app.laruta11.cl
└── Service 2: ruta11-caja
    └── Domain: caja.laruta11.cl
```

**Ventajas:**
- ✅ Mismo código, misma API
- ✅ Un solo repo a mantener
- ✅ Actualizaciones simultáneas
- ✅ Comparten base de datos

**Configuración:**
```javascript
// astro.config.mjs
export default defineConfig({
  output: 'server',
  adapter: node({ mode: 'standalone' }),
  server: { port: 4321 }
});
```

---

### Opción B: **2 Repos Separados**
```
GitHub:
├── ruta11-app (clientes)
│   └── Solo páginas públicas
└── ruta11-admin (interno)
    └── Solo páginas admin

Easypanel:
├── Service 1: ruta11-app → app.laruta11.cl
└── Service 2: ruta11-admin → caja.laruta11.cl
```

**Ventajas:**
- ✅ Separación total
- ✅ Seguridad mejorada
- ❌ Duplicación de código
- ❌ Dos repos a mantener

---

## 📦 CONFIGURACIÓN EASYPANEL

### 1. Crear Proyecto
```
Easypanel Dashboard
└── New Project: "ruta11"
```

### 2. Servicio 1: App Clientes
```yaml
Name: ruta11-app
Type: App
Source: GitHub
Repo: tu-usuario/ruta11-app
Branch: main
Build:
  Command: npm install && npm run build
  Output: dist/
Start:
  Command: node dist/server/entry.mjs
  Port: 4321
Domain: app.laruta11.cl
Environment:
  - NODE_ENV=production
  - PUBLIC_SUPABASE_URL=...
  - PUBLIC_SUPABASE_ANON_KEY=...
```

### 3. Servicio 2: Sistema Caja
```yaml
Name: ruta11-caja
Type: App
Source: GitHub
Repo: tu-usuario/ruta11-app (mismo)
Branch: main
Build:
  Command: npm install && npm run build
  Output: dist/
Start:
  Command: node dist/server/entry.mjs
  Port: 4321
Domain: caja.laruta11.cl
Environment:
  - NODE_ENV=production
  - ADMIN_MODE=true
```

### 4. Base de Datos MySQL
```yaml
Name: ruta11-mysql
Type: MySQL 8.0
Database: u958525313_app
User: u958525313_app
Password: wEzho0-hujzoz-cevzin
Port: 3306
Volume: /var/lib/mysql
```

---

## 🔧 PASOS DE MIGRACIÓN

### 1. Preparar Repositorio
```bash
cd /Users/ricardohuiscaleollafquen/ruta11caja

# Inicializar Git (si no existe)
git init
git add .
git commit -m "Proyecto limpio - 399 archivos obsoletos eliminados"

# Crear repo en GitHub
# Ir a github.com → New Repository → "ruta11-app"

# Conectar y subir
git remote add origin https://github.com/TU-USUARIO/ruta11-app.git
git branch -M main
git push -u origin main
```

### 2. Configurar Variables de Entorno
```bash
# Crear .env.production
cp .env .env.production

# Editar con valores de producción
nano .env.production
```

### 3. Crear Dockerfile (Opcional)
```dockerfile
FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm run build
EXPOSE 4321
CMD ["node", "dist/server/entry.mjs"]
```

### 4. Configurar Easypanel
1. Conectar GitHub a Easypanel
2. Crear proyecto "ruta11"
3. Agregar servicio MySQL
4. Agregar servicio "ruta11-app"
5. Agregar servicio "ruta11-caja"
6. Configurar dominios
7. Deploy

---

## 🌐 CONFIGURACIÓN DNS

En tu proveedor de DNS (Cloudflare/Hostinger):

```
Tipo  | Nombre | Valor              | TTL
------|--------|-------------------|-----
A     | @      | IP_VPS            | Auto
A     | app    | IP_VPS            | Auto
A     | caja   | IP_VPS            | Auto
```

---

## 📁 ESTRUCTURA FINAL

```
VPS con Easypanel:
├── MySQL Container
│   └── u958525313_app (base de datos)
├── ruta11-app Container
│   ├── Domain: app.laruta11.cl
│   ├── Port: 4321 → 443
│   └── SSL: Auto (Let's Encrypt)
└── ruta11-caja Container
    ├── Domain: caja.laruta11.cl
    ├── Port: 4321 → 443
    └── SSL: Auto (Let's Encrypt)
```

---

## ✅ CHECKLIST PRE-MIGRACIÓN

- [ ] Backup completo de base de datos
- [ ] Exportar variables de entorno
- [ ] Crear repositorio GitHub
- [ ] Subir código limpio
- [ ] Configurar .env.production
- [ ] Probar build local: `npm run build`
- [ ] Documentar credenciales
- [ ] Configurar DNS
- [ ] Crear cuenta Easypanel

---

## 🚨 IMPORTANTE

**Base de Datos:**
- Exportar desde Hostinger: `mysqldump -u u958525313_app -p u958525313_app > backup.sql`
- Importar a VPS: `mysql -u root -p u958525313_app < backup.sql`

**Archivos PHP:**
- La carpeta `api/` debe estar accesible
- Configurar PHP 8.1+ en Easypanel
- Verificar extensiones: mysqli, pdo, curl

**Dominios:**
- Esperar propagación DNS (24-48h)
- Usar Cloudflare para CDN/SSL
- Configurar redirects HTTP → HTTPS

---

## 📞 SOPORTE

**Easypanel Docs:** https://easypanel.io/docs
**Astro Deploy:** https://docs.astro.build/en/guides/deploy/

¿Necesitas ayuda con algún paso específico?
