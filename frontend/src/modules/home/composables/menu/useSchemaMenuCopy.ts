import { computed, type MaybeRefOrGetter, type Ref, toValue } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { type ListSchema, useHomeStore } from '@/modules/home/stores';
import { CopyIcon } from '@/components/icon';

export function useSchemaMenuCopy(schemaRef: MaybeRefOrGetter<ListSchema>): Ref<MaybeContextMenuAction> {
  const schema = computed(() => toValue(schemaRef));

  const homeStore = useHomeStore();

  const router = useRouter();
  const route = useRoute();

  return computed(() => ({
    title: 'Зробити Копію',
    icon: CopyIcon,

    async onAction() {
      const created = await homeStore.copySchema.do(schema.value);

      await router.push({
        name: 'schema-editor',
        params: { schemaId: created.id },
        query: { from: route.path },
      });
    },
  }));
}
