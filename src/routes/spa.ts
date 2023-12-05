import indexPage from "../pages/index.htm?raw";
import aboutPage from "../pages/about.htm?raw";
import notFoundPage from "../pages/404.htm?raw";
import { renderLoginSingupPage, renderStaticPage } from "./renders";

const routes = {
  "/": ["nav-index", renderStaticPage(indexPage)],
  "/about": ["nav-about", renderStaticPage(aboutPage)],
  "/account": ["nav-account", renderLoginSingupPage()],
};

export function renderPage(
  path: string,
  previousPath: string | null,
  windowResizeHandler?: (ev?: UIEvent) => any
) {
  let useAnimation = previousPath != path && previousPath != null;
  let page: NodeListOf<Node>;
  let activeLinkID: string | undefined = undefined;

  if (windowResizeHandler) {
    window.onresize = windowResizeHandler;
  }

  if (path in routes) {
    let pageInfo = routes[path as keyof typeof routes];
    activeLinkID = pageInfo[0] as string;
    page = (pageInfo[1] as renderFunc)(windowResizeHandler);
  } else {
    page = renderStaticPage(notFoundPage)();
    useAnimation = false;
  }

  document.querySelector("nav .active")?.classList.remove("active");
  if (activeLinkID) {
    document.getElementById(activeLinkID)!.classList.add("active");
  }

  clonePage(page, useAnimation);
}

export type renderFunc = (
  windowResizeHandler?: (ev?: UIEvent) => any
) => NodeListOf<Node>;

function clonePage(nodes: NodeListOf<Node>, useAnimation: boolean) {
  let destination = document.getElementById("app")!;

  destination.childNodes.forEach((child) => child.remove());

  let copy = () =>
    nodes.forEach(function (element) {
      let clonedElement = element.cloneNode(true);
      destination.appendChild(clonedElement);
    });

  if (!useAnimation) {
    copy();
    return;
  }

  destination.classList.add("pre-animation");
  setTimeout(() => {
    copy();
    destination.classList.remove("pre-animation");
  }, 200);
}
