import "./themes/css/basic.css";
import "./themes/css/nav.css";
import "./themes/css/special.css";
import "./themes/css/modal.css";
import "./styles/forms.css";
import "./styles/style.css";
import "./components/css/nav.css";
import "./components/css/modal.css";

import { renderPage } from "./routes/spa";
import { retrieveTheme, switchTheme } from "./themes/basics";
import { enableNavbar } from "./components/ts/nav";
import { modalFirstSetup } from "./components/ts/modal";

// The ONLY entry point
(function () {
  retrieveTheme();
  renderPage(location.pathname, false);

  // Setting up components
  modalFirstSetup();
  enableNavbar((place) => renderPage(place, true));

  window.addEventListener(
    "popstate",
    () => {
      renderPage(location.pathname, false);
    },
    false
  );

  document
    .getElementById("special-dark-light-theme")
    ?.addEventListener("click", switchTheme);
})();
