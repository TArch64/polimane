import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { PeopleIcon } from '@/components/icon';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import { useModal } from '@/components/modal';

export function useSchemaSelectionEditAccess(actionIds: Ref<string[]>): Ref<MaybeContextMenuAction> {
  const schemaUsersStore = useSchemaUsersStore();

  const accessEditModal = useModal(SchemaAccessEditModal);

  return computed(() => !!actionIds.value.length && {
    title: 'Редагувати доступ',
    icon: PeopleIcon,

    async onAction() {
      await schemaUsersStore.load(actionIds.value);
      void accessEditModal.open();
    },
  });
}
