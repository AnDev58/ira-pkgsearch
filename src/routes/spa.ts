import indexPage from "../pages/index.htm?raw";
import aboutPage from "../pages/about.htm?raw";
import notFoundPage from "../pages/404.htm?raw";
import logInPage from "../pages/login.htm?raw";

function changeLoginSingupSwitcher(where: Document) {
  let switcherText = where
    .getElementById("login-singup-switcher")
    ?.querySelector("span")!;
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

export function renderPage(
  path: string,
  previous: string | null,
  windowResizeHandler?: (ev?: UIEvent) => any
): ((ev?: UIEvent) => any) | undefined {
  let useAnimation = previous != path && previous != null;
  let page: NodeListOf<Node>;
  let activeLinkID: string | undefined = undefined;
  if (windowResizeHandler) {
    window.onresize = windowResizeHandler;
  }
  switch (path) {
    case "/":
      page = new DOMParser().parseFromString(indexPage, "text/html").childNodes;
      activeLinkID = "nav-index";
      break;

    case "/about":
      page = new DOMParser().parseFromString(aboutPage, "text/html").childNodes;
      activeLinkID = "nav-about";
      break;

    case "/account":
      let parsedDOM = new DOMParser().parseFromString(logInPage, "text/html");

      changeLoginSingupSwitcher(parsedDOM);
      window.onresize = (ev) => {
        changeLoginSingupSwitcher(document);
        if (windowResizeHandler) {
          windowResizeHandler(ev);
        }
      };

      parsedDOM.getElementById("login-singup-switcher")!.onclick = () => {
        document.querySelectorAll("form").forEach((form) => {
          form.hidden = !form.hidden;
        });
        changeLoginSingupSwitcher(document);
      };

      activeLinkID = "nav-account";
      page = parsedDOM.childNodes;
      break;

    default:
      useAnimation = false;
      page = new DOMParser().parseFromString(
        notFoundPage,
        "text/html"
      ).childNodes;
      break;
  }

  document.querySelector("nav .active")?.classList.remove("active");
  if (activeLinkID) {
    document.getElementById(activeLinkID)!.classList.add("active");
  }

  let container = document.getElementById("app")!;
  for (let i = 0; i < container.children.length; i++) {
    container.children[i].remove();
  }

  if (!useAnimation) {
    page.forEach((el) => container.appendChild(el));
    return;
  }

  container.classList.add("pre-animation");
  setTimeout(() => {
    page.forEach((el) => container.appendChild(el));
    container.classList.remove("pre-animation");
  }, 200);
}
