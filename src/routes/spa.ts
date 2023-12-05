import indexPage from "../pages/index.htm?raw";
import aboutPage from "../pages/about.htm?raw";
import notFoundPage from "../pages/404.htm?raw";
import { renderLoginSingupPage, renderStaticPage } from "./renders";
import { Listener, PageInfo, RenderFunc } from "./types/render";

// Structure: {"url": ["navbar's anchor ID, if present, or null otherwhise", RenderFunc]}
export const routes = {
  "/": ["nav-index", renderStaticPage(indexPage)],
  "/about": ["nav-about", renderStaticPage(aboutPage)],
  "/account": ["nav-account", renderLoginSingupPage()],
};

export function renderPage(path: string, useAnimation: boolean) {
  let pageInfo: PageInfo;

  document.querySelector("nav .active")?.classList.remove("active");

  if (path in routes) {
    // Known URL
    const pageRoute = routes[path as keyof typeof routes];
    pageInfo = (pageRoute[1] as RenderFunc)();

    if (pageRoute[0]) {
      // Special visual effect - active status of link
      document.getElementById(pageRoute[0] as string)?.classList.add("active");
    }
  } else {
    // Unknown URL, need 404 page
    pageInfo = renderStaticPage(notFoundPage)();

    useAnimation = false; // Special visual effect - no animation
  }

  clonePage(pageInfo, useAnimation);
}

function clonePage(pageInfo: PageInfo, useAnimation: boolean) {
  let destination = document.getElementById("app")!;

  destination.childNodes.forEach((child) => child.remove());

  if (!useAnimation) {
    copyNodes(pageInfo, destination);
    return;
  }

  destination.classList.add("pre-animation");
  setTimeout(() => {
    copyNodes(pageInfo, destination);
    destination.classList.remove("pre-animation");
  }, 200);
}

function copyNodes(pageInfo: PageInfo, destination: HTMLElement) {
  pageInfo.nodes.forEach(function (element) {
    const clonedElement = element.cloneNode(true) as HTMLElement;
    destination.appendChild(clonedElement);
  });
  addAllEventListeners(pageInfo.listeners);
}

function addAllEventListeners(listeners: Listener[]) {
  listeners.forEach((listener) => {
    document
      .querySelector("#app " + listener.elementSelector)!
      .addEventListener(listener.name, listener.listener);
  });
}
