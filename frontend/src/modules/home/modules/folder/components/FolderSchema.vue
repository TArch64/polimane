<template>
  <HomeListSchema :schema :menuActions />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { HomeListSchema } from '@/modules/home/components';
import {
  useSchemaMenuAddToDirectory,
  useSchemaMenuCopy,
  useSchemaMenuDelete,
  useSchemaMenuEditAccess,
  useSchemaMenuRename,
} from '@/modules/home/composables';
import type { ListSchema } from '@/modules/home/stores';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import { FolderRemoveIcon } from '@/components/icon';
import { useFolderSchemasStore, useFolderStore } from '../stores';

const props = defineProps<{
  schema: ListSchema;
}>();

const folderStore = useFolderStore();
const schemasStore = useFolderSchemasStore();

const renameAction = useSchemaMenuCopy(() => props.schema);
const addToDirectoryAction = useSchemaMenuAddToDirectory(() => props.schema, folderStore.folder.id);
const copyAction = useSchemaMenuRename(() => props.schema);
const editAccessAction = useSchemaMenuEditAccess(() => props.schema);
const deleteAction = useSchemaMenuDelete(() => props.schema);

const menuActions = computed((): MaybeContextMenuAction[] => [
  renameAction.value,
  addToDirectoryAction.value,

  {
    title: 'Видалити з Директорії',
    icon: FolderRemoveIcon,
    onAction: () => schemasStore.removeSchemaFromFolder(props.schema),
  },

  copyAction.value,
  editAccessAction.value,
  deleteAction.value,
]);
</script>
