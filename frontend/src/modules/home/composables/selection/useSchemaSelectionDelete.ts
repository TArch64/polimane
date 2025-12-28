import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { TrashIcon } from '@/components/icon';
import { useAsyncAction } from '@/composables';
import { useHomeStore } from '@/modules/home/stores';

export function useSchemaSelectionDelete(actionIds: Ref<string[]>, clear: () => void): Ref<MaybeContextMenuAction> {
  const homeStore = useHomeStore();

  const deleteSchemas = useAsyncAction(async () => {
    await homeStore.deleteSchema.doMany(actionIds.value);
    clear();
  });

  return computed(() => !!actionIds.value.length && {
    title: 'Видалити Схеми',
    icon: TrashIcon,
    danger: true,
    onAction: deleteSchemas,
  });
}
