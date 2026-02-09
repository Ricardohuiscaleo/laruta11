<?php
// Configuración de API Keys y variables de entorno
// IMPORTANTE: Mover este archivo fuera de public/ en producción
// Para testear: config.php?test=1

// Cargar variables de entorno
require_once __DIR__ . '/load-env.php';

$config = [
    // APIs de terceros
    'google_maps_api_key' => getenv('GOOGLE_MAPS_API_KEY') ?: '',
    'whatsapp_api_token' => getenv('WHATSAPP_API_TOKEN') ?: '',
    'instagram_api_key' => getenv('INSTAGRAM_API_KEY') ?: '',
    'facebook_api_key' => getenv('FACEBOOK_API_KEY') ?: '',
    
    // AWS S3 para imágenes
    'aws_access_key_id' => getenv('AWS_ACCESS_KEY_ID') ?: '',
    'aws_secret_access_key' => getenv('AWS_SECRET_ACCESS_KEY') ?: '',
    'aws_region' => getenv('AWS_REGION') ?: 'us-east-1',
    's3_bucket' => getenv('S3_BUCKET') ?: 'laruta11-images',
    's3_url' => getenv('S3_URL') ?: 'https://laruta11-images.s3.amazonaws.com',
    
    // Base de datos
    'db_host' => getenv('DB_HOST') ?: 'localhost',
    'db_name' => getenv('DB_NAME') ?: 'laruta11_db',
    'db_user' => getenv('DB_USER') ?: '',
    'db_pass' => getenv('DB_PASS') ?: '',
    
    // Configuración de la aplicación
    'app_url' => getenv('APP_URL') ?: 'https://laruta11.cl',
    'app_env' => getenv('APP_ENV') ?: 'production',
    'debug' => getenv('DEBUG') === 'true',
    'logo_url' => 'https://laruta11-images.s3.amazonaws.com/menu/1755571382_test.jpg',
    'favicon_url' => 'https://laruta11-images.s3.amazonaws.com/menu/1755571382_test.jpg',
    
    // Email y notificaciones
    'smtp_host' => getenv('SMTP_HOST') ?: '',
    'smtp_user' => getenv('SMTP_USER') ?: '',
    'smtp_pass' => getenv('SMTP_PASS') ?: '',
    'contact_email' => 'hola@laruta11.cl',
    
    // Configuración de food trucks
    'default_location' => [
        'lat' => -33.4489,
        'lng' => -70.6693
    ],
    'business_hours' => [
        'monday' => ['11:00', '21:00'],
        'tuesday' => ['11:00', '21:00'],
        'wednesday' => ['11:00', '21:00'],
        'thursday' => ['11:00', '21:00'],
        'friday' => ['11:00', '21:00'],
        'saturday' => ['10:00', '22:00'],
        'sunday' => ['12:00', '20:00']
    ],
    
    // === CONFIGURACIÓN EXTENDIDA AGENTERAG ===
    // Supabase
    'PUBLIC_SUPABASE_URL' => getenv('PUBLIC_SUPABASE_URL') ?: '',
    'PUBLIC_SUPABASE_ANON_KEY' => getenv('PUBLIC_SUPABASE_ANON_KEY') ?: '',
    
    // APIs
    'gemini_api_key' => getenv('GEMINI_API_KEY') ?: '',
    'unsplash_access_key' => getenv('UNSPLASH_ACCESS_KEY') ?: '',
    
    // Google APIs
    'google_calendar_api_key' => getenv('GOOGLE_CALENDAR_API_KEY') ?: '',
    'google_client_id' => getenv('GOOGLE_CLIENT_ID') ?: '',
    'google_client_secret' => getenv('GOOGLE_CLIENT_SECRET') ?: '',
    
    // Base de datos MySQL - Booking
    'booking_db_host' => getenv('BOOKING_DB_HOST') ?: 'localhost',
    'booking_db_name' => getenv('BOOKING_DB_NAME') ?: '',
    'booking_db_user' => getenv('BOOKING_DB_USER') ?: '',
    'booking_db_pass' => getenv('BOOKING_DB_PASS') ?: '',
    
    // Base de datos MySQL - RAG
    'rag_db_host' => getenv('RAG_DB_HOST') ?: 'localhost',
    'rag_db_name' => getenv('RAG_DB_NAME') ?: '',
    'rag_db_user' => getenv('RAG_DB_USER') ?: '',
    'rag_db_pass' => getenv('RAG_DB_PASS') ?: '',
    
    // Base de datos MySQL - Ruta11Game
    'ruta11game_db_host' => getenv('RUTA11GAME_DB_HOST') ?: 'localhost',
    'ruta11game_db_name' => getenv('RUTA11GAME_DB_NAME') ?: '',
    'ruta11game_db_user' => getenv('RUTA11GAME_DB_USER') ?: '',
    'ruta11game_db_pass' => getenv('RUTA11GAME_DB_PASS') ?: '',
    
    // Base de datos MySQL - Calcularuta11
    'Calcularuta11_db_host' => getenv('CALCULARUTA11_DB_HOST') ?: 'localhost',
    'Calcularuta11_db_name' => getenv('CALCULARUTA11_DB_NAME') ?: '',
    'Calcularuta11_db_user' => getenv('CALCULARUTA11_DB_USER') ?: '',
    'Calcularuta11_db_pass' => getenv('CALCULARUTA11_DB_PASS') ?: '',
    
    // === RUTA11 APP CONFIGURATION ===
    // Base de datos MySQL - Usuarios Ruta11
    'ruta11_db_host' => getenv('RUTA11_DB_HOST') ?: 'localhost',
    'ruta11_db_name' => getenv('RUTA11_DB_NAME') ?: '',
    'ruta11_db_user' => getenv('RUTA11_DB_USER') ?: '',
    'ruta11_db_pass' => getenv('RUTA11_DB_PASS') ?: '',
    
    // OAuth Ruta11 App principal
    'ruta11_google_client_id' => getenv('RUTA11_GOOGLE_CLIENT_ID') ?: '',
    'ruta11_google_client_secret' => getenv('RUTA11_GOOGLE_CLIENT_SECRET') ?: '',
    'ruta11_google_redirect_uri' => getenv('RUTA11_GOOGLE_REDIRECT_URI') ?: 'https://laruta11.cl/api/auth/google/callback.php',
    
    // OAuth Ruta11 Jobs
    'ruta11_jobs_client_id' => getenv('RUTA11_JOBS_CLIENT_ID') ?: '',
    'ruta11_jobs_client_secret' => getenv('RUTA11_JOBS_CLIENT_SECRET') ?: '',
    'ruta11_jobs_redirect_uri' => getenv('RUTA11_JOBS_REDIRECT_URI') ?: 'https://laruta11.cl/api/auth/google/jobs_callback.php',
    
    // OAuth Ruta11 Tracker
    'ruta11_tracker_client_id' => getenv('RUTA11_TRACKER_CLIENT_ID') ?: '',
    'ruta11_tracker_client_secret' => getenv('RUTA11_TRACKER_CLIENT_SECRET') ?: '',
    'ruta11_tracker_redirect_uri' => getenv('RUTA11_TRACKER_REDIRECT_URI') ?: 'https://laruta11.cl/api/auth/google/tracker_callback.php',
    
    // Google Maps para Ruta11
    'ruta11_google_maps_api_key' => getenv('RUTA11_GOOGLE_MAPS_API_KEY') ?: '',
    
    // OAuth Gmail
    'gmail_client_id' => getenv('GMAIL_CLIENT_ID') ?: '',
    'gmail_client_secret' => getenv('GMAIL_CLIENT_SECRET') ?: '',
    'gmail_redirect_uri' => getenv('GMAIL_REDIRECT_URI') ?: 'https://laruta11.cl/api/auth/gmail/callback.php',
    'gmail_sender_email' => getenv('GMAIL_SENDER_EMAIL') ?: ''
];

// Test de configuración
if (isset($_GET['test'])) {
    header('Content-Type: application/json');
    
    $results = [];
    $results['aws_key'] = $config['aws_access_key_id'] ? 'OK' : 'MISSING';
    $results['aws_secret'] = $config['aws_secret_access_key'] ? 'OK' : 'MISSING';
    $results['aws_region'] = $config['aws_region'];
    $results['s3_bucket'] = $config['s3_bucket'];
    
    echo json_encode($results, JSON_PRETTY_PRINT);
    exit;
}

return $config;
?>