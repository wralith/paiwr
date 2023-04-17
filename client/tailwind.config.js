/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  darkMode: "class",
  theme: {
    extend: {
      animation: {
        "bounce-faster": "bounce 0.3s ease-out infinite",
      },
    },
  },
  plugins: [],
}
