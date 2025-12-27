import { computed, reactive, toRef } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import type { ISchemaSelectionAdapter } from '@/modules/home/stores';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction } from '@/composables';
import { TrashIcon } from '@/components/icon';
import { useDeletedSchemasStore } from '../stores';

export function useSchemasSelection(): ISchemaSelectionAdapter {
  const schemasStore = useDeletedSchemasStore();
  const actionIds = computed(() => [...schemasStore.selected]);

  const deleteConfirm = useConfirm({
    danger: true,
    control: false,
    message: () => `Ви впевнені, що хочете видалити остаточно ${actionIds.value.length} схем?`,
    acceptButton: 'Видалити',
  });

  const deleteSchemas = useAsyncAction(async () => {
    await schemasStore.deleteMany(actionIds.value);
    schemasStore.clearSelection();
  });

  const actions = computed((): MaybeContextMenuAction[] => [
    {
      title: 'Видалити Остаточно',
      icon: TrashIcon,
      danger: true,

      async onAction(event) {
        const confirmed = await deleteConfirm.ask({
          virtualTarget: event.menuRect,
        });

        if (confirmed.isAccepted) {
          await deleteSchemas();
        }
      },
    },
  ]);

  return reactive({
    ids: toRef(schemasStore, 'selected'),
    actions,
    onClear: schemasStore.clearSelection,
  });
}
