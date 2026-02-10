#!/bin/bash

echo "🔄 REORGANIZANDO CARPETAS EN RUTA11CAJA"
echo ""

cd /Users/ricardohuiscaleollafquen/ruta11caja

# Renombrar carpetas
echo "📝 Renombrando carpetas..."

# Renombrar digitalapp → app
if [ -d "digitalapp" ]; then
    mv digitalapp app
    echo "   ✅ digitalapp → app"
fi

# Renombrar laruta11cl → landing
if [ -d "laruta11cl" ]; then
    mv laruta11cl landing
    echo "   ✅ laruta11cl → landing"
fi

# Crear carpeta caja y mover todo excepto app y landing
echo ""
echo "📁 Creando carpeta caja..."
mkdir -p caja-temp

# Mover todo excepto app, landing y caja-temp
for item in *; do
    if [ "$item" != "app" ] && [ "$item" != "landing" ] && [ "$item" != "caja-temp" ]; then
        mv "$item" caja-temp/
    fi
done

# Renombrar caja-temp a caja
mv caja-temp caja

echo "   ✅ Contenido movido a caja/"
echo ""

# Crear README
cat > README.md << 'EOF'
# 🍔 La Ruta 11 - Monorepo

Sistema completo con 3 aplicaciones.

## 📁 Estructura

- `/caja` - Sistema admin/caja (caja.laruta11.cl)
- `/app` - Menú clientes (app.laruta11.cl)
- `/landing` - Página principal (laruta11.cl)

## 🌐 Dominios

- **laruta11.cl** → landing/
- **app.laruta11.cl** → app/
- **caja.laruta11.cl** → caja/

## 🚀 Deploy en Easypanel

Cada carpeta = 1 servicio independiente
EOF

# Crear .gitignore
cat > .gitignore << 'EOF'
node_modules/
dist/
.astro/
.env
.env.local
.DS_Store
*.log
EOF

echo "📝 README.md y .gitignore creados"
echo ""
echo "✅ REORGANIZACIÓN COMPLETA"
echo ""
echo "📊 Estructura final:"
echo "   ruta11caja/"
echo "   ├── caja/       (caja.laruta11.cl)"
echo "   ├── app/        (app.laruta11.cl)"
echo "   ├── landing/    (laruta11.cl)"
echo "   ├── README.md"
echo "   └── .gitignore"
echo ""
echo "🚀 Para subir a GitHub:"
echo "   git init"
echo "   git add ."
echo "   git commit -m 'Monorepo completo: 3 apps'"
echo "   git remote add origin https://github.com/Ricardohuiscaleo/laruta11.git"
echo "   git branch -M main"
echo "   git push -u origin main"
