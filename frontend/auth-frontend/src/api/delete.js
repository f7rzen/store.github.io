import config from '../config'; // Импортируем конфигурацию
import axios from "axios";

export const deleteProductById = async (id, token) => {
    try {
        const response = await axios.delete(`${config.apiBaseUrl}/admin/products/${id}`, {
            headers: {
                Authorization: `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        });
        return response.data; // Вернем данные успешного ответа
    } catch (error) {
        console.error('Ошибка при удалении товара:', error);
        throw error.response ? error.response.data : error; // Обработаем ошибку
    }
};