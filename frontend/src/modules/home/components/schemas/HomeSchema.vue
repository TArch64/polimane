<template>
  <Card ref="cardRef" class="home-schema" interactable :binding="cardBinding">
    <img
      :src="screenshotUrl"
      :alt="`Скріншот схеми ${schema.name}`"
      draggable="false"
      decoding="async"
      class="home-schema__screenshot"
      v-if="screenshotUrl"
    >

    {{ schema.name }}
  </Card>
</template>

<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router';
import { computed } from 'vue';
import { Card } from '@/components/card';
import type { ISchema } from '@/models';
import { makeBinding } from '@/components/binding';
import { useContextMenu } from '@/components/contextMenu';
import { useDomRef } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { CopyIcon, TrashIcon } from '@/components/icon';
import { useSchemasStore } from '@/modules/home/stores';
import { buildCdnUrl } from '@/helpers/buildCdnUrl';

const props = defineProps<{
  schema: ISchema;
}>();

const router = useRouter();
const schemasStore = useSchemasStore();
const cardRef = useDomRef<HTMLElement>();

const cardBinding = makeBinding(RouterLink, () => ({
  to: {
    name: 'schema-editor',
    params: { schemaId: props.schema.id },
  },
}));

const screenshotUrl = computed(() => buildCdnUrl(props.schema.screenshotPath));

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
});

useContextMenu({
  el: cardRef,
  title: props.schema.name,
  control: false,

  actions: [
    {
      title: 'Зробити Копію',
      icon: CopyIcon,

      async onAction() {
        const created = await schemasStore.copySchema(props.schema);

        await router.push({
          name: 'schema-editor',
          params: {
            schemaId: created.id,
          },
        });
      },
    },

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

<style scoped>
@layer page {
  .home-schema {
    overflow: clip;
  }

  .home-schema__screenshot {
    display: block;
    aspect-ratio: 16 / 9;
    object-fit: contain;
    object-position: center;
    border-bottom: var(--divider);
    margin: calc(0px - var(--card-padding-top)) calc(0px - var(--card-padding-right)) 8px calc(0px - var(--card-padding-left));
    width: calc(100% + var(--card-padding-left) + var(--card-padding-right));
    background-color: v-bind("schema.backgroundColor");
  }
}
</style>
