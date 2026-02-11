// Feature Flags para migración PHP → Go
export const API_FLAGS = {
  USE_GO_AUTH: true,
  USE_GO_COMPRAS: true,
  USE_GO_INVENTORY: true,
  USE_GO_QUALITY: true,
  USE_GO_CATALOG: true,
  USE_GO_ORDERS: true
};

const GO_API = 'https://websites-api-go-caja-r11.dj3bvg.easypanel.host';
const PHP_API = '';

export const api = (module, endpoint) => {
  const useGo = API_FLAGS[`USE_GO_${module.toUpperCase()}`];
  return `${useGo ? GO_API : PHP_API}${endpoint}`;
};

// Shortcuts
export const authApi = (e) => api('auth', e);
export const comprasApi = (e) => api('compras', e);
export const inventoryApi = (e) => api('inventory', e);
export const qualityApi = (e) => api('quality', e);
export const catalogApi = (e) => api('catalog', e);
export const ordersApi = (e) => api('orders', e);
