const themeButton = document.getElementById("theme-button");
const html = document.documentElement;

let theme
const setTheme = (theme) => {
   document.documentElement.removeAttribute('data-theme');

  if (theme === 'dark') {
    html.setAttribute('data-theme', 'dark');
  }

  localStorage.setItem('theme', theme);
};

const loadTheme = () => {
  const savedTheme = localStorage.getItem('theme');
  theme = savedTheme
  if (savedTheme) {
    setTheme(savedTheme);
  }
};

themeButton.addEventListener('click', () => {
    if (theme === "dark") {
        setTheme("light")
        theme = "light"
    } else {
        setTheme("dark")
        theme = "dark"
    }
});

loadTheme();