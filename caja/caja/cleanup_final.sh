#!/bin/bash

echo "🗑️  LIMPIEZA MASIVA DE ARCHIVOS OBSOLETOS"
echo ""

deleted=0
not_found=0

# ARCHIVOS RAÍZ
echo "📁 Limpiando raíz del proyecto..."

files_root=(
"analyze_project.js" "analyze_project.py" "analyze-project.js" "detailed-analysis.js"
"project-summary.js" "quick-stats.js" "generate-tech-report.js" "download-images.js"
"concurso-tracking.js" "test-job-application.php" "test-technical-report.php"
"test_sync.php" "test_tuu_api.php" "simulate_mysql_insert.php" "simulate_tuu_data.php"
"cleanup_apis.sh" "cleanup_complete.sh" "cleanup_obsolete_databases.mjs"
"cleanup_selective.php" "cleanup_smart.sh" "cleanup_unused_apis.php" "cleanup.py"
"cleanup.sh" "delete_obsolete_apis.mjs" "delete_phase2.mjs" "delete_phase3.mjs"
"delete_unused_apis.js" "delete_unused_apis.mjs" "execute_cleanup.mjs" "preview_delete.mjs"
"add_cafe_te_subcategories.sql" "add_empanadas_subcategory.sql" "add_score_lata.sql"
"check_categories_and_add_empanadas.sql" "complete_analytics_tables.sql"
"concurso_tracking.sql" "count_webpay_real.sql" "create_combos_subcategories.sql"
"create_extras_and_notes_tables.sql" "create_tuu_orders_table.sql" "insert_sample_data.sql"
"setup_app_database_real.sql" "setup_app_database.sql" "setup_interview_questions.sql"
"test_order_cleanup.sql" "update_existing_db.sql" "verificar_combos.sql"
"galaga-simple.html" "galaga.html" "test_import.html" "test-skills.html"
"ANALISIS_APIS_ELIMINAR.md" "ANALISIS_FASE2.md" "ANALISIS_FINAL.md"
"BASES_DATOS_OBSOLETAS.md" "Copia de MIGRACION-VPS.md" "api-openfactura-tuu.md"
"api-tuu.md" "DELIVERY_SYSTEM_README.md" "INSTALL_DOMPDF.md" "README_INVENTARIO.md"
"README-DELIVERY-SYSTEM.md" "README-EMAIL-SYSTEM.md" "README-TUU-PAGO-ONLINE.md"
"setup_cron.md" "TUU_SYSTEM_README.md" "config_root_update.php" "fix_config_paths.sh"
"install-ocr.sh" "sql_add_order_reference.txt" "project-analysis.json" "deploy.sh"
"ARCHIVOS_OBSOLETOS.md"
)

for file in "${files_root[@]}"; do
  if [ -f "$file" ]; then
    rm "$file"
    ((deleted++))
  else
    ((not_found++))
  fi
done

# CARPETAS COMPLETAS
echo "📁 Eliminando carpetas completas..."

folders=(
"backup_deleted_apis"
"tuu-pluguin"
"public/jobs"
"src/pages/jobs"
"src/pages/jobsTracker"
"src/pages/concurso_disabled"
"api/jobs"
"api/tracker"
"api/food_trucks"
"api/test"
)

for folder in "${folders[@]}"; do
  if [ -d "$folder" ]; then
    rm -rf "$folder"
    echo "  ✅ Eliminada carpeta: $folder"
    ((deleted++))
  else
    ((not_found++))
  fi
done

# ARCHIVOS ESPECÍFICOS EN SUBCARPETAS
echo "📁 Limpiando archivos específicos..."

specific_files=(
"public/js/candidate-detail.js"
"public/js/job-application.js"
"public/js/jobs-tracker.js"
"public/js/kanban.js"
"public/js/keywords.js"
"src/components/MenuApp.jsx.backup_20250926_082156"
"src/components/ProductEditModal.jsx.backup"
)

for file in "${specific_files[@]}"; do
  if [ -f "$file" ]; then
    rm "$file"
    ((deleted++))
  else
    ((not_found++))
  fi
done

echo ""
echo "📊 RESUMEN:"
echo "✅ Eliminados: $deleted archivos/carpetas"
echo "⚠️  No encontrados: $not_found"
echo ""
echo "🎉 Limpieza completada"
