#!/bin/bash
# Limpiar archivos duplicados

cd "$(dirname "$0")"

echo "🧹 Limpiando archivos duplicados..."

# Eliminar archivos viejos
rm -f auth.go
rm -f compras.go
rm -f resources.go
rm -f handlers.go
rm -f handlers_*.go
rm -f main_optimized.go
rm -f handlers_optimized.go

echo "✅ Archivos duplicados eliminados"
echo ""
echo "📁 Estructura actual:"
ls -la

echo ""
echo "📂 Carpetas creadas:"
ls -d */ 2>/dev/null || echo "No hay carpetas"
