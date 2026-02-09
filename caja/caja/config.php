<?php
// Configuración de API Keys y variables de entorno
// IMPORTANTE: Mover este archivo fuera de public/ en producción
// Para testear: config.php?test=1

// Cargar variables de entorno
require_once __DIR__ . '/load-env.php';

$config = [
    'PUBLIC_SUPABASE_URL' => getenv('PUBLIC_SUPABASE_URL') ?: '',
    'PUBLIC_SUPABASE_ANON_KEY' => getenv('PUBLIC_SUPABASE_ANON_KEY') ?: '',

    'gemini_api_key' => getenv('GEMINI_API_KEY') ?: '',

    'unsplash_access_key' => getenv('UNSPLASH_ACCESS_KEY') ?: '',

    'google_calendar_api_key' => getenv('GOOGLE_CALENDAR_API_KEY') ?: '',
    'google_client_id' => getenv('GOOGLE_CLIENT_ID') ?: '',
    'google_client_secret' => getenv('GOOGLE_CLIENT_SECRET') ?: '',

    // === RUTA11 APP CONFIGURATION ===
    // OAuth para Ruta11 App principal (dedicado)
    'ruta11_google_client_id' => getenv('RUTA11_GOOGLE_CLIENT_ID') ?: '',
    'ruta11_google_client_secret' => getenv('RUTA11_GOOGLE_CLIENT_SECRET') ?: '',
    'ruta11_google_redirect_uri' => getenv('RUTA11_GOOGLE_REDIRECT_URI') ?: 'https://app.laruta11.cl/api/auth/google/callback.php',
    
    // OAuth para Ruta11 App (solo app DB)
    'ruta11_app_redirect_uri' => getenv('RUTA11_APP_REDIRECT_URI') ?: 'https://app.laruta11.cl/api/auth/google/app_callback.php',

    // OAuth para Ruta11 Jobs (proceso de selección)
    'ruta11_jobs_client_id' => getenv('RUTA11_JOBS_CLIENT_ID') ?: '',
    'ruta11_jobs_client_secret' => getenv('RUTA11_JOBS_CLIENT_SECRET') ?: '',
    'ruta11_jobs_redirect_uri' => getenv('RUTA11_JOBS_REDIRECT_URI') ?: 'https://app.laruta11.cl/api/auth/google/jobs_callback.php',
    
    // OAuth para Ruta11 Jobs Tracker (dashboard candidatos)
    'ruta11_tracker_client_id' => getenv('RUTA11_TRACKER_CLIENT_ID') ?: '',
    'ruta11_tracker_client_secret' => getenv('RUTA11_TRACKER_CLIENT_SECRET') ?: '',
    'ruta11_tracker_redirect_uri' => getenv('RUTA11_TRACKER_REDIRECT_URI') ?: 'https://app.laruta11.cl/api/auth/google/tracker_callback.php',

    // Google Maps API para Ruta11
    'ruta11_google_maps_api_key' => getenv('RUTA11_GOOGLE_MAPS_API_KEY') ?: '',
    
    // OAuth Gmail para envío de emails de postulaciones
    'gmail_client_id' => getenv('GMAIL_CLIENT_ID') ?: '',
    'gmail_client_secret' => getenv('GMAIL_CLIENT_SECRET') ?: '',
    'gmail_redirect_uri' => getenv('GMAIL_REDIRECT_URI') ?: 'https://app.laruta11.cl/api/auth/gmail/callback.php',
    'gmail_sender_email' => getenv('GMAIL_SENDER_EMAIL') ?: '',

    // TUU.cl Payment Gateway Configuration
    'tuu_api_key' => getenv('TUU_API_KEY') ?: '',
    
    // TUU Pago Online (Webpay) - PRODUCCIÓN
    'tuu_online_rut' => getenv('TUU_ONLINE_RUT') ?: '',
    'tuu_online_secret' => getenv('TUU_ONLINE_SECRET') ?: '',
    'tuu_online_env' => getenv('TUU_ONLINE_ENV') ?: 'production',
    'tuu_environment' => getenv('TUU_ENVIRONMENT') ?: 'dev',
    
    // Múltiples dispositivos POS
    'tuu_devices' => [
        'pos1' => [
            'serial' => '6010B232541610747',
            'name' => 'POS Principal - La Ruta 11',
            'location' => 'Mostrador Principal'
        ],
        'pos2' => [
            'serial' => '6010B232541609909',
            'name' => 'POS Secundario - La Ruta 11',
            'location' => 'Caja 2'
        ]
    ],
    
    'tuu_device_serial' => getenv('TUU_DEVICE_SERIAL') ?: '',

    // Base de datos MySQL para Ruta11 App (Nueva)
    'app_db_host' => getenv('APP_DB_HOST') ?: 'localhost',
    'app_db_name' => getenv('APP_DB_NAME') ?: '',
    'app_db_user' => getenv('APP_DB_USER') ?: '',
    'app_db_pass' => getenv('APP_DB_PASS') ?: '',
    
    // Credenciales Admin
    'admin_users' => [
        'admin' => getenv('ADMIN_PASSWORD') ?: '',
        'ricardo' => getenv('RICARDO_PASSWORD') ?: '',
        'manager' => getenv('MANAGER_PASSWORD') ?: '',
        'ruta11' => getenv('RUTA11_PASSWORD') ?: ''
    ],
    
    // Credenciales Inventario Móvil
    'inventario_user' => getenv('INVENTARIO_USER') ?: 'inventario',
    'inventario_password' => getenv('INVENTARIO_PASSWORD') ?: '',
    
    // Credenciales Externas
    'external_credentials' => [
        'pedidosya' => [
            'platform' => 'PedidosYA (Gowin)',
            'email' => getenv('PEDIDOSYA_EMAIL') ?: '',
            'password' => getenv('PEDIDOSYA_PASSWORD') ?: ''
        ],
        'instagram' => [
            'platform' => 'Instagram',
            'email' => getenv('INSTAGRAM_EMAIL') ?: '',
            'password' => getenv('INSTAGRAM_PASSWORD') ?: ''
        ],
        'tuu_platform' => [
            'platform' => 'TUU Platform',
            'email' => getenv('TUU_PLATFORM_EMAIL') ?: '',
            'password' => getenv('TUU_PLATFORM_PASSWORD') ?: ''
        ]
    ],
    
    // Configuración S3 AWS
    'aws_access_key_id' => getenv('AWS_ACCESS_KEY_ID') ?: '',
    'aws_secret_access_key' => getenv('AWS_SECRET_ACCESS_KEY') ?: '',
    's3_bucket' => getenv('S3_BUCKET') ?: 'laruta11-images',
    's3_region' => getenv('S3_REGION') ?: 'us-east-1',
    's3_url' => getenv('S3_URL') ?: 'https://laruta11-images.s3.amazonaws.com'
];

// Test de configuración
if (isset($_GET['test'])) {
    header('Content-Type: application/json');
    
    $results = [];
    $results['tuu_api_key'] = $config['tuu_api_key'] ? 'OK' : 'MISSING';
    $results['tuu_device'] = $config['tuu_device_serial'];
    
    echo json_encode($results, JSON_PRETTY_PRINT);
    exit;
}

return $config;
?>