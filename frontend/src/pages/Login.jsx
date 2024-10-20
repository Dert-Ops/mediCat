import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
	const [formData, setFormData] = useState({
		username: '',
		email: '',
		password: '',
		fullname: '', // Fullname added for registration
	});
	const [isLogin, setIsLogin] = useState(true);
	const navigate = useNavigate();

	const handleChange = (e) => {
		const { name, value } = e.target;
		setFormData({ ...formData, [name]: value });
	};

	function getCookie(name) {
		let cookieArr = document.cookie.split(";");
		for (let i = 0; i < cookieArr.length; i++) {
			let cookiePair = cookieArr[i].split("=");
			if (name === cookiePair[0].trim()) {
				return decodeURIComponent(cookiePair[1]);
			}
		}
		return null; // Eğer cookie bulunamazsa null döner
	}

	const handleLoginSubmit = async (e) => {
		e.preventDefault();

		const endpoint = 'http://45.9.30.65:8083/auth/signin';

		const requestData = {
			username: formData.username,
			password: formData.password,
		};

		try {
			const response = await fetch(endpoint, {
				method: 'POST', // POST isteği
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(requestData), // Kullanıcı adı ve şifreyi JSON formatında gönderiyoruz
			});

			if (response.ok) {
				console.log('Login işlemi başarılı');

				// Token'ı alıyoruz
				const { token } = await response.json(); // API'den token'ı alın
				document.cookie = `token=${token}; path=/`; // Cookie'yi ayarlayın

				// if (!token) {
				// 	return
				// }

				// Başarılı girişten sonra yönlendirme
				const userData = await fetch(`http://45.9.30.65:8083/users/${formData.username}`, {
					method: 'GET',
					headers: {
						'Authorization': `Bearer ${token}`,
						'Content-Type': 'application/json',
					},
				});

				if (userData.ok) {
					const userProfile = await userData.json();
					// Kullanıcı bilgilerini yönlendirdiğimiz User sayfasına geçirelim
					navigate('/user', { state: { user: userProfile } });
				} else {
					console.error('Kullanıcı bilgileri alınamadı:', userData.status);
				}
			} else {
				console.error('Login işlemi başarısız:', response.status);
			}
		} catch (error) {
			console.error('Bir hata oluştu:', error);
		}
	};


	// Handle signup submit (POST request)
	const handleSignupSubmit = async (e) => {
		e.preventDefault();

		// POST request for signup
		const endpoint = 'http://45.9.30.65:8083/auth/signup';

		try {
			const response = await fetch(endpoint, {
				mode: 'cors',
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({
					username: formData.username,
					email: formData.email,
					password: formData.password,
					fullname: formData.fullname,
				}),
			});

			if (response.ok) {
				const data = await response.json(); // JSON yanıtı al
				console.log('Register işlemi başarılı:', data);
			} else {
				const errorText = await response.text(); // Hata metnini al
				console.error('Register işlemi başarısız:', response.status, errorText);
			}
		} catch (error) {
			console.error('Bir hata oluştu:', error);
		}
	};

	const toggleForm = () => {
		setIsLogin(!isLogin);
	};

	return (
		<main className="absolute inset-0 top-16 flex items-center justify-center w-full h-screen bg-black p-4 bg-cover bg-center">
			<div className="w-full sm:w-full md:w-full max-w-[500px] bg-[#222222] p-8 rounded-xl z-10">
				<h2 className="text-2xl font-bold mb-6">{isLogin ? 'Login' : 'Register'}</h2>

				{/* Conditional form rendering based on isLogin */}
				<form onSubmit={isLogin ? handleLoginSubmit : handleSignupSubmit}>
					{!isLogin && (
						<>
							<div className="mb-4">
								<label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="fullname">
									Full Name
								</label>
								<input
									type="text"
									id="fullname"
									name="fullname"
									className="w-full px-3 py-2 border border-gray-300 rounded-lg"
									placeholder="Enter your full name"
									value={formData.fullname}
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
						</>
					)}
					<div className="mb-4">
						<label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="username">
							Username
						</label>
						<input
							type="text"
							id="username"
							name="username"
							className="w-full px-3 py-2 border border-gray-300 rounded-lg"
							placeholder="Enter your username"
							value={formData.username}
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
					{isLogin ? "Don't have an account? " : 'Already have an account? '}
					<span onClick={toggleForm} className="text-blue-500 cursor-pointer hover:underline">
						{isLogin ? 'Register here' : 'Login here'}
					</span>
				</p>
			</div>
		</main>
	);
};

export default Login;