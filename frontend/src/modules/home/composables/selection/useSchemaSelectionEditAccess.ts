import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { PeopleIcon } from '@/components/icon';
import { SchemaAccessEditModal } from '@/modules/schemas/shared/modals/accessEdit';
import { useModal } from '@/components/modal';

export function useSchemaSelectionEditAccess(actionIds: Ref<string[]>): Ref<MaybeContextMenuAction> {
  const accessEditModal = useModal(SchemaAccessEditModal);

  return computed(() => !!actionIds.value.length && {
    title: 'Редагувати Доступ',
    icon: PeopleIcon,

    async onAction() {
      void accessEditModal.open({
        schemaIds: actionIds.value,
      });
    },
  });
}
