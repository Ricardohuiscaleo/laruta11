import { useState, useCallback } from 'react';

/**
 * Hook para gestionar el proceso de checkout
 * Extrae lógica de MenuApp.jsx
 */
export function useCheckout() {
  const [showCheckout, setShowCheckout] = useState(false);
  const [showPayment, setShowPayment] = useState(false);
  const [currentOrder, setCurrentOrder] = useState(null);
  const [isProcessing, setIsProcessing] = useState(false);
  
  const [customerInfo, setCustomerInfo] = useState({
    name: '',
    phone: '',
    email: '',
    address: '',
    deliveryType: 'pickup',
    pickupTime: '',
    customerNotes: '',
    deliveryDiscount: false,
    pickupDiscount: false,
    birthdayDiscount: false
  });

  // Iniciar checkout
  const startCheckout = useCallback((user) => {
    if (!user) {
      return { success: false, error: 'Usuario no autenticado' };
    }
    setShowCheckout(true);
    return { success: true };
  }, []);

  // Actualizar info del cliente
  const updateCustomerInfo = useCallback((field, value) => {
    setCustomerInfo(prev => ({
      ...prev,
      [field]: value
    }));
  }, []);

  // Validar formulario
  const validateForm = useCallback(() => {
    const { name, phone, deliveryType, address } = customerInfo;
    
    if (!name || !phone) {
      return { valid: false, error: 'Nombre y teléfono son requeridos' };
    }

    if (deliveryType === 'delivery' && !address) {
      return { valid: false, error: 'Dirección es requerida para delivery' };
    }

    return { valid: true };
  }, [customerInfo]);

  // Procesar orden
  const processOrder = useCallback(async (cart, paymentMethod) => {
    setIsProcessing(true);
    
    try {
      const validation = validateForm();
      if (!validation.valid) {
        throw new Error(validation.error);
      }

      const orderData = {
        customer: customerInfo,
        items: cart,
        paymentMethod,
        timestamp: new Date().toISOString()
      };

      // Aquí iría la llamada a la API
      const response = await fetch('/api/create_order.php', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(orderData)
      });

      const data = await response.json();

      if (data.success) {
        setCurrentOrder(data.order);
        setShowPayment(true);
        return { success: true, order: data.order };
      }

      throw new Error(data.error || 'Error al crear orden');
    } catch (error) {
      return { success: false, error: error.message };
    } finally {
      setIsProcessing(false);
    }
  }, [customerInfo, validateForm]);

  // Completar orden
  const completeOrder = useCallback(() => {
    setShowCheckout(false);
    setShowPayment(false);
    setCurrentOrder(null);
    setCustomerInfo({
      name: '',
      phone: '',
      email: '',
      address: '',
      deliveryType: 'pickup',
      pickupTime: '',
      customerNotes: '',
      deliveryDiscount: false,
      pickupDiscount: false,
      birthdayDiscount: false
    });
  }, []);

  // Cancelar checkout
  const cancelCheckout = useCallback(() => {
    setShowCheckout(false);
    setShowPayment(false);
  }, []);

  return {
    showCheckout,
    showPayment,
    currentOrder,
    customerInfo,
    isProcessing,
    startCheckout,
    updateCustomerInfo,
    validateForm,
    processOrder,
    completeOrder,
    cancelCheckout,
    setShowCheckout,
    setShowPayment
  };
}
