#!/bin/bash
# Script para limpiar cache corrupto de Vite/Astro

echo "🧹 Limpiando cache corrupto..."

# Matar procesos
pkill -f "astro" 2>/dev/null
pkill -f "vite" 2>/dev/null
pkill -f "node.*dev" 2>/dev/null

# Limpiar cache
rm -rf .astro
rm -rf node_modules/.vite
rm -rf node_modules/.cache
rm -rf dist

echo "✅ Cache limpiado"
echo ""
echo "📝 Ahora ejecuta:"
echo "   npm run dev"
