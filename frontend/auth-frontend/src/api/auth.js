import { jwtDecode } from 'jwt-decode';
import config from '../config'; // Импортируем конфигурацию

export const isTokenValid = (token) => {
    if (!token) {
        console.warn("Токен отсутствует");
        return false;
    }

    try {
        const decoded = jwtDecode(token);

        // Проверяем, истёк ли токен
        if (decoded.exp * 1000 < Date.now()) {
            console.warn("Токен истёк");
            return false;
        }

        return true; // Токен валиден
    } catch (error) {
        console.error("Ошибка при декодировании токена:", error);
        return false; // Невалидный токен
    }
};


export const tryRefreshAccessToken = async (navigate) => {
    try {
        const response = await fetch(`${config.apiBaseUrl}/refresh`, {
            method: "POST",
            credentials: "include", // Отправляем куки с refresh токеном
        });

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem("accessToken", data.accessToken);
            console.log("Сохранён новый accessToken:", localStorage.getItem("accessToken"));
            return true;
        } else {
            throw new Error("Не удалось обновить токен");
        }
    } catch (error) {
        console.error("Ошибка обновления токена:", error);
        navigate("/login"); // Перенаправляем на страницу логина
        return false;
    }
};

export const checkAuth = async (navigate) => {
    const currentAccessToken = localStorage.getItem("accessToken");
    console.log("Текущий токен перед проверкой:", currentAccessToken);
    // Проверяем валидность access токена
    if (!isTokenValid(currentAccessToken)) {
        console.warn("Access токен отсутствует или истёк. Пытаемся обновить...");
        return await tryRefreshAccessToken(navigate);
    }

    console.log("Access токен валиден");
    return true;
};
