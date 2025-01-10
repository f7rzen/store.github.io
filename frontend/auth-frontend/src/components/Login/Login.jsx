import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import './login.css';

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault(); // Предотвращаем перезагрузку страницы

        try {
            // Отправляем запрос на сервер
            const response = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ email, password }), // Отправляем логин и пароль
                credentials: "include", // передача куков
            });

            if (response.ok) {
                const data = await response.json();

                // Сохраняем токен в localStorage
                localStorage.setItem("accessToken", data.accessToken);

                // Перенаправляем пользователя на страницу admin
                navigate("/admin");
            } else {
                setError("Ошибка входа: неверный логин или пароль.");
            }
        } catch (error) {
            console.error("Ошибка:", error);
            setError("Произошла ошибка при подключении к серверу.");
        }
    };

    return (
        <div className="login-container">
            <h1>Login</h1>
            <form onSubmit={handleLogin}>
                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">Login</button>
            </form>
            {error && <p>{error}</p>}
        </div>
    );
};

export default Login;
