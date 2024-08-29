import React, { useState } from 'react';

const Login = () => {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        try {
            const response = await fetch('http://your-backend-url/endpoint', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            });

            if (response.ok) {
                // Başarılı durum
                console.log('Form successfully submitted');
                // Başarılı bir yanıt aldıktan sonra işlemler yapabilirsiniz
            } else {
                // Hata durumunda
                console.error('Form submission failed:', response.status);
                // Hata mesajı veya işlem yapabilirsiniz
            }
        } catch (error) {
            console.error('An error occurred:', error);
            // Hata işleme
        }
    };

    return (
        <main className="absolute inset-0 top-16 flex items-center justify-center w-full h-screen bg-black p-4">
            <div className="w-full max-w-[500px] bg-[#222222] p-8 rounded-xl">
                <h2 className="text-2xl font-bold mb-6">Register</h2>
                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="name">
                            Name
                        </label>
                        <input
                            type="text"
                            id="name"
                            name="username"  // formData'da "username" olarak tanımlı
                            className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                            placeholder="Enter your name"
                            value={formData.username}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
                            Email
                        </label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                            placeholder="Enter your email"
                            value={formData.email}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                            placeholder="Enter your password"
                            value={formData.password}
                            onChange={handleChange}
                        />
                    </div>
                    <button
                        type="submit"
                        className="w-full bg-nav-color text-white py-2 px-4 rounded-xl hover:bg-nav-toggle-color"
                    >
                        Register
                    </button>
                </form>
            </div>
        </main>
    );
};

export default Login;
