import { useState, useCallback } from 'react';

/**
 * Hook para gestionar el carrito de compras
 * Extrae lógica de MenuApp.jsx para reducir complejidad
 */
export function useCart() {
  const [cart, setCart] = useState([]);
  const [isCartOpen, setIsCartOpen] = useState(false);

  // Agregar producto al carrito
  const addToCart = useCallback((product, quantity = 1, extras = []) => {
    setCart(prevCart => {
      const existingItem = prevCart.find(
        item => item.id === product.id && 
        JSON.stringify(item.extras) === JSON.stringify(extras)
      );

      if (existingItem) {
        return prevCart.map(item =>
          item.id === product.id && 
          JSON.stringify(item.extras) === JSON.stringify(extras)
            ? { ...item, quantity: item.quantity + quantity }
            : item
        );
      }

      return [...prevCart, { 
        ...product, 
        quantity, 
        extras,
        cartId: Date.now() + Math.random()
      }];
    });
  }, []);

  // Remover producto del carrito
  const removeFromCart = useCallback((cartId) => {
    setCart(prevCart => prevCart.filter(item => item.cartId !== cartId));
  }, []);

  // Actualizar cantidad
  const updateQuantity = useCallback((cartId, quantity) => {
    if (quantity <= 0) {
      removeFromCart(cartId);
      return;
    }

    setCart(prevCart =>
      prevCart.map(item =>
        item.cartId === cartId ? { ...item, quantity } : item
      )
    );
  }, [removeFromCart]);

  // Limpiar carrito
  const clearCart = useCallback(() => {
    setCart([]);
  }, []);

  // Calcular total
  const getTotal = useCallback(() => {
    return cart.reduce((total, item) => {
      const itemPrice = parseFloat(item.price) || 0;
      const extrasPrice = (item.extras || []).reduce(
        (sum, extra) => sum + (parseFloat(extra.price) || 0),
        0
      );
      return total + (itemPrice + extrasPrice) * item.quantity;
    }, 0);
  }, [cart]);

  // Calcular cantidad total de items
  const getTotalItems = useCallback(() => {
    return cart.reduce((total, item) => total + item.quantity, 0);
  }, [cart]);

  // Toggle carrito
  const toggleCart = useCallback(() => {
    setIsCartOpen(prev => !prev);
  }, []);

  return {
    cart,
    isCartOpen,
    addToCart,
    removeFromCart,
    updateQuantity,
    clearCart,
    getTotal,
    getTotalItems,
    toggleCart,
    setIsCartOpen
  };
}
