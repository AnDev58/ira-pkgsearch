import { PackageModel } from "../models/packages";
import logInPage from "../pages/login.htm?raw";
import searchPage from "../pages/search.htm?raw";
import { RenderFunc } from "./types/render";

export function renderStaticPage(page: string): RenderFunc {
  return (_?: (ev?: UIEvent) => any) => {
    return {
      nodes: new DOMParser().parseFromString(page, "text/html").body.childNodes,
      listeners: [],
    };
  };
}

export function renderLoginSingupPage(): RenderFunc {
  const page = new DOMParser().parseFromString(logInPage, "text/html");
  return (windowResizeHandler?: (ev?: UIEvent) => any) => {
    changeLoginSingupSwitcher(page);

    window.onresize = (ev) => {
      changeLoginSingupSwitcher(document);
      if (windowResizeHandler) {
        windowResizeHandler(ev);
      }
    };
    return {
      nodes: page.body.childNodes,
      listeners: [
        {
          name: "click",
          elementSelector: "#login-singup-switcher",
          listener: () => {
            document.querySelectorAll("form").forEach((form) => {
              form.hidden = !form.hidden;
            });
            changeLoginSingupSwitcher(document);
          },
        },
      ],
    };
  };
}

export function renderSearchPage(): RenderFunc {
  const page = new DOMParser().parseFromString(searchPage, "text/html");
  return (windowResizeHandler?: (ev?: UIEvent) => any) => {
    changeVisibility(page);

    window.onresize = (ev) => {
      changeVisibility(document);
      if (windowResizeHandler) {
        windowResizeHandler(ev);
      }
    };

    return {
      nodes: page.body.childNodes,
      listeners: [
        {
          name: "keyup",
          elementSelector: "search input",
          listener: searchPackageHandler,
        },
      ],
    };
  };
}

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
      if (form.hidden) {
        place = form.querySelector("h2")!.textContent!.toLowerCase();
        break;
      }
    }
    switcherText.textContent = "Switch to " + place;
  }
}

function searchPackageHandler() {
  const tableBody = document.querySelector<HTMLTableSectionElement>(
    "#search-results tbody"
  )!;
  const section = document.querySelector<HTMLElement>(
    "#search-results section"
  )!;

  tableBody.innerHTML = "";
  section.innerHTML = "";

  const searchField = document.querySelector<HTMLInputElement>("search input")!;
  if (searchField.value.length <= 1) {
    return;
  }
  const pkgs: PackageModel.Package[] = PackageModel.search(searchField.value);
  pkgs.forEach((pkg) => {
    const article = buildResultArticle(pkg);
    tableBody.appendChild(article[0]);
    section.appendChild(article[1]);
  });
}

function buildResultArticle(
  pkg: PackageModel.Package
): [HTMLTableRowElement, HTMLElement] {
  const resultRow = document.createElement("tr");
  const resultArticle = document.createElement("article");
  resultArticle.onclick = (ev: MouseEvent) => showPackageDetails(ev, pkg);

  const tableName = document.createElement("td");
  const articleName = document.createElement("h2");
  tableName.textContent = articleName.textContent = pkg.name;

  const tableVersion = document.createElement("td");
  const articleVersion = document.createElement("span");
  tableVersion.textContent = articleVersion.textContent = pkg.version;

  const tableDescription = document.createElement("td");
  const articleDescription = document.createElement("p");
  tableDescription.textContent = articleDescription.textContent =
    pkg.description;

  const owner = document.createElement("td");
  owner.textContent = pkg.owner;

  const info = document.createElement("td");
  info.classList.add("material-symbols-outlined");
  info.style.cursor = "pointer";
  info.textContent = "info";
  info.onclick = (ev) => showPackageDetails(ev, pkg);

  resultRow.append(tableName, tableVersion, tableDescription, owner, info);
  articleName.appendChild(articleVersion);
  resultArticle.append(articleName, articleDescription);
  return [resultRow, resultArticle];
}

function changeVisibility(where: Document) {
  const section = where.querySelector<HTMLElement>("#search-results section")!;
  const table = where.querySelector<HTMLTableElement>("#search-results table")!;

  table.hidden = innerWidth <= 600;
  section.hidden = !table.hidden;
}

function showPackageDetails(ev: MouseEvent, pkg: PackageModel.Package) {
  ev.preventDefault();
  alert(pkg.owner);
}
