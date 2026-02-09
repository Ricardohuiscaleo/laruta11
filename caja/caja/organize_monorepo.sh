#!/bin/bash

echo "🚀 ORGANIZANDO MONOREPO LARUTA11"
echo ""

cd /Users/ricardohuiscaleollafquen/ruta11caja

# Verificar que existen las 3 carpetas
if [ ! -d "digitalapp" ] || [ ! -d "laruta11cl" ]; then
    echo "❌ Error: Faltan carpetas digitalapp o laruta11cl"
    exit 1
fi

echo "✅ Carpetas encontradas:"
echo "   - ruta11caja (actual - caja.laruta11.cl)"
echo "   - digitalapp (app.laruta11.cl)"
echo "   - laruta11cl (laruta11.cl)"
echo ""

# Crear estructura temporal
cd ..
mkdir -p laruta11-temp
cd laruta11-temp

echo "📁 Creando estructura del monorepo..."

# Copiar las 3 carpetas con nombres claros
cp -r ../ruta11caja caja
cp -r ../ruta11caja/digitalapp app
cp -r ../ruta11caja/laruta11cl landing

# Limpiar carpetas duplicadas dentro de caja
rm -rf caja/digitalapp
rm -rf caja/laruta11cl

echo "✅ Estructura creada"
echo ""

# Crear README principal
cat > README.md << 'EOF'
# 🍔 La Ruta 11 - Monorepo

Sistema completo de La Ruta 11 con 3 aplicaciones.

## 📁 Estructura

```
laruta11/
├── caja/       → caja.laruta11.cl (Sistema admin/caja)
├── app/        → app.laruta11.cl (Menú clientes)
└── landing/    → laruta11.cl (Página principal)
```

## 🌐 Dominios

- **laruta11.cl** - Landing principal
- **app.laruta11.cl** - App de menú para clientes
- **caja.laruta11.cl** - Sistema interno (admin, comandas, inventario)

## 🚀 Deployment

Cada carpeta es un servicio independiente en Easypanel:

### Servicio 1: Caja (Admin)
```yaml
Name: ruta11-caja
Root: /caja
Build: npm install && npm run build
Start: node dist/server/entry.mjs
Domain: caja.laruta11.cl
```

### Servicio 2: App (Clientes)
```yaml
Name: ruta11-app
Root: /app
Build: npm install && npm run build
Start: node dist/server/entry.mjs
Domain: app.laruta11.cl
```

### Servicio 3: Landing
```yaml
Name: ruta11-landing
Root: /landing
Type: Static
Domain: laruta11.cl
```

## 🗄️ Base de Datos

Todas las apps comparten la misma base de datos MySQL:
- **DB**: u958525313_app
- **Host**: localhost (en VPS)

## 📝 Notas

- Limpieza masiva: 399 archivos obsoletos eliminados
- APIs activas: ~240
- Sistema optimizado para producción
EOF

# Crear .gitignore global
cat > .gitignore << 'EOF'
# Dependencies
node_modules/
vendor/

# Build
dist/
.astro/

# Environment
.env
.env.local
.env.production

# IDE
.vscode/
.idea/

# OS
.DS_Store
Thumbs.db

# Logs
*.log
npm-debug.log*

# Temp
*.tmp
*.temp
EOF

echo "📝 README y .gitignore creados"
echo ""

# Inicializar Git
echo "🔧 Inicializando Git..."
git init
git add .
git commit -m "Initial commit: Monorepo completo con 3 apps (caja + app + landing)"

# Conectar con GitHub
echo "🔗 Conectando con GitHub..."
git remote add origin https://github.com/Ricardohuiscaleo/laruta11.git
git branch -M main

echo ""
echo "✅ LISTO PARA SUBIR"
echo ""
echo "📊 Resumen:"
echo "   - Carpeta: laruta11-temp/"
echo "   - Apps: 3 (caja, app, landing)"
echo "   - Repo: github.com/Ricardohuiscaleo/laruta11"
echo ""
echo "🚀 Para subir a GitHub ejecuta:"
echo "   cd /Users/ricardohuiscaleollafquen/laruta11-temp"
echo "   git push -u origin main"
echo ""
echo "⚠️  IMPORTANTE: Después de verificar que todo está bien,"
echo "   puedes eliminar la carpeta ruta11caja original"
EOF

chmod +x /Users/ricardohuiscaleollafquen/ruta11caja/organize_monorepo.sh
