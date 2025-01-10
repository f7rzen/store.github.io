import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./components/Login/Login";
import Admin from "./components/Admin/Admin";

const App = () => {
    return (
        <Router>
            <div>
                <Routes>
                    {/* Перенаправление с главной страницы на /admin */}
                    <Route path="/" element={<Navigate to="/admin" />} />

                    {/* Страница входа */}
                    <Route path="/login" element={<Login />} />

                    {/* Страница администратора */}
                    <Route path="/admin" element={<Admin />} />

                    {/* Обработка неизвестных маршрутов */}
                    <Route path="*" element={<h1>404: Страница не найдена</h1>} />
                </Routes>
            </div>
        </Router>
    );
};

export default App;
