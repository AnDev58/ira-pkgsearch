export const themeIsDark = () => localStorage.getItem("color_scheme") == "dark";

export const pageIsDark = () => document.body.classList.contains("dark_mode");

export const themeSwitcher = document.querySelector<HTMLSpanElement>(
  "#special-dark-light-theme>.material-symbols-outlined"
)!;

// UPDATERS
export const updateTheme = () =>
  localStorage.setItem("color_scheme", pageIsDark() ? "dark" : "light");

export const updateIcon = () =>
  (themeSwitcher.textContent = themeIsDark() ? "dark_mode" : "light_mode");

export const restorePageTheme = () => {
  if (themeIsDark()) {
    document.body.classList.add("dark_mode");
  } else {
    document.body.classList.remove("dark_mode");
  }
};

export function switchTheme() {
  document.body.classList.toggle("dark_mode");
  updateTheme();
  updateIcon();
}

export function retrieveTheme() {
  restorePageTheme();
  updateIcon();
}
