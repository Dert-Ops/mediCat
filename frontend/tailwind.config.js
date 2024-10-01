/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        // colors: {
        //     'blue': '#1fb6ff',
        //     'purple': '#7e5bef',
        //     'pink': '#ff49db',
        //     'orange': '#ff7849',
        //     'green': '#13ce66',
        //     'yellow': '#ffc82c',
        //     'gray-dark': '#273444',
        //     'gray': '#8492a6',
        //     'gray-light': '#d3dce6',
        //     'navbar-purple': '#4c1d95',
        // },
        extend: {
            keyframes: {
                'fade-out': {
                    '0%': { opacity: '1' },
                    '100%': { opacity: '0' },
                },
                'fade-in': {
                    '0%': { opacity: '0' },
                    '100%': { opacity: '1' },
                },
            },
            animation: {
                'fade-out': 'fade-out 0.5s ease-out forwards',
                'fade-in': 'fade-in 0.5s ease-in forwards',
            },
            colors: {
                'nav-color': 'rgb(42, 3, 54)',
                'nav-toggle-color': 'rgb(50, 3, 64)',
            },
        },
    },
    plugins: [],
}

