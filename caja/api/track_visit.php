<?php
header('Content-Type: application/json');
$url = 'https://websites-api-go-caja-r11.dj3bvg.easypanel.host/api/track/visit';
$data = file_get_contents('php://input');
$opts = ['http' => ['method' => 'POST', 'header' => 'Content-Type: application/json', 'content' => $data]];
echo file_get_contents($url, false, stream_context_create($opts));
