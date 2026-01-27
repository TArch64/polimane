import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import type { IContextMenuAction } from '@/components/contextMenu';
import { type ListSchema, useHomeStore } from '@/modules/home/stores';
import { CopyIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { SchemasLimitReachedModal } from '@/modules/home/components';
import { useSchemasCreatedCounter } from '@/composables/subscription';

export function useSchemaMenuCopy(schemaRef: MaybeRefOrGetter<ListSchema>): IContextMenuAction {
  const schema = computed(() => toValue(schemaRef));

  const homeStore = useHomeStore();

  const router = useRouter();
  const route = useRoute();
  const schemasCreatedCounter = useSchemasCreatedCounter();
  const schemasLimitReachedModal = useModal(SchemasLimitReachedModal);

  return {
    title: 'Зробити Копію',
    icon: CopyIcon,

    async onAction() {
      if (schemasCreatedCounter.isReached) {
        const isUpgraded = await schemasLimitReachedModal.open({
          overflowCount: 1,
        });

        if (!isUpgraded) {
          return;
        }
      }

      const created = await homeStore.copySchema.do(schema.value);

      await router.push({
        name: 'schema-editor',
        params: { schemaId: created.id },
        query: { from: route.path },
      });
    },
  };
}
