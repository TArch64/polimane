import { computed, type MaybeRefOrGetter, type Ref, toValue } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useAccessPermissions } from '@/composables';
import type { ListSchema } from '@/modules/home/stores';
import { PeopleIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import { useSessionStore } from '@/stores';
import { SubscriptionLimit } from '@/enums';

export function useSchemaMenuEditAccess(schemaRef: MaybeRefOrGetter<ListSchema>): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const schemaUsersStore = useSchemaUsersStore();
  const sessionStore = useSessionStore();

  const permissions = useAccessPermissions(() => schema.value.access);
  const sharedAccessLimit = sessionStore.getLimit(SubscriptionLimit.SHARED_ACCESS);
  const accessEditModal = useModal(SchemaAccessEditModal);

  return computed((): MaybeContextMenuAction => permissions.admin && {
    title: 'Редагувати Доступ',
    icon: PeopleIcon,
    disabled: sharedAccessLimit === 1,

    async onAction() {
      await schemaUsersStore.load([schema.value.id]);
      void accessEditModal.open();
    },
  });
}
