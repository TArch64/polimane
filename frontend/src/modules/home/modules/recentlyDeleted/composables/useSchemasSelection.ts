import { computed, reactive, toRef } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import type { ISchemaSelectionAdapter } from '@/modules/home/stores';
import { useDeletedSchemasStore } from '../stores';

export function useSchemasSelection(): ISchemaSelectionAdapter {
  const schemasStore = useDeletedSchemasStore();
  const actionIds = computed(() => [...schemasStore.selected]);

  const actions = computed((): MaybeContextMenuAction[] => {
    if (!actionIds.value.length) {
      return [];
    }

    return [];
  });

  return reactive({
    ids: toRef(schemasStore, 'selected'),
    actions,
    onClear: schemasStore.clearSelection,
  });
}
