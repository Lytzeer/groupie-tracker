/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./static/**/*.html"],
  theme: {
    extend: {
      colors: {
        "mwg": "white power",
        //"mww": "red",
        "title":"red",
        "search-b":"#6d071a",
      },
    },
  },
  plugins: [],
}
