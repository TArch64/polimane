import { computed, reactive, toRef } from 'vue';
import { useRouter } from 'vue-router';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import type { ISchemaSelectionAdapter } from '@/modules/home/stores';
import { useConfirm } from '@/components/confirm';
import { useAsyncAction } from '@/composables';
import { CornerUpLeftIcon, TrashIcon } from '@/components/icon';
import { useSchemasCreatedLimit } from '@/composables/subscription';
import { useDeletedSchemasStore } from '../stores';

export function useSchemasSelection(): ISchemaSelectionAdapter {
  const schemasStore = useDeletedSchemasStore();
  const actionIds = computed(() => [...schemasStore.selected]);

  const router = useRouter();
  const schemasCreatedLimit = useSchemasCreatedLimit();

  const deleteConfirm = useConfirm({
    danger: true,
    control: false,
    message: () => `Ви впевнені, що хочете видалити остаточно ${actionIds.value.length} схем?`,
    acceptButton: 'Видалити',
  });

  async function onDeletableComplete() {
    if (schemasStore.schemas.length) {
      schemasStore.clearSelection();
    } else {
      await router.push({ name: 'home' });
    }
  }

  const deleteSchemas = useAsyncAction(async () => {
    await schemasStore.deleteMany(actionIds.value);
    await onDeletableComplete();
  });

  const restoreSchemas = useAsyncAction(async () => {
    await schemasStore.restoreMany(actionIds.value);
    await onDeletableComplete();
  });

  const actions = computed((): MaybeContextMenuAction[] => [
    {
      title: 'Відновити Схеми',
      icon: CornerUpLeftIcon,
      disabled: schemasCreatedLimit.isNear(actionIds.value.length),
      onAction: restoreSchemas,
    },

    {
      title: 'Видалити Остаточно',
      icon: TrashIcon,
      danger: true,

      async onAction(event) {
        const confirmation = await deleteConfirm.ask({
          virtualTarget: event.menuRect,
        });

        if (confirmation.isAccepted) {
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
