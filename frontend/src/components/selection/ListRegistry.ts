export type SelectionItem = { id: string };

export interface ISelectionItem<I extends SelectionItem> {
  data: I;
  el: HTMLElement;
}

export type SelectionListRegistry<I extends SelectionItem> = Map<I['id'], ISelectionItem<I>>;
