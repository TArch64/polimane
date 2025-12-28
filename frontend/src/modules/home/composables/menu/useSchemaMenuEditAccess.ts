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

export function useSchemaMenuEditAccess(schemaRef: MaybeRefOrGetter<ListSchema>): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const schemaUsersStore = useSchemaUsersStore();

  const permissions = useAccessPermissions(() => schema.value.access);
  const accessEditModal = useModal(SchemaAccessEditModal);

  return computed(() => permissions.admin && {
    title: 'Редагувати Доступ',
    icon: PeopleIcon,

    async onAction() {
      await schemaUsersStore.load([schema.value.id]);
      void accessEditModal.open();
    },
  });
}
