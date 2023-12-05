import { renderPage } from "./routes/spa";

import "./styles/color.css";
import "./styles/style.css";
import "./styles/forms.css";
import { retrieveTheme } from "./themes/basics";

// The ONLY entry point
(function () {
  retrieveTheme();
  renderPage(location.pathname, false);

  window.addEventListener(
    "popstate",
    () => {
      renderPage(location.pathname, false);
    },
    false
  );

  // Making navigation bar work
  document.querySelectorAll<HTMLAnchorElement>("nav li a").forEach((link) => {
    if (link.href == "javascript:void(0)") {
      return;
    }

    link.onclick = (ev) => {
      ev.preventDefault();

      if (link.classList.contains("active")) {
        return;
      }

      let target = ev.target as HTMLAnchorElement;
      let url = new URL(target.href);

      renderPage(url.pathname, true);
      history.pushState({}, "", url.pathname);
    };
  });

  document
    .getElementById("special-dark-light-theme")
    ?.addEventListener("click", themeSwitchHandler);
})();

let icon = document.querySelector<HTMLSpanElement>(
  "#special-mode>.material-symbols-outlined"
)!;

function themeSwitchHandler(ev: MouseEvent) {
  ev.preventDefault();
  document.body.classList.toggle("dark_mode");

  if (document.body.classList.contains("dark_mode")) {
    localStorage.setItem("color_scheme", "dark");
  } else {
    localStorage.setItem("color_scheme", "light");
  }
  icon.textContent = localStorage.getItem("color_scheme") + "_mode";
}
