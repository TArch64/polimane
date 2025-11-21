<template>
  <HomeListCard
    :to="editorRoute"
    :selected="isSelected"
    :menu-title="schema.name"
    :menuActions
  >
    <HomeListScreenshot
      :path="schema.screenshotPath"
      :alt="`Скріншот схеми ${schema.name}`"
      :background-color="schema.backgroundColor"
    />

    {{ schema.name }}
  </HomeListCard>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { type RouteLocationRaw, useRouter } from 'vue-router';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { useAccessPermissions } from '@/composables';
import { useConfirm } from '@/components/confirm';
import { CopyIcon, EditIcon, FolderIcon, PeopleIcon, TrashIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { SchemaRenameModal } from '@/modules/schemas/shared/modals/rename';
import {
  SchemaAccessEditModal,
  useSchemaUsersStore,
} from '@/modules/schemas/shared/modals/accessEdit';
import type { ISchema } from '@/models';
import { useSchemasStore } from '@/modules/home/modules/list/stores';
import type { ListSchema } from '../../stores';
import { FolderAddSchemaModal } from '../modals';
import HomeListCard from './HomeListCard.vue';
import HomeListScreenshot from './HomeListScreenshot.vue';

const props = defineProps<{
  schema: ListSchema;
}>();

const router = useRouter();

const schemasStore = useSchemasStore();
const schemaUsersStore = useSchemaUsersStore();

const isSelected = computed(() => schemasStore.selected.has(props.schema.id));
const permissions = useAccessPermissions(() => props.schema.access);

const renameModal = useModal(SchemaRenameModal);
const folderAddModal = useModal(FolderAddSchemaModal);
const accessEditModal = useModal(SchemaAccessEditModal);

const editorRoute = computed((): RouteLocationRaw => ({
  name: 'schema-editor',
  params: { schemaId: props.schema.id },
}));

const deleteConfirm = useConfirm({
  danger: true,
  control: false,
  message: 'Ви впевнені, що хочете видалити цю схему?',
  acceptButton: 'Видалити',
});

const menuActions = computed((): MaybeContextMenuAction[] => [
  permissions.write && {
    title: 'Змінити назву',
    icon: EditIcon,

    onAction() {
      renameModal.open({
        schema: props.schema as ISchema,
        updateSchema: (attrs) => schemasStore.updateSchema(props.schema, attrs),
      });
    },
  },

  {
    title: 'Перемістити в Директорію',
    icon: FolderIcon,

    onAction() {
      folderAddModal.open({
        schemaIds: [props.schema.id],
        folderId: null,
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
      await schemaUsersStore.load([props.schema.id]);
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
]);
</script>
