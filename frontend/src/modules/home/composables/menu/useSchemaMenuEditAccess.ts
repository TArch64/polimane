import { computed, type MaybeRefOrGetter, type Ref, toValue } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useAccessPermissions } from '@/composables';
import type { ListSchema } from '@/modules/home/stores';
import { PeopleIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { SchemaAccessEditModal } from '@/modules/schemas/shared/modals/accessEdit';

export function useSchemaMenuEditAccess(schemaRef: MaybeRefOrGetter<ListSchema>): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const permissions = useAccessPermissions(() => schema.value.access);
  const accessEditModal = useModal(SchemaAccessEditModal);

  return computed((): MaybeContextMenuAction => permissions.admin && {
    title: 'Редагувати Доступ',
    icon: PeopleIcon,

    async onAction() {
      void accessEditModal.open({
        schemaIds: [schema.value.id],
      });
    },
  });
}
