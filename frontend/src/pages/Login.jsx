import React, { useState } from 'react';

const Login = () => {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
    });
    const [isLogin, setIsLogin] = useState(true);

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
                console.log('Form successfully submitted');
            } else {
                console.error('Form submission failed:', response.status);
            }
        } catch (error) {
            console.error('An error occurred:', error);
        }
    };

    const toggleForm = () => {
        setIsLogin(!isLogin);
    };

    const backgroundImageUrl = isLogin
        ? "url('/images/gofret.png')"
        : "url('/images/necmi.JPG')";

    return (
        <main className="absolute inset-0 top-16 flex items-center justify-center w-full h-screen bg-black p-4 bg-cover bg-center">
            <div
                className="absolute inset-0 bg-cover bg-center opacity-95 z-0 transition-all duration-500"
                style={{ backgroundImage: backgroundImageUrl }}
            ></div>
            <div className={`w-full sm:w-full md:w-full max-w-[500px] opacity-95 bg-[#222222] p-8 rounded-xl z-10 transform transition-transform duration-500 ${isLogin ? 'rotate-y-180' : ''}`}>
                <h2 className="text-2xl font-bold mb-6">{isLogin ? 'Login' : 'Register'}</h2>
                <form onSubmit={handleSubmit}>
                    {!isLogin && (
                        <div className="mb-4">
                            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="name">
                                Name
                            </label>
                            <input
                                type="text"
                                id="name"
                                name="username"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                placeholder="Enter your name"
                                value={formData.username}
                                onChange={handleChange}
                            />
                        </div>
                    )}
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
                    <button type="submit" className="w-full bg-nav-color text-white py-2 px-4 rounded-xl hover:bg-nav-toggle-color">
                        {isLogin ? 'Login' : 'Register'}
                    </button>
                </form>
                <p className="mt-4 text-gray-500 text-sm text-center">
                    {isLogin ? 'Don\'t have an account? ' : 'Already have an account? '}
                    <span onClick={toggleForm} className="text-blue-500 cursor-pointer hover:underline">
                        {isLogin ? 'Register here' : 'Login here'}
                    </span>
                </p>
            </div>
        </main>
    );
};

export default Login;
