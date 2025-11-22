import { computed, reactive, toRef } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { AccessLevel } from '@/enums';
import type { ISchemaSelectionStrategy } from '@/modules/home/stores';
import {
  useSchemaSelectionAddToFolder,
  useSchemaSelectionDelete,
  useSchemaSelectionEditAccess,
} from '@/modules/home/composables';
import { useFolderSchemasStore } from '../stores';

export function useSchemasSelection(): ISchemaSelectionStrategy {
  const schemasStore = useFolderSchemasStore();

  const allActionIds = computed(() => [...schemasStore.selected]);

  const adminActionIds = computed(() => {
    return schemasStore.filterIdsByAccess(schemasStore.selected, AccessLevel.ADMIN);
  });

  const addToFolderAction = useSchemaSelectionAddToFolder(allActionIds);
  const editAccessAction = useSchemaSelectionEditAccess(adminActionIds);
  const deleteAction = useSchemaSelectionDelete(adminActionIds, schemasStore.clearSelection);

  const actions = computed((): MaybeContextMenuAction[] => {
    if (!allActionIds.value.length) {
      return [];
    }

    return [
      addToFolderAction.value,
      editAccessAction.value,
      deleteAction.value,
    ];
  });

  return reactive({
    ids: toRef(schemasStore, 'selected'),
    actions,
    onClear: schemasStore.clearSelection,
  });
}
