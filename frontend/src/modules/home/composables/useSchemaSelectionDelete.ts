import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { TrashIcon } from '@/components/icon';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction } from '@/composables';
import { useHomeStore } from '@/modules/home/stores';

export function useSchemaSelectionDelete(actionIds: Ref<string[]>, clear: () => void): Ref<MaybeContextMenuAction> {
  const homeStore = useHomeStore();

  const deleteConfirm = useConfirm({
    danger: true,
    control: false,
    message: () => `Ви впевнені, що хочете видалити ${actionIds.value.length} схем?`,
    acceptButton: 'Видалити',
  });

  const deleteSchemas = useAsyncAction(async () => {
    await homeStore.deleteSchema.doMany(actionIds.value);
    clear();
  });

  return computed(() => !!actionIds.value.length && {
    title: 'Видалити Схеми',
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
  });
}
