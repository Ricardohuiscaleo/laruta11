Cómo Configurar y Ejecutar el Proyecto
Sigue estos pasos para tener la aplicación funcionando en tu computadora con Visual Studio Code.

Requisitos Previos

Tener Node.js instalado (versión 18 o superior).

Tener Visual Studio Code o tu editor de código preferido.

Pasos

Crear la Carpeta del Proyecto:

Crea una nueva carpeta en tu computadora con el nombre que prefieras (por ejemplo, menu-la-ruta-11).

Abre esta carpeta con Visual Studio Code.

Crear la Estructura de Archivos:

Dentro de la carpeta principal, crea la siguiente estructura de archivos y carpetas. Puedes hacerlo manualmente o usando la terminal.

menu-la-ruta-11/
├── src/
│   ├── components/
│   │   └── MenuApp.jsx
│   └── pages/
│       └── index.astro
├── astro.config.mjs
├── package.json
├── tailwind.config.mjs
└── tsconfig.json  (Opcional, pero recomendado por Astro)

Copiar y Pegar el Código:

Abre cada uno de los archivos que creaste en VS Code.

Copia el contenido de los bloques de código que te proporcioné arriba y pégalo en su archivo correspondiente. Asegúrate de que el nombre del archivo coincida exactamente.

Crear tsconfig.json (Recomendado):

Crea un archivo llamado tsconfig.json en la raíz de tu proyecto y pega el siguiente contenido. Esto ayuda con el autocompletado y la verificación de tipos.

{
  "extends": "astro/tsconfigs/strict"
}

Instalar Dependencias:

Abre la terminal integrada en Visual Studio Code (Ctrl + ``  o Cmd + `` ).

Ejecuta el siguiente comando. Esto leerá tu archivo package.json e instalará todas las herramientas necesarias (Astro, React, Tailwind, etc.).

npm install

Iniciar el Servidor de Desarrollo:

Una vez que la instalación termine, ejecuta el siguiente comando en la misma terminal:

npm run dev

¡Listo!

La terminal te mostrará una dirección local, usualmente http://localhost:4321.

Abre esa dirección en tu navegador web para ver y probar tu aplicación de menú.

## Deployment en Hostinger (PWA)

### Preparar para Producción

1. Construye el proyecto:
```bash
npm run build
```

2. Sube la carpeta `dist` completa a tu hosting Hostinger via FTP/cPanel

3. Configura tu dominio para apuntar a la carpeta donde subiste los archivos

### Backend PHP/MySQL

1. Crea una carpeta `api` en tu hosting
2. Configura la conexión a MySQL en tus archivos PHP
3. Actualiza la URL en `components/api.js` con tu dominio real

### Características PWA Incluidas

- ✅ Optimizado para móviles
- ✅ Carga rápida
- ✅ Funciona offline (básico)
- ✅ Instalable como app
- ✅ Responsive design
- ✅ Sistema de Control de Calidad integrado
- ✅ Gestión de sesiones persistentes con cookies
- ✅ Analytics avanzado con tracking de usuarios

### Cache Busting

El sistema implementa técnicas de "Cache Busting" para garantizar que los datos mostrados siempre estén actualizados en tiempo real:

- ✅ **Timestamps únicos** en todas las llamadas API
- ✅ **Headers anti-caché** en respuestas del servidor
- ✅ **Botón refresh** que limpia caché completo
- ✅ **Parámetros dinámicos** para evitar caché del navegador
- ✅ **Versionado automático** de assets estáticos

## Estructura Actual del Proyecto

### Frontend (React/Astro)
- Aplicación de menú con tarjetas de productos
- Loader inicial con logo
- Carrito de compras funcional
- Modales para detalles de productos
- Navegación por categorías
- Imágenes locales para productos principales
- **Sistema de Control de Calidad** (`/admin/calidad`)
  - Checklists para Maestro Planchero y Cajero
  - Respuestas Sí/No con observaciones
  - Subida de fotos evidencia comprimidas
  - Cálculo automático de scores de calidad

### Backend PHP/MySQL

**Configuración de Base de Datos** (`config.php`):
- **Base de Datos Principal (APP)**: `u958525313_app`
- **Usuario**: `u958525313_app`
- **Contraseña**: `wEzho0-hujzoz-cevzin`
- **Servidor**: localhost
- API Gemini integrada
- Búsqueda automática de config.php hasta 5 niveles

**API Endpoints Disponibles** (`api/` - 80+ archivos PHP):

**Productos:**
- `get_productos.php` - Obtener productos
- `add_producto.php` - Agregar producto
- `update_producto.php` - Actualizar producto
- `create_producto.php` - Crear producto

**Ingredientes y Recetas:**
- `get_ingredientes.php` - Obtener ingredientes
- `save_ingrediente.php` - Guardar ingrediente
- `get_recetas.php` - Obtener recetas
- `update_receta.php` - Actualizar receta

**Categorías:**
- `get_categories.php` - Obtener categorías
- `save_category.php` - Guardar categoría
- `categorias_hardcoded.php` - Categorías predefinidas

**Proyecciones Financieras:**
- `proyeccion.php` - Proyección básica
- `proyeccion_v2.php` - Proyección v2
- `proyeccion_v3.php` - Proyección v3
- `get_proyeccion.php` - Obtener proyección
- `save_proyeccion.php` - Guardar proyección

**Ventas:**
- `registrar_venta.php` - Registrar venta
- `ventas_update.php` - Actualizar ventas
- `ventas_get_all.php` - Obtener todas las ventas

**Dashboard y KPIs:**
- `get_dashboard_kpis.php` - Obtener KPIs
- `setup_dashboard_tables.php` - Configurar tablas

**Testing y Debug:**
- `test_api.php` - Probar API
- `test_connection.php` - Probar conexión
- `debug_db_connection.php` - Debug conexión DB

**Control de Calidad:**
- `get_questions.php` - Obtener preguntas por rol desde `quality_questions`
- `save_checklist.php` - Guardar checklists en `quality_checklists`
- `get_quality_score.php` - Obtener promedio de calidad para dashboard
- Tabla `quality_questions` - 20 preguntas (14 planchero + 6 cajero)
- Tabla `quality_checklists` - Almacena respuestas y scores
- Compresión automática de imágenes evidencia
- Integración con sistema de subida AWS existente

**Setup y Configuración:**
- `setup_tables.php` - Configurar tablas
- `check_config.php` - Verificar configuración
- `save_config.php` - Guardar configuración

### Imágenes de Productos (`public/`)
- `icon.png` - Logo de la aplicación
- `Completo-italiano.png` - Completo Tradicional
- `completo-talquino.png` - Completo Talquino
- `salchi-papas.png` - Salchipapas
- `papas-ruta11.png` - Papas Ruta 11
- `toma-provoleta.png` - Tomahawk Provoleta
- `tomahawk-full.png` - Tomahawk Full Ruta 11

## Sistema de Control de Calidad

### Funcionalidades
- **Checklists Diarios**: Maestro Planchero (14 preguntas) y Cajero (6 preguntas)
- **Secciones Organizadas**: Pre-Servicio, Durante Servicio, Post-Servicio
- **Respuestas Estructuradas**: Sí/No + observaciones de texto libre
- **Evidencia Fotográfica**: Compresión automática y subida a AWS S3
- **Scoring Automático**: Cálculo de porcentaje de calidad por rol
- **Dashboard Integration**: Promedio de calidad visible en dashboard principal
- **Histórico**: Un checklist por día por rol (actualizable)
- **Base de Datos**: Usa `u958525313_app` (no Calcularuta11)

### Base de Datos

**Tabla de Preguntas:**
```sql
CREATE TABLE quality_questions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    role ENUM('planchero', 'cajero') NOT NULL,
    question TEXT NOT NULL,
    requires_photo TINYINT(1) DEFAULT 0,
    order_index INT NOT NULL,
    active TINYINT(1) DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Tabla de Respuestas:**
```sql
CREATE TABLE quality_checklists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    role ENUM('planchero', 'cajero') NOT NULL,
    checklist_date DATE NOT NULL,
    responses JSON NOT NULL,
    total_questions INT NOT NULL,
    passed_questions INT NOT NULL,
    score_percentage DECIMAL(5,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY unique_role_date (role, checklist_date)
);
```

### Acceso y Características
- **URL**: `/admin/calidad` (integrado en SaaS admin)
- **Navegación**: Botón "Control Calidad" en menú admin
- **Responsive**: Optimizado para móviles y tablets
- **Acordeones**: Secciones colapsables para mejor UX
- **Progreso Visual**: Barra de progreso en tiempo real
- **Persistencia**: LocalStorage guarda progreso automáticamente
- **Dashboard**: Métrica "Calidad Promedio" en dashboard principal
- **API Endpoints**: 3 APIs dedicados para funcionalidad completa

