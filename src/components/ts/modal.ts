export function modalFirstSetup() {
  window.addEventListener("click", (ev) => {
    let target = ev.target as HTMLElement;
    if (target.className == "modal") {
      target.style.display = "none";
    }
  });
}

export function setupModalWindow(modal: HTMLDivElement) {
  modal.querySelector<HTMLSpanElement>("span.close")!.onclick = () =>
    (modal.style.display = "none");
}

export function openModal(modal: HTMLElement | string) {
  if (typeof modal == "string") {
    modal = document.getElementById(modal)!;
  }
  modal.style.display = "block";
}
