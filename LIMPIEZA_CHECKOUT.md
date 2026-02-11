# Plan de Limpieza - Código Muerto Checkout

**Fecha**: 2026-02-11  
**Decisión**: Solo usar `CheckoutApp.jsx` para caja

---

## ✅ Componente en Uso

**CheckoutApp.jsx** (1,013 líneas)
- ✅ Usado en sistema de caja
- ✅ Checkout completo con TUU
- ✅ **MANTENER**

---

## ❌ Código Muerto (Eliminar)

| Archivo | Líneas | Estado | Acción |
|---------|--------|--------|--------|
| `CheckoutWithTUU.jsx` | 140 | ❌ No usado | **ELIMINAR** |
| `MultiPOSCheckout.jsx` | 123 | ❌ No usado | **ELIMINAR** |
| `TUUCheckout.jsx` | 100 | ❌ No usado | **ELIMINAR** |
| `TuuNativeCheckout.jsx` | 199 | ❌ No usado | **ELIMINAR** |

**Total a eliminar**: 562 líneas de código muerto

---

## 🗑️ Comandos de Limpieza

```bash
cd /Users/ricardohuiscaleollafquen/laruta11/caja/src/components

# Eliminar checkouts no usados
rm CheckoutWithTUU.jsx
rm MultiPOSCheckout.jsx
rm TUUCheckout.jsx
rm TuuNativeCheckout.jsx

# Commit
git add -A
git commit -m "chore: remove unused checkout components (562 lines)"
git push
```

---

## 📊 Impacto

### Antes:
```
CheckoutApp.jsx          1,013 líneas ✅ USADO
CheckoutWithTUU.jsx        140 líneas ❌ NO USADO
MultiPOSCheckout.jsx       123 líneas ❌ NO USADO
TUUCheckout.jsx            100 líneas ❌ NO USADO
TuuNativeCheckout.jsx      199 líneas ❌ NO USADO
─────────────────────────────────────
TOTAL:                   1,575 líneas
```

### Después:
```
CheckoutApp.jsx          1,013 líneas ✅ USADO
─────────────────────────────────────
TOTAL:                   1,013 líneas
```

**Reducción**: 562 líneas (36% menos código)

---

## ✅ Beneficios

1. ✅ **Menos confusión** - Solo 1 checkout, no 5
2. ✅ **Build más rápido** - Menos archivos que compilar
3. ✅ **Bundle más pequeño** - Menos código en producción
4. ✅ **Más fácil de mantener** - Solo 1 archivo que actualizar
5. ✅ **Menos bugs** - Menos código = menos superficie de error

---

## 🎯 Próximo Paso

**¿Elimino los 4 archivos no usados?**

```bash
# Ejecutar limpieza
cd /Users/ricardohuiscaleollafquen/laruta11
rm caja/src/components/CheckoutWithTUU.jsx
rm caja/src/components/MultiPOSCheckout.jsx
rm caja/src/components/TUUCheckout.jsx
rm caja/src/components/TuuNativeCheckout.jsx
git add -A
git commit -m "chore: remove 562 lines of unused checkout code"
git push
```

**Confirma para proceder** ✅
