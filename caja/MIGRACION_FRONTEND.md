# 🚀 Guía de Migración PHP → Go con Feature Flags

## 📋 Estado Actual

✅ **25 endpoints Go creados** (52 PHP reemplazados)
❌ **Frontend usando PHP** (pendiente migración)

## 🎯 Plan de Migración

### Fase 1: Activar Auth (Día 1)

```javascript
// src/utils/api-config.js
USE_GO_AUTH: true  // ✅ Activar
```

**Componentes a actualizar:**
```javascript
// Antes
fetch('/api/admin_auth.php', ...)

// Después
import { authApi } from '../utils/api-config';
fetch(authApi('/api/auth/login'), ...)
```

**Archivos afectados:**
- `components/AdminSPA.jsx` (3 usos)
- `pages/login.astro` (1 uso)
- `pages/comandas/index.astro` (1 uso)
- `components/MenuApp.jsx` (2 usos)

### Fase 2: Activar Compras (Día 2-3)

```javascript
USE_GO_COMPRAS: true
```

**Archivo principal:**
- `components/ComprasApp.jsx` (9 endpoints)

### Fase 3: Activar Inventory (Día 4-5)

```javascript
USE_GO_INVENTORY: true
```

**Archivos:**
- `pages/admin/inventario.astro`
- `pages/admin/ingredients.astro`
- `components/ProductEditModal.jsx`

### Fase 4: Activar Quality + Catalog + Orders (Día 6-7)

```javascript
USE_GO_QUALITY: true
USE_GO_CATALOG: true
USE_GO_ORDERS: true
```

## 🔧 Cómo Usar

### Ejemplo: Migrar ComprasApp.jsx

```javascript
// 1. Importar helper
import { comprasApi } from '../utils/api-config';

// 2. Reemplazar URLs
// Antes:
fetch('/api/compras/get_compras.php')
fetch('/api/compras/get_saldo_disponible.php')

// Después:
fetch(comprasApi('/api/compras'))
fetch(comprasApi('/api/compras/saldo'))
```

### Mapeo de Rutas

| PHP | Go | Helper |
|-----|-----|--------|
| `/api/admin_auth.php` | `/api/auth/login` | `authApi('/api/auth/login')` |
| `/api/compras/get_compras.php` | `/api/compras` | `comprasApi('/api/compras')` |
| `/api/get_ingredientes.php` | `/api/ingredientes` | `inventoryApi('/api/ingredientes')` |
| `/api/checklist.php` | `/api/checklist` | `qualityApi('/api/checklist')` |
| `/api/create_order.php` | `/api/create_order` | `ordersApi('/api/create_order')` |

## 🧪 Testing

### 1. Activar flag
```javascript
USE_GO_AUTH: true
```

### 2. Rebuild frontend
```bash
cd caja
npm run build
```

### 3. Probar en dev
```bash
npm run dev
```

### 4. Si falla → Rollback
```javascript
USE_GO_AUTH: false  // Volver a PHP
```

## 📊 Progreso

- [ ] Auth (17 usos)
- [ ] Compras (9 usos)
- [ ] Inventory (19 usos)
- [ ] Quality (4 usos)
- [ ] Catalog (2 usos)
- [ ] Orders (3 usos)

**Total: 0/54 usos migrados (0%)**

## ⚠️ Notas

- **Siempre probar en dev primero**
- **Monitorear errores en consola**
- **Rollback inmediato si hay problemas**
- **Un módulo a la vez**
