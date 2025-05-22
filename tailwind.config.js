/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./templates/**/*.html",  // Add leading ./ for relative path
    "./templates/**/*.{html,js}",  // Optional: include JS files if you use them
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}