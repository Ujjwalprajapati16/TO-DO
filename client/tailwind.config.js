/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        primary: "#964d4d", // Adding the color palette to the theme
        secondary: "#642b2b",
        tertiary: "#591f1f",
        quaternary: "#500000",
        quinary: "#7b5151",
      },
    },
  },
  plugins: [require("daisyui")],
};
