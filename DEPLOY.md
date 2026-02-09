# 🚀 Despliegue en Easypanel

## ✅ Archivos preparados

- ✅ `.gitignore` creado
- ✅ `.env.example` en cada carpeta (app, caja, landing)
- ✅ `config.php` modificados para usar variables de entorno
- ✅ `SECRETS.txt` con todas las credenciales (NO SUBIR A GITHUB)
- ✅ Backups de config originales: `*.BACKUP`

## 📝 Pasos para subir a GitHub

1. **Remover archivos con secretos del commit anterior:**
```bash
git rm --cached app/config.php caja/caja/config.php landing/config.php
git commit -m "Remove config files with secrets"
```

2. **Agregar cambios nuevos:**
```bash
git add .gitignore app/.env.example caja/.env.example landing/.env.example
git add app/config.php caja/caja/config.php landing/config.php
git commit -m "Add environment variables support"
```

3. **Subir a GitHub:**
```bash
git push origin main
```

## 🔧 Configuración en Easypanel

Para cada servicio (app, caja, landing), agrega estas variables de entorno desde `SECRETS.txt`:

### Variables comunes para todos:
- `PUBLIC_SUPABASE_URL`
- `PUBLIC_SUPABASE_ANON_KEY`
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `S3_BUCKET`
- `S3_REGION`

### Variables específicas de APP y CAJA:
- `TUU_API_KEY`
- `TUU_ONLINE_RUT`
- `TUU_ONLINE_SECRET`
- `APP_DB_HOST`, `APP_DB_NAME`, `APP_DB_USER`, `APP_DB_PASS`
- Todas las credenciales OAuth de Google
- Credenciales de admin

### Variables específicas de LANDING:
- `GOOGLE_MAPS_API_KEY`
- `DB_HOST`, `DB_NAME`, `DB_USER`, `DB_PASS`

## 🌐 Dominios en Easypanel

- **laruta11.cl** → servicio `landing`
- **app.laruta11.cl** → servicio `app`
- **caja.laruta11.cl** → servicio `caja`

## ⚠️ IMPORTANTE

- **NO subas** `SECRETS.txt` a GitHub
- **NO subas** archivos `*.BACKUP` a GitHub
- **NO subas** archivos `config.php` con valores reales
- Guarda `SECRETS.txt` en un lugar seguro (1Password, Bitwarden, etc.)
