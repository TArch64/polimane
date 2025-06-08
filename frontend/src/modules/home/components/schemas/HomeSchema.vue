<template>
  <Card ref="cardRef" interactable :binding="cardBinding">
    {{ schema.name }}
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { RouterLink } from '@/components/button';
import type { ISchema } from '@/models';
import { makeBinding } from '@/components/binding';
import { useContextMenu } from '@/components/contextMenu';
import { useDomRef } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { TrashIcon } from '@/components/icon';
import { useSchemasStore } from '@/modules/home/stores';

const props = defineProps<{
  schema: ISchema;
}>();

const schemasStore = useSchemasStore();
const cardRef = useDomRef<HTMLElement>();

const cardBinding = makeBinding(RouterLink, () => ({
  to: {
    name: 'schema-editor',
    params: { schemaId: props.schema.id },
  },
}));

const deleteConfirm = useConfirm({
  danger: true,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
});

useContextMenu({
  el: cardRef,
  title: props.schema.name,

  actions: [
    {
      danger: true,
      title: 'Видалити Схему',
      icon: TrashIcon,

      async onAction(event) {
        if (await deleteConfirm.ask({ virtualTarget: event.menuRect })) {
          await schemasStore.deleteSchema(props.schema);
        }
      },
    },
  ],
});
</script>
