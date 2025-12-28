import { computed, reactive, toRef } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { AccessLevel } from '@/enums';
import type { ISchemaSelectionAdapter } from '@/modules/home/stores';
import {
  useSchemaSelectionAddToFolder,
  useSchemaSelectionDelete,
  useSchemaSelectionEditAccess,
} from '@/modules/home/composables';
import { FolderRemoveIcon } from '@/components/icon';
import { useFolderSchemasStore } from '../stores';

export function useSchemasSelection(): ISchemaSelectionAdapter {
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

      {
        title: 'Видалити з Директорії',
        icon: FolderRemoveIcon,
        async onAction() {
          await schemasStore.removeManyFromFolder(allActionIds.value);
          schemasStore.clearSelection();
        },
      },

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
