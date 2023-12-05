import { renderPage } from "./routes/spa";

import "./styles/color.css";
import "./styles/style.css";
import "./styles/forms.css";

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

function retrieveTheme() {
  var theme = localStorage.getItem("color_scheme");
  if (theme != null) {
    document.body.classList.remove("light_mode", "dark_mode");
    let mode = theme + "_mode";
    document.body.classList.add(mode);
    icon.textContent = mode;
  }
}

// The ONLY entry point
(function () {
  retrieveTheme();
  renderPage(location.pathname, null);

  window.addEventListener(
    "popstate",
    () => {
      renderPage(location.pathname, null);
    },
    false
  );

  document.querySelectorAll<HTMLAnchorElement>("nav li a").forEach((link) => {
    if (link.href == "javascript:void(0)") {
      return;
    }

    link.onclick = (ev) => {
      ev.preventDefault();

      let target = ev.target as HTMLAnchorElement;
      let url = new URL(target.href);

      renderPage(url.pathname, location.pathname);
      history.pushState({}, "", url.pathname);
    };
  });

  document
    .getElementById("special-mode")
    ?.addEventListener("click", themeSwitchHandler);
})();
