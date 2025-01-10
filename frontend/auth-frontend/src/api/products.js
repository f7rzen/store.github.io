// api/products.js

import axios from 'axios';
import config from '../config'; // Импортируем конфигурацию

export const getProducts = async (token) => {
    try {
        const response = await axios.get(`${config.apiBaseUrl}/admin/products`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });
        return response.data.data; // Данные о товарах
    } catch (error) {
        console.error("Ошибка при получении товаров:", error);
        throw error;
    }
};
