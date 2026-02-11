#!/bin/bash

API="https://websites-api-go-caja-r11.dj3bvg.easypanel.host"

echo "🧪 Testing Go API Endpoints"
echo "================================"

# Health
echo -e "\n✅ Health Check"
curl -s "$API/api/health" | jq

# Auth
echo -e "\n🔐 Auth - Login (admin)"
curl -s -X POST "$API/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"user":"admin","pass":"R11adm2025x7k9","type":"admin"}' | jq

echo -e "\n🔐 Auth - Check"
curl -s "$API/api/auth/check" -H "X-User: admin" | jq

echo -e "\n🔐 Auth - Logout"
curl -s -X POST "$API/api/auth/logout" | jq

# Compras
echo -e "\n💰 Compras - Get All"
curl -s "$API/api/compras" | jq '.success, (.compras | length)'

echo -e "\n💰 Compras - Get Items"
curl -s "$API/api/compras/items" | jq 'length'

echo -e "\n💰 Compras - Get Proveedores"
curl -s "$API/api/compras/proveedores" | jq

echo -e "\n💰 Compras - Get Saldo"
curl -s "$API/api/compras/saldo" | jq

echo -e "\n💰 Compras - Get Historial Saldo"
curl -s "$API/api/compras/historial-saldo" | jq '.success, (.movimientos | length)'

echo -e "\n💰 Compras - Get Precio Historico (ingrediente_id=1)"
curl -s "$API/api/compras/precio-historico?ingrediente_id=1" | jq

# Inventory
echo -e "\n📦 Inventory - Get Ingredientes"
curl -s "$API/api/ingredientes" | jq '.success, (.ingredientes | length)'

echo -e "\n📦 Inventory - Get Categories"
curl -s "$API/api/categories" | jq '.success, (.categories | length)'

# Quality
echo -e "\n✔️ Quality - Get Checklists"
curl -s "$API/api/checklist" | jq '.success, (.checklists | length)'

# Catalog
echo -e "\n📋 Catalog - Get Products"
curl -s "$API/api/products" | jq '.success, (.products | length)'

echo -e "\n📋 Catalog - Get Product by ID (id=1)"
curl -s "$API/api/products/1" | jq

# Orders
echo -e "\n📦 Orders - Get Pending"
curl -s "$API/api/orders/pending" | jq '.success, (.orders | length)'

echo -e "\n================================"
echo "✅ Tests completed!"
