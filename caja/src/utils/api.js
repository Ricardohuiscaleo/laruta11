// API Helper - Usa API Go en producción
const API_BASE = import.meta.env.PUBLIC_API_BASE_URL || 'https://websites-api-go-caja-r11.dj3bvg.easypanel.host';

export const apiUrl = (endpoint) => {
  // Si el endpoint ya tiene el dominio completo, devolverlo tal cual
  if (endpoint.startsWith('http')) return endpoint;
  
  // Si empieza con /api/, usar API Go
  if (endpoint.startsWith('/api/')) {
    return `${API_BASE}${endpoint}`;
  }
  
  // Agregar /api/ si no lo tiene
  return `${API_BASE}/api/${endpoint.replace(/^\//, '')}`;
};

// Helper para fetch con API Go
export const apiFetch = async (endpoint, options = {}) => {
  const url = apiUrl(endpoint);
  const response = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
  });
  return response;
};
