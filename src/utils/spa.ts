import indexPage from "../pages/index.htm?raw";
import aboutPage from "../pages/about.htm?raw";
import notFoundPage from "../pages/404.htm?raw";
import logInPage from "../pages/login.htm?raw";

export function renderPage(path: string, previous: string | null) {
  let useAnimation = previous != path && previous != null;
  let page: string;
  let activeLinkID: string | undefined = undefined;
  switch (path) {
    case "/":
      page = indexPage;
      activeLinkID = "nav-index";
      break;
    case "/about":
      page = aboutPage;
      activeLinkID = "nav-about";
      break;
    case "/account":
      page = logInPage;
      activeLinkID = "nav-account";
      break;
    default:
      useAnimation = false;
      page = notFoundPage;
      break;
  }

  document.querySelector("nav .active")?.classList.remove("active");
  if (activeLinkID) {
    document.getElementById(activeLinkID)!.classList.add("active");
  }

  let container = document.getElementById("app")!;
  if (!useAnimation) {
    container.innerHTML = page;
    return;
  }
  container.classList.add("pre-animation");
  setTimeout(() => {
    container.innerHTML = page;
    container.classList.remove("pre-animation");
  }, 200);
}
