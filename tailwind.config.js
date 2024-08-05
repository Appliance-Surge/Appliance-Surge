/** @type {import('tailwindcss').Config} */
export default {
  content: ["./templates/**/*.html"],
  theme: {
    extend: {
        fontFamily: {
            intertight: ['interTight', 'sans-serif'],
            inter: ['inter', 'sans-serif'],
        },
        colors: {
            'brand-gray-200': '#858589',
            'brand-gray-300': '#202024',
            'brand-gray-800': '#18181c',
            'brand-gray-900': '#10131E',
            'brand-blue-500': '#26bbff',
            'brand-green': '#123334',
        },
    },
  },
  plugins: [],
}

