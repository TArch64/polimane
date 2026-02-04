import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import type { IContextMenuAction } from '@/components/contextMenu';
import { type ListSchema, useHomeStore } from '@/modules/home/stores';
import { CopyIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { SchemasLimitReachedModal } from '@/modules/home/components';
import { useLimitedAction, useSchemasCreatedCounter } from '@/composables/subscription';

export function useSchemaMenuCopy(schemaRef: MaybeRefOrGetter<ListSchema>): IContextMenuAction {
  const schema = computed(() => toValue(schemaRef));

  const homeStore = useHomeStore();

  const router = useRouter();
  const route = useRoute();

  const copySchema = useLimitedAction({
    counter: useSchemasCreatedCounter(),
    modal: useModal(SchemasLimitReachedModal),

    async onAction() {
      const created = await homeStore.copySchema.do(schema.value);

      await router.push({
        name: 'schema-editor',
        params: { schemaId: created.id },
        query: { from: route.path },
      });
    },
  });

  return {
    title: 'Зробити Копію',
    icon: CopyIcon,
    onAction: () => copySchema({ overflowCount: 1 }),
  };
}
