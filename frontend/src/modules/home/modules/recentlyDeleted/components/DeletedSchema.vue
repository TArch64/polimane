<template>
  <HomeListSchema
    disabled
    :schema
    :menuActions
  />
</template>

<script setup lang="ts">
import { HomeListSchema } from '@/modules/home/components';
import type { ListSchema } from '@/modules/home/stores';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useConfirm } from '@/components/confirm';
import { TrashIcon } from '@/components/icon';
import { useDeletedSchemasStore } from '../stores';

const props = defineProps<{
  schema: ListSchema;
}>();

const schemasStore = useDeletedSchemasStore();

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: 'Ви впевнені, що хочете видалити цю схему остаточно?',
  acceptButton: 'Видалити',
});

const menuActions: MaybeContextMenuAction[] = [
  {
    danger: true,
    title: 'Видалити Остаточно',
    icon: TrashIcon,

    async onAction(event) {
      const confirmed = await deleteConfirm.ask({
        virtualTarget: event.menuRect,
      });

      if (confirmed.isAccepted) {
        await schemasStore.deleteSchema(props.schema);
      }
    },
  },
];
</script>
