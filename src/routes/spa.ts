import indexPage from "../pages/index.htm?raw";
import aboutPage from "../pages/about.htm?raw";
import notFoundPage from "../pages/404.htm?raw";
import {
  renderLoginSingupPage,
  renderSearchPage,
  renderStaticPage,
} from "./renders";
import { Listener, PageInfo } from "./types/render";

export const routes = {
  "/": {
    activeLinkID: "nav-index", // "" if don't have
    renderFunc: renderStaticPage(indexPage),
  },
  "/about": {
    activeLinkID: "nav-about",
    renderFunc: renderStaticPage(aboutPage),
  },
  "/account": {
    activeLinkID: "nav-account",
    renderFunc: renderLoginSingupPage(),
  },
  "/search": {
    activeLinkID: "nav-search",
    renderFunc: renderSearchPage(),
  },
};

export function renderPage(path: string, useAnimation: boolean) {
  let pageInfo: PageInfo;

  document.querySelector("nav .active")?.classList.remove("active");

  if (path in routes) {
    // Known URL
    const pageRoute = routes[path as keyof typeof routes];
    pageInfo = pageRoute.renderFunc();

    if (pageRoute.activeLinkID !== "") {
      // Special visual effect - active status of link
      document.getElementById(pageRoute.activeLinkID)?.classList.add("active");
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

  if (!useAnimation) {
    destination.innerHTML = "";
    copyNodes(pageInfo, destination);
    return;
  }

  destination.classList.add("pre-animation");
  setTimeout(() => {
    destination.innerHTML = "";
    copyNodes(pageInfo, destination);
    destination.classList.remove("pre-animation");
  }, 300);
}

function copyNodes(pageInfo: PageInfo, destination: HTMLElement) {
  pageInfo.nodes.forEach(function (element) {
    const clonedElement = element.cloneNode(true) as HTMLElement;
    destination.appendChild(clonedElement);
  });
  addAllEventListeners(pageInfo.listeners);
  if (pageInfo.postRender) {
    pageInfo.postRender(destination);
  }
}

function addAllEventListeners(listeners: Listener[]) {
  listeners.forEach((listener) => {
    document
      .querySelector("#app " + listener.elementSelector)!
      .addEventListener(listener.name, listener.listener);
  });
}
