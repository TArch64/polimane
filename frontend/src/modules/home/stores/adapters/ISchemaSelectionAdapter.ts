import type { MaybeContextMenuAction } from '@/components/contextMenu';

export interface ISchemaSelectionAdapter {
  ids: Set<string>;
  actions: MaybeContextMenuAction[];
  onClear: () => void;
}
