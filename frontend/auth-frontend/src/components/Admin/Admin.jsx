import React, { useEffect, useState, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { checkAuth } from "../../api/auth";
import { getProducts } from "../../api/products";
import { uploadFile } from "../../api/upload";
import { deleteProductById } from "../../api/delete";
import "./admin.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFileImport, faUpload, faTrash } from "@fortawesome/free-solid-svg-icons";

const Admin = () => {
    const navigate = useNavigate();
    const [products, setProducts] = useState([]);
    const [selectedProducts, setSelectedProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [selectedFile, setSelectedFile] = useState(null);
    const [uploadStatus, setUploadStatus] = useState("");
    const fileInputRef = useRef(null);

    useEffect(() => {
        const authenticate = async () => {
            const isAuthenticated = await checkAuth(navigate);
            if (!isAuthenticated) {
                navigate("/login");
            }
        };

        authenticate();

        const interval = setInterval(() => {
            authenticate();
        }, 4000);

        return () => clearInterval(interval);
    }, [navigate]);

    // Функция для получения продуктов
    const fetchProducts = async () => {
        try {
            const token = localStorage.getItem("accessToken");
            const data = await getProducts(token);
            setProducts(data);
            setLoading(false);
        } catch (err) {
            setError("Не удалось загрузить товары");
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchProducts();
    }, []);

    const handleFileChange = (event) => {
        const file = event.target.files[0];
        if (file) {
            setSelectedFile(file);
        }
    };

    const handleUpload = async () => {
        if (!selectedFile) {
            setUploadStatus("Пожалуйста, выберите файл.");
            return;
        }

        try {
            const token = localStorage.getItem("accessToken");
            const statusMessage = await uploadFile(selectedFile, token);
            setUploadStatus(statusMessage);
            setSelectedFile(null);
            fetchProducts(); // Повторно загружаем данные после удаления
        } catch (error) {
            setUploadStatus(error.message || "Ошибка загрузки файла");
        }
    };

    const handleSelectProduct = (id) => {
        setSelectedProducts((prevSelected) =>
            prevSelected.includes(id)
                ? prevSelected.filter((productId) => productId !== id)
                : [...prevSelected, id]
        );
    };

    const handleDeleteSelected = async () => {
        try {
            const token = localStorage.getItem("accessToken");
            for (const id of selectedProducts) {
                await deleteProductById(id, token);
            }
            setSelectedProducts([]); // Очистка выбранных товаров
            fetchProducts(); // Повторно загружаем данные после удаления
        } catch (error) {
            alert("Ошибка при удалении товаров");
        }
    };

    if (loading) {
        return <div>Загрузка...</div>;
    }

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div>
            <div className="header">
                <h1>Admin Page</h1>
            </div>

            <div className="admin-container">
                <table>
                    <thead>
                    <tr>
                        <th>
                            <input
                                type="checkbox"
                                onChange={(e) => {
                                    if (e.target.checked) {
                                        setSelectedProducts(products.map((p) => p.ID));
                                    } else {
                                        setSelectedProducts([]);
                                    }
                                }}
                                checked={
                                    selectedProducts.length === products.length &&
                                    products.length > 0
                                }
                            />
                        </th>
                        <th>ID</th>
                        <th>Название</th>
                        <th>Описание</th>
                        <th>Цена</th>
                        <th>Ссылки на изображения</th>
                        <th>Категория</th>
                    </tr>
                    </thead>
                    <tbody>
                    {products.map((product) => (
                        <tr key={product.ID}>
                            <td>
                                <input
                                    type="checkbox"
                                    checked={selectedProducts.includes(product.ID)}
                                    onChange={() => handleSelectProduct(product.ID)}
                                />
                            </td>
                            <td>{product.ID}</td>
                            <td>{product.name}</td>
                            <td>{product.description}</td>
                            <td>{product.price}</td>
                            <td>{product.image_url}</td>
                            <td>{product.category_id}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>

                <div className="admin-container">
                    <div className="import-button-container">
                        <button
                            className={`import-button ${selectedProducts.length > 0 ? "delete-button" : ""}`}
                            onClick={() => {
                                if (selectedProducts.length > 0) {
                                    handleDeleteSelected();
                                } else if (selectedFile) {
                                    handleUpload();
                                } else {
                                    fileInputRef.current.click();
                                }
                            }}
                        >
                            <FontAwesomeIcon
                                icon={
                                    selectedProducts.length > 0
                                        ? faTrash
                                        : selectedFile
                                            ? faUpload
                                            : faFileImport
                                }
                                style={{ marginRight: "8px" }}
                            />
                            {selectedProducts.length > 0 ? "DELETE" : selectedFile ? "UPLOAD" : "IMPORT"}
                        </button>
                    </div>
                    <input
                        type="file"
                        ref={fileInputRef}
                        style={{ display: "none" }}
                        onChange={handleFileChange}
                    />
                    {uploadStatus && <p>{uploadStatus}</p>}
                </div>
            </div>
        </div>
    );
};

export default Admin;
