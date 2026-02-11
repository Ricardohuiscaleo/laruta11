⚠️ 24 APIs endpoint no están siendo usadas.
🔧 Clases PHP / Helpers (2)
▼
/api/S3Manager.php
4 usos
Incluida en: /api/add_producto.php (PHP include), /api/checklist.php (PHP include), /api/update_producto.php (PHP include), /api/upload_image.php (PHP include)
/api/youtube_helper.php
1 usos
Incluida en: /api/save_youtube_live.php (PHP include)
✅ APIs USADAS (168)
▼

/api/add_concursante_manual.php
2 usos 1 tablas
Usada en: pages/admin/concurso-stats.astro, pages/admin/index.astro
Tablas: concurso_registros
/api/add_producto.php
2 usos 1 tablas
Usada en: pages/admin/index.astro, pages/admin/inventario.astro
Tablas: products
/api/add_review.php
1 usos 2 tablas
Usada en: components/ReviewsModal.jsx
Tablas: products, reviews
/api/admin/get_keys.php
1 usos
Usada en: pages/admin/keys.astro
/api/admin/verify_keys_access.php
1 usos
Usada en: pages/admin/keys.astro
✅ /api/admin_auth.php → GO /api/auth/login
8 usos
Usada en: components/AdminSPA.jsx, pages/admin/analytics.astro, pages/admin/app.astro, pages/admin/caja-config.astro, pages/admin/dashboard.astro, pages/admin/index.astro, pages/admin/login.astro, pages/admin/test.astro
/api/admin_dashboard.php
3 usos 1 tablas
Usada en: components/AdminDashboard.jsx, components/AdminPanel.jsx, pages/admin/test.astro
Tablas: products
✅ /api/admin_logout.php → GO /api/auth/logout
2 usos
Usada en: components/AdminSPA.jsx, pages/admin/index.astro
/api/api_status.php
2 usos
Usada en: components/ApiMonitor.jsx, pages/admin/test.astro
/api/app/check_tables.php
1 usos
Usada en: pages/admin/test.astro
/api/app/get_analytics.php
3 usos 4 tablas
Usada en: pages/admin/app.astro, pages/admin/dashboard.astro, pages/admin/test.astro
Tablas: site_visits, app_users, user_orders, products
/api/app/get_user_behavior.php
3 usos 4 tablas
Usada en: pages/admin/analytics.astro, pages/admin/dashboard.astro, pages/admin/index.astro
Tablas: product_analytics, user_interactions, user_journey, site_visits
/api/app/track_interaction.php
1 usos 2 tablas
Usada en: pages/index.astro
Tablas: user_interactions, product_analytics
/api/app/track_journey.php
1 usos 1 tablas
Usada en: pages/index.astro
Tablas: user_journey
/api/app/track_visit.php
2 usos 1 tablas
Usada en: pages/admin/test.astro, pages/index.astro
Tablas: site_visits
/api/approve_militar_rl6.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: usuarios
✅ /api/auth/check_session.php → GO /api/auth/check
3 usos
Usada en: components/CheckoutApp.jsx, components/MenuApp.jsx, pages/admin/test.astro
/api/auth/get_profile.php
1 usos 1 tablas
Usada en: components/modals/ProfileModal.jsx
Tablas: usuarios
/api/auth/gmail/auto_refresh.php
1 usos
Usada en: /api/cron/refresh_gmail_token.php (PHP include)
/api/auth/gmail/refresh_token.php
1 usos
Usada en: /api/auth/gmail/auto_refresh.php (PHP include)
✅ /api/auth/login_comandas.php → GO /api/auth/login
1 usos
Usada en: pages/comandas/index.astro
✅ /api/auth/login_v2.php → GO /api/auth/login
1 usos 1 tablas
Usada en: pages/login.astro
Tablas: cashiers
✅ /api/auth/logout.php → GO /api/auth/logout
2 usos
Usada en: components/MenuApp.jsx, components/modals/ProfileModal.jsx
/api/auth/update_profile.php
1 usos 1 tablas
Usada en: components/modals/ProfileModal.jsx
Tablas: usuarios
/api/bulk_adjust_price.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: products
/api/bulk_delete_products.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: products
/api/bulk_edit_products.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: productos
/api/bulk_update_products.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: products
/api/calculate_product_cost.php
1 usos 3 tablas
Usada en: pages/admin/edit-product.astro
Tablas: product_recipes, products, ingredients
/api/cancel_order.php
2 usos 5 tablas
Usada en: components/MiniComandas.jsx, pages/comandas/index.astro
Tablas: tuu_orders, inventory_transactions, ingredients, products, refunds
✅ /api/categories.php → GO /api/categories
8 usos 1 tablas
Usada en: components/ProductEditModal.jsx, components/ProductsManager.jsx, components/api.js, pages/admin/edit-product.astro, pages/admin/index.astro, pages/admin/inventario.astro, pages/admin/test.astro, pages/catalogo/index.astro
Tablas: categories
✅ /api/check_admin_auth.php → GO /api/auth/check
7 usos
Usada en: components/AdminSPA.jsx, pages/admin/analytics.astro, pages/admin/app.astro, pages/admin/caja-config.astro, pages/admin/dashboard.astro, pages/admin/index.astro, pages/admin/test.astro
/api/check_all_connections.php
1 usos
Usada en: pages/admin/caja-config.astro
/api/check_tables.php
1 usos
Usada en: pages/admin/test.astro
/api/check_tracking_data.php
2 usos 4 tablas
Usada en: pages/admin/dashboard.astro, pages/admin/index.astro
Tablas: site_visits, user_interactions, user_journey, product_analytics
/api/check_visits_table.php
2 usos 1 tablas
Usada en: pages/admin/dashboard.astro, pages/admin/index.astro
Tablas: site_visits
✅ /api/checklist.php → GO /api/checklist
4 usos 3 tablas
Usada en: components/ChecklistApp.jsx, components/ChecklistsListener.jsx, components/MiniComandas.jsx, pages/admin/calidad.astro
Tablas: checklists, checklist_items, checklist_templates
/api/cleanup_fake_data.php
1 usos 5 tablas
Usada en: pages/admin/index.astro
Tablas: usuarios, app_visits, productos, ventas, CURRENT_TIMESTAMP
✅ /api/compras/delete_compra.php → GO DELETE /api/compras/:id
1 usos 4 tablas
Usada en: components/ComprasApp.jsx
Tablas: compras_detalle, compras, ingredients, products
✅ /api/compras/get_compras.php → GO /api/compras
1 usos 2 tablas
Usada en: components/ComprasApp.jsx
Tablas: compras, compras_detalle
✅ /api/compras/get_historial_saldo.php → GO /api/compras/historial-saldo
1 usos 2 tablas
Usada en: components/ComprasApp.jsx
Tablas: tuu_orders, compras
✅ /api/compras/get_items_compra.php → GO /api/compras/items
1 usos 3 tablas
Usada en: components/ComprasApp.jsx
Tablas: ingredients, products, categories
✅ /api/compras/get_precio_historico.php → GO /api/compras/precio-historico
1 usos 2 tablas
Usada en: components/ComprasApp.jsx
Tablas: compras_detalle, compras
✅ /api/compras/get_proveedores.php → GO /api/compras/proveedores
1 usos 1 tablas
Usada en: components/ComprasApp.jsx
Tablas: compras
✅ /api/compras/get_saldo_disponible.php → GO /api/compras/saldo
1 usos 2 tablas
Usada en: components/ComprasApp.jsx
Tablas: tuu_orders, compras
✅ /api/compras/registrar_compra.php → GO POST /api/compras
1 usos 7 tablas
Usada en: components/ComprasApp.jsx
Tablas: ingredients, products, compras_detalle, tuu_orders, compras, capital_trabajo, egresos_compras
✅ /api/compras/upload_respaldo.php → GO POST /api/compras/:id/respaldo
1 usos 1 tablas
Usada en: components/ComprasApp.jsx
Tablas: compras
/api/confirm_transfer_payment.php
1 usos 7 tablas
Usada en: components/MiniComandas.jsx
Tablas: tuu_orders, tuu_order_items, caja_movimientos, product_recipes, products, inventory_transactions, ingredients
/api/create_ingredient.php
1 usos 1 tablas
Usada en: pages/admin/edit-product.astro
Tablas: ingredients
✅ /api/create_order.php → GO POST /api/create_order
3 usos 6 tablas
Usada en: components/CheckoutApp.jsx, components/MenuApp.jsx, components/OrderPOSApp.jsx
Tablas: product_recipes, products, caja_movimientos, tuu_orders, tuu_order_items, ingredients
✅ /api/delete_category.php → GO DELETE /api/categories/:id
1 usos 1 tablas
Usada en: pages/admin/inventario.astro
Tablas: categories
/api/delete_combo.php
2 usos 1 tablas
Usada en: pages/admin/combos.astro, pages/admin/index.astro
Tablas: combos
/api/delete_concursante.php
2 usos 1 tablas
Usada en: pages/admin/concurso-stats.astro, pages/admin/index.astro
Tablas: concurso_registros
/api/delete_image_from_gallery.php
1 usos 1 tablas
Usada en: pages/admin/edit-product.astro
Tablas: products
✅ /api/delete_ingrediente.php → GO DELETE /api/ingredientes/:id
2 usos 1 tablas
Usada en: pages/admin/ingredients.astro, pages/admin/inventario.astro
Tablas: ingredients
/api/delete_producto.php
2 usos 1 tablas
Usada en: pages/admin/index.astro, pages/admin/inventario.astro
Tablas: products
/api/generate_whatsapp_message.php
1 usos
Usada en: pages/cash-pending.astro
/api/get_analytics.php
3 usos 4 tablas
Usada en: pages/admin/app.astro, pages/admin/dashboard.astro, pages/admin/test.astro
Tablas: site_visits, app_users, user_orders, products
✅ /api/get_categories.php → GO /api/categories
4 usos 1 tablas
Usada en: components/ProductEditModal.jsx, pages/admin/inventario.astro, pages/admin/test.astro, pages/catalogo/index.astro
Tablas: categories
/api/get_combos.php
3 usos 5 tablas
Usada en: components/modals/ComboModal.jsx, pages/admin/combos.astro, pages/admin/index.astro
Tablas: combos, combo_items, combo_selections, categories, products
/api/get_concurso_stats.php
2 usos 1 tablas
Usada en: pages/admin/concurso-stats.astro, pages/admin/index.astro
Tablas: concurso_tracking
/api/get_dashboard_analytics.php
1 usos 6 tablas
Usada en: pages/admin/index.astro
Tablas: site_visits, usuarios, products, product_analytics, user_interactions, user_journey
/api/get_dashboard_cards.php
1 usos 6 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_orders, compras, ingredients, products, tuu_order_items, product_recipes
/api/get_delivery_fee.php
1 usos 1 tablas
Usada en: components/CheckoutApp.jsx
Tablas: food_trucks
/api/get_financial_reports.php
1 usos 2 tablas
Usada en: pages/admin/reportes.astro
Tablas: tuu_orders, tuu_order_items
/api/get_historial_caja.php
1 usos 1 tablas
Usada en: components/modals/SaldoCajaModal.jsx
Tablas: caja_movimientos
✅ /api/get_ingredientes.php → GO /api/ingredientes
5 usos 1 tablas
Usada en: components/MermasApp.jsx, pages/admin/edit-product.astro, pages/admin/ingredients.astro, pages/admin/inventario.astro, pages/admin/mermas.astro
Tablas: ingredients
/api/get_ingredients.php
2 usos 1 tablas
Usada en: pages/admin/edit-product.astro, pages/inventario/index.astro
Tablas: ingredients
/api/get_menu_products.php
1 usos 3 tablas
Usada en: components/MenuApp.jsx
Tablas: products, subcategories, reviews
/api/get_mermas.php
1 usos 3 tablas
Usada en: components/MermasApp.jsx
Tablas: mermas, ingredients, products
/api/get_militares_rl6.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: usuarios
/api/get_month_comparison.php
1 usos
Usada en: pages/admin/index.astro
/api/get_movimientos_caja.php
1 usos 1 tablas
Usada en: components/ArqueoResumen.jsx
Tablas: caja_movimientos
/api/get_nearby_trucks.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: food_trucks
/api/get_order_notifications.php
2 usos 2 tablas
Usada en: components/MenuApp.jsx, components/OrderNotifications.jsx
Tablas: tuu_orders, order_notifications
/api/get_participantes_concurso.php
2 usos 1 tablas
Usada en: pages/admin/concurso-stats.astro, pages/admin/index.astro
Tablas: concurso_registros
/api/get_pending_orders.php
1 usos 1 tablas
Usada en: components/OrderManagement.jsx
Tablas: orders
/api/get_previous_month_summary.php
1 usos 2 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_orders, tuu_order_items
/api/get_product_recipe.php
1 usos 2 tablas
Usada en: pages/admin/edit-product.astro
Tablas: product_recipes, ingredients
/api/get_product_recipe_details.php
1 usos 2 tablas
Usada en: pages/inventario/index.astro
Tablas: product_recipes, ingredients
/api/get_productos.php
10 usos 2 tablas
Usada en: components/MermasApp.jsx, components/ProductEditModal.astro, components/ProductEditModal.jsx, pages/admin/caja-config.astro, pages/admin/edit-product.astro, pages/admin/index.astro, pages/admin/inventario.astro, pages/admin/test.astro, pages/catalogo/index.astro, pages/inventario/index.astro
Tablas: products, combos
/api/get_purchase_plan.php
1 usos 6 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_order_items, product_recipes, products, tuu_orders, a, ingredients
/api/get_purchase_recommendation.php
1 usos 2 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_order_items, tuu_orders
/api/get_quality_score.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: quality_checklists
/api/get_questions.php
1 usos 1 tablas
Usada en: pages/admin/calidad.astro
Tablas: quality_questions
/api/get_reviews.php
1 usos 2 tablas
Usada en: components/ReviewsModal.jsx
Tablas: reviews, products
/api/get_saldo_caja.php
2 usos 2 tablas
Usada en: components/ArqueoApp.jsx, components/ArqueoResumen.jsx
Tablas: caja_movimientos, tuu_orders
/api/get_sales_analytics.php
1 usos 3 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_orders, tuu_order_items, products
/api/get_sales_by_address.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_orders
/api/get_sales_by_period.php
1 usos
Usada en: pages/admin/index.astro
/api/get_sales_detail.php
2 usos 5 tablas
Usada en: components/ArqueoResumen.jsx, components/VentasDetalle.jsx
Tablas: tuu_orders, tuu_order_items, inventory_transactions, ingredients, products
/api/get_sales_summary.php
2 usos 1 tablas
Usada en: components/ArqueoApp.jsx, components/ArqueoResumen.jsx
Tablas: tuu_orders
/api/get_smart_projection.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_orders
/api/get_smart_projection_shifts.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: tuu_orders
/api/get_subcategories.php
3 usos 1 tablas
Usada en: components/ProductEditModal.jsx, pages/admin/edit-product.astro, pages/admin/index.astro
Tablas: subcategories
/api/get_technical_report.php
3 usos
Usada en: components/LiveMetrics.jsx, components/TechnicalReport.jsx, pages/admin/index.astro
/api/get_transfer_order.php
4 usos 2 tablas
Usada en: pages/card-pending.astro, pages/cash-pending.astro, pages/pedidosya-pending.astro, pages/transfer-pending.astro
Tablas: tuu_orders, tuu_order_items
/api/get_truck_schedules.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: food_truck_schedules
/api/get_truck_status.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: food_trucks
/api/get_user_detail.php
1 usos 3 tablas
Usada en: pages/admin/index.astro
Tablas: app_users, user_orders, user_order_items
/api/get_user_orders.php
1 usos 2 tablas
Usada en: components/MenuApp.jsx
Tablas: usuarios, tuu_orders
/api/get_users.php
3 usos 2 tablas
Usada en: pages/admin/index.astro, pages/admin/test.astro, pages/admin/users.astro
Tablas: app_users, user_orders
/api/get_ventas_turno.php
1 usos 2 tablas
Usada en: components/ArqueoResumen.jsx
Tablas: tuu_orders, tuu_order_items
/api/inventario_login.php
1 usos
Usada en: pages/inventario/index.astro
/api/location/calculate_delivery_time.php
1 usos
Usada en: components/MenuApp.jsx
/api/location/check_delivery_zone.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: delivery_zones
/api/location/geocode.php
1 usos
Usada en: components/MenuApp.jsx
/api/location/get_location.php
2 usos 1 tablas
Usada en: components/modals/ProfileModal.jsx, pages/admin/test.astro
Tablas: usuarios
/api/location/get_nearby_products.php
1 usos
Usada en: components/MenuApp.jsx
/api/location/save_location.php
1 usos 2 tablas
Usada en: components/MenuApp.jsx
Tablas: user_locations, usuarios
/api/notify_admin_payment.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: admin_notifications
/api/products.php
4 usos 2 tablas
Usada en: components/AdminPanel.jsx, components/MenuApp.jsx, components/ProductsManager.jsx, pages/admin/index.astro
Tablas: products, categories
/api/recalcular_stock_productos.php
1 usos 3 tablas
Usada en: pages/inventario/index.astro
Tablas: product_recipes, products, ingredients
/api/registrar_merma.php
2 usos 3 tablas
Usada en: components/MermasApp.jsx, pages/admin/mermas.astro
Tablas: ingredients, products, mermas
/api/registrar_movimiento_caja.php
1 usos 1 tablas
Usada en: components/modals/SaldoCajaModal.jsx
Tablas: caja_movimientos
/api/reprocess_order_inventory.php
1 usos 6 tablas
Usada en: pages/admin/pagos-tuu.astro
Tablas: product_recipes, products, tuu_orders, inventory_transactions, tuu_order_items, ingredients
/api/rl6_refund_credit.php
1 usos 6 tablas
Usada en: components/MiniComandas.jsx
Tablas: tuu_orders, inventory_transactions, ingredients, products, rl6_credit_transactions, usuarios
/api/robot_notifications.php
1 usos 1 tablas
Usada en: components/RobotAlerts.jsx
Tablas: robot_test_logs
/api/robot_test_tracking.php
2 usos
Usada en: pages/admin/dashboard.astro, pages/admin/index.astro
/api/save_category.php
1 usos 1 tablas
Usada en: pages/admin/inventario.astro
Tablas: categorias
/api/save_checklist.php
1 usos 2 tablas
Usada en: pages/admin/calidad.astro
Tablas: quality_checklists, responses
/api/save_combo.php
2 usos 4 tablas
Usada en: pages/admin/edit-product.astro, pages/admin/index.astro
Tablas: combos, products, combo_items, combo_selections
/api/save_image_url.php
1 usos 1 tablas
Usada en: pages/admin/edit-product.astro
Tablas: products
/api/save_ingrediente.php
3 usos 1 tablas
Usada en: components/ComprasApp.jsx, pages/admin/ingredients.astro, pages/admin/inventario.astro
Tablas: ingredients
/api/save_product_recipe.php
1 usos 1 tablas
Usada en: pages/admin/edit-product.astro
Tablas: product_recipes
/api/setup_ingredients.php
1 usos 1 tablas
Usada en: pages/admin/edit-product.astro
Tablas: ingredients
/api/simulate_visit.php
1 usos 1 tablas
Usada en: pages/admin/index.astro
Tablas: site_visits
/api/test_analytics_system.php
2 usos
Usada en: pages/admin/dashboard.astro, pages/admin/index.astro
/api/test_connection.php
1 usos
Usada en: pages/admin/test.astro
/api/test_inventory_deduction.php
1 usos 3 tablas
Usada en: pages/admin/test-inventory.astro
Tablas: products, product_recipes, ingredients
/api/toggle_like.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: products
/api/toggle_product_status.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: products
/api/track_usage.php
1 usos 2 tablas
Usada en: components/MenuApp.jsx
Tablas: user_metrics, usuarios
/api/tuu/check_connection.php
1 usos 2 tablas
Usada en: pages/admin/caja-config.astro
Tablas: tuu_pos_transactions, tuu_orders
/api/tuu/check_status.php
1 usos
Usada en: pages/admin/test.astro
/api/tuu/create_payment_direct.php
2 usos 2 tablas
Usada en: components/CheckoutApp.jsx, components/TUUPaymentIntegration.jsx
Tablas: tuu_order_items, tuu_orders
/api/tuu/create_payment_simple.php
1 usos
Usada en: pages/admin/test.astro
/api/tuu/delete_transaction.php
2 usos 1 tablas
Usada en: components/PagosTuu.jsx, pages/admin/pagos-tuu.astro
Tablas: tuu_orders
/api/tuu/get_comandas_v2.php
3 usos 5 tablas
Usada en: components/MiniComandas.jsx, components/OrdersListener.jsx, pages/comandas/index.astro
Tablas: tuu_orders, tuu_order_items, products, product_recipes, ingredients
/api/tuu/get_devices.php
2 usos
Usada en: components/TuuReportsAdmin.jsx, pages/admin/test.astro
/api/tuu/get_from_mysql.php
5 usos 5 tablas
Usada en: components/TUUTransactions.jsx, components/TuuReportsAdmin.jsx, pages/admin/dashboard.astro, pages/admin/index.astro, pages/admin/pagos-tuu.astro
Tablas: tuu_order_items, tuu_orders, inventory_transactions, products, ingredients
/api/tuu/get_order_delivery.php
1 usos 1 tablas
Usada en: pages/payment-success.astro
Tablas: tuu_orders
/api/tuu/get_order_products_with_extras.php
1 usos 2 tablas
Usada en: pages/payment-success.astro
Tablas: tuu_order_items, tuu_orders
/api/tuu/get_orders_with_inventory.php
1 usos 5 tablas
Usada en: components/PagosTuu.jsx
Tablas: tuu_orders, tuu_order_items, inventory_transactions, ingredients, products
/api/tuu/get_reports.php
1 usos
Usada en: pages/admin/test.astro
/api/tuu/get_shift_transactions.php
1 usos 6 tablas
Usada en: pages/admin/pagos-tuu.astro
Tablas: tuu_order_items, tuu_orders, inventory_transactions, product_recipes, products, ingredients
/api/tuu/register_success.php
1 usos 1 tablas
Usada en: pages/payment-success.astro
Tablas: tuu_orders
/api/tuu/save_delivery_info.php
1 usos 1 tablas
Usada en: components/CheckoutApp.jsx
Tablas: tuu_orders
/api/tuu/setup_remote_payments_table.php
1 usos 1 tablas
Usada en: pages/admin/caja-config.astro
Tablas: CURRENT_TIMESTAMP
/api/tuu/sync_reports.php
1 usos 1 tablas
Usada en: pages/admin/test.astro
Tablas: tuu_reports
/api/tuu/update_order_status.php
2 usos 1 tablas
Usada en: components/MiniComandas.jsx, pages/comandas/index.astro
Tablas: tuu_orders
/api/tuu-pagos-online/save_transaction.php
1 usos 1 tablas
Usada en: components/TuuPayment.jsx
Tablas: tuu_pagos_online
/api/update_cashier_profile.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: cashiers
/api/update_category.php
1 usos 1 tablas
Usada en: pages/admin/inventario.astro
Tablas: categories
/api/update_concursante.php
2 usos 1 tablas
Usada en: pages/admin/concurso-stats.astro, pages/admin/index.astro
Tablas: concurso_registros
/api/update_ingredient_stock.php
1 usos 1 tablas
Usada en: pages/inventario/index.astro
Tablas: ingredients
/api/update_ingrediente.php
1 usos 2 tablas
Usada en: pages/admin/inventario.astro
Tablas: ingrediente, ingredients
/api/update_ingrediente_full.php
1 usos 1 tablas
Usada en: pages/admin/ingredients.astro
Tablas: ingredients
/api/update_order_status.php
3 usos 2 tablas
Usada en: components/MiniComandas.jsx, components/OrderManagement.jsx, pages/comandas/index.astro
Tablas: tuu_orders, order_notifications
/api/update_producto.php
4 usos 1 tablas
Usada en: components/ProductEditModal.astro, components/ProductEditModal.jsx, pages/admin/edit-product.astro, pages/admin/inventario.astro
Tablas: products
/api/update_producto_stock.php
1 usos 1 tablas
Usada en: pages/inventario/index.astro
Tablas: products
/api/update_truck_schedule.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: food_truck_schedules
/api/update_truck_status.php
1 usos 1 tablas
Usada en: components/MenuApp.jsx
Tablas: food_trucks
/api/update_visits_table.php
2 usos
Usada en: pages/admin/dashboard.astro, pages/admin/index.astro
/api/upload_image.php
4 usos
Usada en: pages/admin/calidad.astro, pages/admin/concurso-stats.astro, pages/admin/edit-product.astro, pages/admin/index.astro
/api/users/get_user_detail.php
1 usos 3 tablas
Usada en: pages/admin/index.astro
Tablas: app_users, user_orders, user_order_items
/api/users/get_users.php
3 usos 2 tablas
Usada en: pages/admin/index.astro, pages/admin/test.astro, pages/admin/users.astro
Tablas: usuarios, tuu_orders
/api/verify_inventario_token.php
1 usos
Usada en: pages/inventario/index.astro