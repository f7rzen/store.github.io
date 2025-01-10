// src/api/upload.js
import config from '../config'; // Импортируем конфигурацию

export const uploadFile = async (file, token) => {
    const formData = new FormData();
    formData.append("file", file);

    try {
        const response = await fetch(`${config.apiBaseUrl}/admin/upload`, {
            method: "POST",
            headers: {
                Authorization: `Bearer ${token}`, // Добавляем токен в заголовок
            },
            body: formData,
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || "Ошибка загрузки файла");
        }
    } catch (error) {
        console.error("Ошибка при загрузке файла:", error);
        throw error;
    }
};
