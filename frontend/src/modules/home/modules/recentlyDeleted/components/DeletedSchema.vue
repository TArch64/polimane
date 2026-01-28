<template>
  <HomeListSchema disabled :schema :menuActions />
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { HomeListSchema, SchemasLimitReachedModal } from '@/modules/home/components';
import type { ListSchema } from '@/modules/home/stores';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useConfirm } from '@/components/confirm';
import { CornerUpLeftIcon, TrashIcon } from '@/components/icon';
import { useLimitedAction, useSchemasCreatedCounter } from '@/composables/subscription';
import { useModal } from '@/components/modal';
import { useDeletedSchemasStore } from '../stores';

const props = defineProps<{
  schema: ListSchema;
}>();

const schemasStore = useDeletedSchemasStore();

const router = useRouter();

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цю схему остаточно?',
  acceptButton: 'Видалити',
});

async function onDeletableFinish() {
  if (!schemasStore.schemas.length) {
    await router.push({ name: 'home' });
  }
}

const restoreSchema = useLimitedAction({
  counter: useSchemasCreatedCounter(),
  modal: useModal(SchemasLimitReachedModal),

  async onAction() {
    await schemasStore.restoreSchema(props.schema);
    await onDeletableFinish();
  },
});

const menuActions: MaybeContextMenuAction[] = [
  {
    title: 'Відновити Схему',
    icon: CornerUpLeftIcon,
    onAction: () => restoreSchema({ overflowCount: 1 }),
  },

  {
    title: 'Видалити Остаточно',
    danger: true,
    icon: TrashIcon,

    async onAction(event) {
      const confirmation = await deleteConfirm.ask({
        virtualTarget: event.menuRect,
      });

      if (confirmation.isAccepted) {
        await schemasStore.deleteSchema(props.schema);
        await onDeletableFinish();
      }
    },
  },
];
</script>
