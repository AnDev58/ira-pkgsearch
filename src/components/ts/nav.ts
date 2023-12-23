export function enableNavbar(renderPage: (place: string) => void) {
  document.querySelectorAll<HTMLAnchorElement>("nav li a").forEach((link) => {
    if (link.href == "javascript:void(0)") {
      return;
    }

    link.onclick = (ev) => {
      ev.preventDefault();

      if (link.classList.contains("active")) {
        return;
      }

      let target = ev.target as HTMLAnchorElement;
      let url = new URL(target.href);

      renderPage(url.pathname);
      let base = import.meta.env.PROD ? "/ira-pkgsearch" : "";
      history.pushState({}, "", base + url.pathname);
    };
  });
}
