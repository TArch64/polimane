import { computed, type Ref } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { PeopleIcon } from '@/components/icon';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import { useModal } from '@/components/modal';
import { SubscriptionLimit } from '@/enums';
import { useSessionStore } from '@/stores';

export function useSchemaSelectionEditAccess(actionIds: Ref<string[]>): Ref<MaybeContextMenuAction> {
  const sessionStore = useSessionStore();
  const schemaUsersStore = useSchemaUsersStore();

  const sharedAccessLimit = sessionStore.getLimit(SubscriptionLimit.SHARED_ACCESS);
  const accessEditModal = useModal(SchemaAccessEditModal);

  return computed((): MaybeContextMenuAction => !!actionIds.value.length && {
    title: 'Редагувати Доступ',
    icon: PeopleIcon,
    disabled: sharedAccessLimit === 1,

    async onAction() {
      await schemaUsersStore.load(actionIds.value);
      void accessEditModal.open();
    },
  });
}
