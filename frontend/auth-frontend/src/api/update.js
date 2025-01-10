import axios from "axios";
import config from '../config'; // Импортируем конфигурацию

// Функция для обновления товара по ID
export const updateProductById = async (id, updatedValues, token) => {
    try {
        const response = await axios.put(
            `${config.apiBaseUrl}/admin/products/${id}`, // URL для PUT запроса с ID товара
            {
                product: updatedValues, // Обновляемые данные товара
            },
            {
                headers: {
                    Authorization: `Bearer ${token}`, // Заголовок с Bearer токеном
                    "Content-Type": "application/json", // Устанавливаем тип контента
                },
            }
        );
        return response.data; // Возвращаем ответ (обновленные данные товара)
    } catch (error) {
        if (error.response) {
            // Обработка ошибок от сервера (например, 400, 404, 500)
            throw new Error(error.response.data.message || "Ошибка обновления товара");
        } else {
            // Обработка других ошибок (например, ошибки сети)
            throw new Error("Ошибка соединения с сервером");
        }
    }
};
