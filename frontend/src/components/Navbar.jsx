import React, { useState, useEffect } from 'react';
import { Link, useLocation } from 'react-router-dom';

const Navbar = () => {
    const [isOpen, setIsOpen] = useState(false);
    const location = useLocation();

    const toggleMenu = () => {
        setIsOpen(!isOpen);
    };

    useEffect(() => {
        setIsOpen(false);
    }, [location]);

    return (
        <nav className="fixed top-0 left-0 w-full bg-nav-color h-16 p-4 z-10">
            <div className="relative flex items-center justify-center w-full h-full">
                <Link to="/login" className="text-white focus:outline-none bg-inherit absolute left-4">
                    <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                        <circle cx="12" cy="8" r="4" stroke="currentColor" strokeWidth="2" />
                        <path d="M5 20c0-3.313 4.687-6 8-6s8 2.687 8 6" stroke="currentColor" strokeWidth="2" />
                    </svg>
                </Link>

                <div className="absolute flex items-center">
                    <Link to="/" className="text-white text-2xl font-bold flex items-center">
                        <img src="/images/favicon.ico" alt="MediCat" className="h-8 mr-2" />
                        <h2>MediCat</h2>
                    </Link>
                </div>

                <button onClick={toggleMenu} className="text-white focus:outline-none bg-inherit absolute right-4">
                    <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16"></path>
                    </svg>
                </button>
            </div>

            <div className={`absolute inset-x-0 top-16 bg-nav-toggle-color text-white space-y-3 p-4 overflow-hidden transition-all duration-500 ease-in-out ${isOpen ? 'max-h-96 opacity-100 visible' : 'max-h-0 opacity-0 invisible'}`}>
                <a href="#" className="block hover:bg-nav-color p-2">Home</a>
                <a href="#" className="block hover:bg-nav-color p-2">About</a>
                <a href="#" className="block hover:bg-nav-color p-2">Services</a>
                <a href="#" className="block hover:bg-nav-color p-2">Contact</a>
            </div>
        </nav>
    );
};

export default Navbar;
