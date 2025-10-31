<template>
  <Card
    interactable
    ref="cardRef"
    class="home-schema"
    :active="isSelected"
    :binding="cardBinding"
  >
    <img
      :src="screenshotUrl"
      :alt="`Скріншот схеми ${schema.name}`"
      draggable="false"
      decoding="async"
      loading="lazy"
      class="home-schema__screenshot"
      v-if="screenshotUrl"
    >

    <div class="home-schema__screenshot" v-else />

    {{ schema.name }}
  </Card>
</template>

<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router';
import { computed } from 'vue';
import { Card } from '@/components/card';
import { makeBinding } from '@/components/binding';
import { useContextMenu } from '@/components/contextMenu';
import { useAccessPermissions, useDomRef } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { CopyIcon, EditIcon, PeopleIcon, TrashIcon } from '@/components/icon';
import { buildCdnUrl } from '@/helpers/buildCdnUrl';
import { useModal } from '@/components/modal';
import SchemaRenameModal from '@/modules/schemas/shared/modals/SchemaRenameModal.vue';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import type { ISchema } from '@/models';
import { type SchemaListItem, useSchemasStore } from '../../stores';

const props = defineProps<{
  schema: SchemaListItem;
}>();

const router = useRouter();

const schemasStore = useSchemasStore();
const schemaUsersStore = useSchemaUsersStore();

const cardRef = useDomRef<HTMLElement>();

const isSelected = computed(() => schemasStore.selected.has(props.schema.id));
const permissions = useAccessPermissions(() => props.schema.access);

const renameModal = useModal(SchemaRenameModal);
const accessEditModal = useModal(SchemaAccessEditModal);

const cardBinding = makeBinding(RouterLink, () => ({
  draggable: false,

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
    permissions.write && {
      title: 'Переназвати',
      icon: EditIcon,

      onAction() {
        renameModal.open({
          schema: props.schema as ISchema,
          updateSchema: (attrs) => schemasStore.updateSchema(props.schema, attrs),
        });
      },
    },

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

    permissions.admin && {
      title: 'Редагувати Доступ',
      icon: PeopleIcon,

      async onAction() {
        await schemaUsersStore.load(props.schema.id);
        accessEditModal.open();
      },
    },

    permissions.admin && {
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
    box-shadow: var(--box-shadow);
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
