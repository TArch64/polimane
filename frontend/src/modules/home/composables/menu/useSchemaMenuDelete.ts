import { computed, type MaybeRefOrGetter, type Ref, toValue } from 'vue';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useAccessPermissions } from '@/composables';
import { type ListSchema, useHomeStore } from '@/modules/home/stores';
import { TrashIcon } from '@/components/icon';

export function useSchemaMenuDelete(schemaRef: MaybeRefOrGetter<ListSchema>): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const homeStore = useHomeStore();

  const permissions = useAccessPermissions(() => schema.value.access);

  return computed(() => permissions.admin && {
    danger: true,
    title: 'Видалити',
    icon: TrashIcon,
    onAction: () => homeStore.deleteSchema.do(schema.value),
  });
}
