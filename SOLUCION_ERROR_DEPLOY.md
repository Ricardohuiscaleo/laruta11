# 🚀 Solución Error: Cannot find module '/app/dist/server/entry.mjs'

## ❌ Problema
El error ocurre porque Easypanel intenta ejecutar las apps como servidores Node.js, pero son aplicaciones **estáticas** de Astro.

## ✅ Solución

### 1. Archivos Creados
Se crearon 3 archivos `nixpacks.toml` (uno por carpeta):
- `/app/nixpacks.toml`
- `/caja/nixpacks.toml`
- `/landing/nixpacks.toml`

### 2. Configuración en Easypanel

Para cada servicio en Easypanel:

#### **Opción A: Usar nixpacks.toml (Recomendado)**
1. Ve a cada servicio en Easypanel
2. En "Build Settings" → "Build Path" configura:
   - **app**: `/app`
   - **caja**: `/caja`
   - **landing**: `/landing`
3. Easypanel detectará automáticamente el `nixpacks.toml`
4. Redeploy cada servicio

#### **Opción B: Configuración Manual**
Si nixpacks.toml no funciona, configura manualmente:

**Build Command:**
```bash
npm ci && npm run build
```

**Start Command:**
```bash
npx serve dist -l 3000
```

**Port:** `3000`

### 3. Variables de Entorno
Copia las variables de `SECRETS.txt` a cada servicio en Easypanel:

**Para /app:**
```env
PUBLIC_SUPABASE_URL=https://uznvakpuuxnpdhoejrog.supabase.co
PUBLIC_SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
# ... resto de variables necesarias
```

**Para /caja y /landing:**
Agregar las variables que necesite cada app.

### 4. Verificación
Después del deploy, verifica:
- ✅ `laruta11.cl` → landing funcionando
- ✅ `app.laruta11.cl` → app funcionando
- ✅ `caja.laruta11.cl` → caja funcionando

## 📝 Notas
- Las 3 apps son **estáticas** (no SSR)
- Se sirven con `serve` en puerto 3000
- Los archivos PHP en `/app/api/` deben configurarse por separado si necesitas backend PHP
