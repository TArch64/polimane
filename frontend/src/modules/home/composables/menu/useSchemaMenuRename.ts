import { computed, type MaybeRefOrGetter, type Ref, toValue } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useAccessPermissions } from '@/composables';
import { type ListSchema, useHomeStore } from '@/modules/home/stores';
import { EditIcon } from '@/components/icon';
import type { ISchema } from '@/models';
import { useModal } from '@/components/modal';
import { SchemaRenameModal } from '@/modules/schemas/shared/modals/rename';

export function useSchemaMenuRename(schemaRef: MaybeRefOrGetter<ListSchema>): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const homeStore = useHomeStore();

  const permissions = useAccessPermissions(() => schema.value.access);
  const renameModal = useModal(SchemaRenameModal);

  return computed(() => permissions.write && {
    title: 'Змінити Назву',
    icon: EditIcon,

    onAction() {
      void renameModal.open({
        schema: schema.value as ISchema,
        updateSchema: (attrs) => homeStore.updateSchema.do(schema.value, attrs),
      });
    },
  });
}
