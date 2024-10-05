import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './App.css'
import Home from './pages/Home.jsx'
import Login from './pages/Login.jsx'
import Navbar from './components/Navbar.jsx'

function App() {
    return (
        <>
            <BrowserRouter>
                <div>
                    <Navbar />
                    <Routes>
                        <Route index element={<Home />} />
                        <Route path="home" element={<Home />} />
                        <Route path="/" element={<Home />} />
                        <Route path="login" element={<Login />} />
                    </Routes>
                </div>
            </BrowserRouter>
        </>
    )
}

export default App
