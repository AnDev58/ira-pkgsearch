import logInPage from "../pages/login.htm?raw";
import { renderFunc } from "./spa";
function changeLoginSingupSwitcher(where: Document) {
  let switcherText = where
    .getElementById("login-singup-switcher")
    ?.querySelector("span");

  if (!switcherText) {
    return;
  }
  if (innerWidth <= 400) {
    switcherText.classList.add("material-symbols-outlined");
    switcherText.textContent = "swap_horiz";
  } else {
    switcherText.classList.remove("material-symbols-outlined");
    let place = "another page";
    for (let i = 0; i < where.forms.length; i++) {
      const form = where.forms[i];
      if (!form.hidden) {
        place = form.querySelector("h2")!.textContent!.toLowerCase();
        break;
      }
    }
    switcherText.textContent = "Switch to " + place;
  }
}

export function renderStaticPage(page: string): () => NodeListOf<Node> {
  return () => new DOMParser().parseFromString(page, "text/html").childNodes;
}

export function renderLoginSingupPage(): renderFunc {
  const page = new DOMParser().parseFromString(logInPage, "text/html");
  page.getElementById("login-singup-switcher")!.onclick = () => {
    document.querySelectorAll("form").forEach((form) => {
      form.hidden = !form.hidden;
    });
    changeLoginSingupSwitcher(document);
  };
  return (windowResizeHandler?: (ev?: UIEvent) => any) => {
    changeLoginSingupSwitcher(page);

    window.onresize = (ev) => {
      changeLoginSingupSwitcher(document);
      if (windowResizeHandler) {
        windowResizeHandler(ev);
      }
    };
    return page.childNodes;
  };
}
