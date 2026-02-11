<?php
header('Content-Type: application/json');
$url = 'https://websites-api-go-caja-r11.dj3bvg.easypanel.host/api/auth/check_session.php';
echo file_get_contents($url);
