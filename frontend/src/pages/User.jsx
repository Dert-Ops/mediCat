import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';

const User = () => {
    const [userData, setUserData] = useState(null);
    const [isEditing, setIsEditing] = useState(false);
    const [formData, setFormData] = useState({
        profile_picture: '',
        age: '',
        bio: '',
        github_account: '',
        linkedin_account: '',
        google_account: '',
        job: '',
        fav_email: '',
        location: ''
    });

    const navigate = useNavigate();
    const location = useLocation();
    const user = location.state?.user;

    useEffect(() => {
        const fetchUserData = async () => {
            const token = getCookie("token");
            if (!token) {
                navigate('/login');
                return;
            }

            try {
                const response = await fetch(`http://45.9.30.65:8083/users/${user.username}`, {
                    method: 'GET',
                    headers: {
                        'Authorization': `Bearer ${token}`,
                        'Content-Type': 'application/json',
                    },
                });

                if (response.ok) {
                    const data = await response.json();
                    setUserData(data);
                    setFormData({
                        profile_picture: data.profile_picture,
                        age: data.age,
                        bio: data.bio,
                        github_account: data.github_account,
                        linkedin_account: data.linkedin_account,
                        google_account: data.google_account,
                        job: data.job,
                        fav_email: data.fav_email,
                        location: data.location
                    });
                } else {
                    console.error('Kullanıcı bilgileri alınamadı:', response.status);
                }
            } catch (error) {
                console.error('Bir hata oluştu:', error);
            }
        };

        fetchUserData();
    }, [navigate]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const token = getCookie("token");

        try {
            const response = await fetch('http://45.9.30.65:8083/users/update', {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            });

            if (response.ok) {
                const updatedData = await response.json();
                setUserData(updatedData);
                setIsEditing(false);
            } else {
                console.error('Güncelleme başarısız:', response.status);
            }
        } catch (error) {
            console.error('Bir hata oluştu:', error);
        }
    };

    const toggleEdit = () => {
        setIsEditing(!isEditing);
    };

    const getCookie = (name) => {
        let cookieArr = document.cookie.split(";");
        for (let i = 0; i < cookieArr.length; i++) {
            let cookiePair = cookieArr[i].split("=");
            if (name === cookiePair[0].trim()) {
                return decodeURIComponent(cookiePair[1]);
            }
        }
        return null;
    };

    if (!userData) {
        return <div>Loading...</div>;
    }

    return (
        <main className="absolute inset-0 top-16 flex items-center justify-center w-full h-screen bg-black p-4 bg-cover bg-center">
            <div className="w-full sm:w-full md:w-full max-w-[500px] bg-[#222222] p-8 rounded-xl z-10">
                <h2 className="text-2xl font-bold mb-6">Kullanıcı Profili</h2>

                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="profile_picture">
                            Profile Picture
                        </label>
                        {isEditing ? (
                            <input
                                type="text"
                                id="profile_picture"
                                name="profile_picture"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.profile_picture}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.profile_picture}</p>
                        )}
                    </div>

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="age">
                            Age
                        </label>
                        {isEditing ? (
                            <input
                                type="number"
                                id="age"
                                name="age"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.age}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.age}</p>
                        )}
                    </div>

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="bio">
                            Bio
                        </label>
                        {isEditing ? (
                            <textarea
                                id="bio"
                                name="bio"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.bio}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.bio}</p>
                        )}
                    </div>

                    {/* Diğer alanlar için benzer yapıyı kullanıyoruz */}
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="github_account">
                            Github Account
                        </label>
                        {isEditing ? (
                            <input
                                type="text"
                                id="github_account"
                                name="github_account"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.github_account}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.github_account}</p>
                        )}
                    </div>

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="linkedin_account">
                            Linkedin Account
                        </label>
                        {isEditing ? (
                            <input
                                type="text"
                                id="linkedin_account"
                                name="linkedin_account"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.linkedin_account}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.linkedin_account}</p>
                        )}
                    </div>

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="google_account">
                            Google Account
                        </label>
                        {isEditing ? (
                            <input
                                type="text"
                                id="google_account"
                                name="google_account"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.google_account}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.google_account}</p>
                        )}
                    </div>

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="job">
                            Job
                        </label>
                        {isEditing ? (
                            <input
                                type="text"
                                id="job"
                                name="job"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.job}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.job}</p>
                        )}
                    </div>

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="fav_email">
                            Favorite Email
                        </label>
                        {isEditing ? (
                            <input
                                type="email"
                                id="fav_email"
                                name="fav_email"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.fav_email}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.fav_email}</p>
                        )}
                    </div>

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="location">
                            Location
                        </label>
                        {isEditing ? (
                            <input
                                type="text"
                                id="location"
                                name="location"
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg"
                                value={formData.location}
                                onChange={handleChange}
                            />
                        ) : (
                            <p>{userData.location}</p>
                        )}
                    </div>

                    <button
                        type="submit"
                        className={`w-full bg-nav-color text-white px-4 py-2 rounded-lg ${isEditing ? '' : 'hidden'}`}
                    >
                        Save
                    </button>
                </form>

                <button
                    onClick={toggleEdit}
                    className="w-full bg-blue-500 text-white px-4 py-2 mt-4 rounded-lg"
                >
                    {isEditing ? 'Cancel' : 'Edit'}
                </button>
            </div>
        </main>
    );
};

export default User;
