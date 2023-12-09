export type RenderFunc = (
  windowResizeHandler?: (ev?: UIEvent) => any
) => PageInfo;

export type Listener = {
  name: keyof HTMLElementEventMap;
  elementSelector: string;
  listener: (ev?: Event) => any;
};

export type PageInfo = {
  nodes: NodeListOf<Node>;
  listeners: Listener[];
  postRender?: (page: HTMLElement) => void;
};
