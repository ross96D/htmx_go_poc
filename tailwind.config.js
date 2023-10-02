/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pages/**/*.html", "./pages/*.html"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark"],
  },
};
