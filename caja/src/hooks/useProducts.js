import { useState, useCallback, useMemo } from 'react';

/**
 * Hook para gestionar productos y categorías
 * Extrae lógica de MenuApp.jsx
 */
export function useProducts(initialProducts = []) {
  const [products, setProducts] = useState(initialProducts);
  const [activeCategory, setActiveCategory] = useState('hamburguesas');
  const [selectedProduct, setSelectedProduct] = useState(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [showInactiveProducts, setShowInactiveProducts] = useState(false);

  // Filtrar productos por categoría
  const productsByCategory = useMemo(() => {
    return products.filter(product => {
      const matchesCategory = product.category === activeCategory;
      const matchesActive = showInactiveProducts || product.active !== false;
      return matchesCategory && matchesActive;
    });
  }, [products, activeCategory, showInactiveProducts]);

  // Buscar productos
  const searchProducts = useMemo(() => {
    if (!searchQuery.trim()) return [];

    const query = searchQuery.toLowerCase();
    return products.filter(product =>
      product.name.toLowerCase().includes(query) ||
      product.description?.toLowerCase().includes(query)
    );
  }, [products, searchQuery]);

  // Cambiar categoría
  const changeCategory = useCallback((category) => {
    setActiveCategory(category);
    setSearchQuery(''); // Limpiar búsqueda al cambiar categoría
  }, []);

  // Seleccionar producto
  const selectProduct = useCallback((product) => {
    setSelectedProduct(product);
  }, []);

  // Cerrar detalle de producto
  const closeProductDetail = useCallback(() => {
    setSelectedProduct(null);
  }, []);

  // Actualizar producto
  const updateProduct = useCallback((productId, updates) => {
    setProducts(prev =>
      prev.map(product =>
        product.id === productId ? { ...product, ...updates } : product
      )
    );
  }, []);

  // Toggle like producto
  const toggleLike = useCallback(async (productId) => {
    try {
      const response = await fetch('/api/toggle_like.php', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ product_id: productId })
      });

      const data = await response.json();
      
      if (data.success) {
        updateProduct(productId, { likes: data.likes });
        return { success: true, likes: data.likes };
      }

      return { success: false };
    } catch (error) {
      console.error('Error toggling like:', error);
      return { success: false, error: error.message };
    }
  }, [updateProduct]);

  // Obtener producto por ID
  const getProductById = useCallback((productId) => {
    return products.find(p => p.id === productId);
  }, [products]);

  return {
    products,
    activeCategory,
    selectedProduct,
    searchQuery,
    showInactiveProducts,
    productsByCategory,
    searchProducts,
    changeCategory,
    selectProduct,
    closeProductDetail,
    updateProduct,
    toggleLike,
    getProductById,
    setSearchQuery,
    setShowInactiveProducts,
    setProducts
  };
}
