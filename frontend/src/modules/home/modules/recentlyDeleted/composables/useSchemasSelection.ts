import { computed, reactive, toRef } from 'vue';
import { useRouter } from 'vue-router';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import type { ISchemaSelectionAdapter } from '@/modules/home/stores';
import { useConfirm } from '@/components/confirm';
import { CornerUpLeftIcon, TrashIcon } from '@/components/icon';
import { useLimitedAction, useSchemasCreatedCounter } from '@/composables/subscription';
import { useModal } from '@/components/modal';
import { SchemasLimitReachedModal } from '@/modules/home/components';
import { useDeletedSchemasStore } from '../stores';

export function useSchemasSelection(): ISchemaSelectionAdapter {
  const schemasStore = useDeletedSchemasStore();
  const actionIds = computed(() => [...schemasStore.selected]);

  const router = useRouter();
  const schemasCreatedCounter = useSchemasCreatedCounter();

  const deleteConfirm = useConfirm({
    danger: true,
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

  const restoreSchema = useLimitedAction({
    counter: schemasCreatedCounter,
    overflow: () => actionIds.value.length,
    modal: useModal(SchemasLimitReachedModal),

    async onAction() {
      await schemasStore.restoreMany(actionIds.value);
      await onDeletableComplete();
    },
  });

  const actions: MaybeContextMenuAction[] = [
    {
      title: 'Відновити Схеми',
      icon: CornerUpLeftIcon,

      async onAction() {
        const afterAdd = schemasCreatedCounter.current + actionIds.value.length;
        return restoreSchema({ overflowCount: afterAdd - schemasCreatedCounter.max });
      },
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
          await schemasStore.deleteMany(actionIds.value);
          await onDeletableComplete();
        }
      },
    },
  ];

  return reactive({
    ids: toRef(schemasStore, 'selected'),
    actions,
    onClear: schemasStore.clearSelection,
  });
}
