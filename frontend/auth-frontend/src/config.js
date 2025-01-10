// src/config.js

const config = {
    apiBaseUrl: process.env.REACT_APP_API_BASE_URL || 'http://localhost:8080', // Использует переменную окружения, если она установлена, или localhost по умолчанию
};

export default config;
