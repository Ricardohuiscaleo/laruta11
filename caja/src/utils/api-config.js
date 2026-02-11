// Feature Flags para migración PHP → Go
export const API_FLAGS = {
  USE_GO_AUTH: false,
  USE_GO_COMPRAS: false,
  USE_GO_INVENTORY: false,
  USE_GO_QUALITY: false,
  USE_GO_CATALOG: false,
  USE_GO_ORDERS: false
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
