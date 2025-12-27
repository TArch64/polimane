<template>
  <HomeListSchema
    disabled
    :schema
    :menuActions
  />
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { HomeListSchema } from '@/modules/home/components';
import type { ListSchema } from '@/modules/home/stores';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useConfirm } from '@/components/confirm';
import { CornerUpLeftIcon, TrashIcon } from '@/components/icon';
import { useDeletedSchemasStore } from '../stores';

const props = defineProps<{
  schema: ListSchema;
}>();

const schemasStore = useDeletedSchemasStore();
const router = useRouter();

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: 'Ви впевнені, що хочете видалити цю схему остаточно?',
  acceptButton: 'Видалити',
});

async function onDeletableFinish() {
  if (!schemasStore.schemas.length) {
    await router.push({ name: 'home' });
  }
}

const menuActions: MaybeContextMenuAction[] = [
  {
    title: 'Відновити Схему',
    icon: CornerUpLeftIcon,
    async onAction() {
      await schemasStore.restoreSchema(props.schema);
      await onDeletableFinish();
    },
  },

  {
    danger: true,
    title: 'Видалити Остаточно',
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
