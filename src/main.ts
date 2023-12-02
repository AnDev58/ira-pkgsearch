import indexPage from "./pages/index.htm?raw";
import "./styles/color.css";
import "./styles/style.css";

let icon = document.querySelector<HTMLSpanElement>(
  "#special-mode>.material-symbols-outlined"
)!;

function retrieve_theme() {
  var theme = localStorage.getItem("color_scheme");
  if (theme != null) {
    document.body.classList.remove("light_mode", "dark_mode");
    let mode = theme + "_mode";
    document.body.classList.add(mode);
    icon.textContent = mode;
  }
}

(function () {
  let data = new DOMParser().parseFromString(indexPage, "text/html");
  document.getElementById("app")?.appendChild(data.documentElement);

  retrieve_theme();

  window.addEventListener(
    "storage",
    function () {
      retrieve_theme();
    },
    false
  );

  document.getElementById("special-mode")?.addEventListener("click", (ev) => {
    ev.preventDefault();
    document.body.classList.toggle("dark_mode");

    if (document.body.classList.contains("dark_mode")) {
      localStorage.setItem("color_scheme", "dark");
    } else {
      localStorage.setItem("color_scheme", "light");
    }
    icon.textContent = localStorage.getItem("color_scheme") + "_mode";
  });
})();
