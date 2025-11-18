document.addEventListener('DOMContentLoaded', function() {
	const themeButton = document.getElementById("theme-button");
	loadTheme(themeButton);
	themeButton.addEventListener('click', () => {
		if (theme === "dark") {
			setTheme("light", themeButton)
			theme = "light"
		} else {
			setTheme("dark", themeButton)
			theme = "dark"
		}
	});
})
const html = document.documentElement;

let theme
const setTheme = (theme, btn) => {
   document.documentElement.removeAttribute('data-theme');

   if (theme === 'dark') {
	   html.setAttribute('data-theme', 'dark');
	   btn.innerHTML = "<img src='/static/aurinko.svg' class='light' />"
	} else {
		btn.innerHTML = "<img src='/static/kuu.svg' class='dark' />"
	}
	
	localStorage.setItem('theme', theme);
};

const loadTheme = (btn) => {
  const savedTheme = localStorage.getItem('theme');
  theme = savedTheme
  if (savedTheme) {
	  setTheme(savedTheme, btn);
	}
};

