import type { MaybeContextMenuAction } from '@/components/contextMenu';

export interface ISchemaSelectionStrategy {
  ids: Set<string>;
  actions: MaybeContextMenuAction[];
  onClear: () => void;
}
