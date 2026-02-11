// API Go en producción (funciona desde local)
const API_BASE = 'https://websites-api-go-caja-r11.dj3bvg.easypanel.host';

export const API = {
  dashboard: `${API_BASE}/api/dashboard`,
  products: `${API_BASE}/api/products`,
  orders: `${API_BASE}/api/orders`,
  auth: `${API_BASE}/api/auth`,
  compras: `${API_BASE}/api/compras`,
  ingredientes: `${API_BASE}/api/ingredientes`,
};
