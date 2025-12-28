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

const props = defineProps<{
  schema: ListSchema;
}>();

const renameAction = useSchemaMenuCopy(() => props.schema);
const addToDirectoryAction = useSchemaMenuAddToDirectory(() => props.schema);
const copyAction = useSchemaMenuRename(() => props.schema);
const editAccessAction = useSchemaMenuEditAccess(() => props.schema);
const deleteAction = useSchemaMenuDelete(() => props.schema);

const menuActions = computed((): MaybeContextMenuAction[] => [
  renameAction.value,
  addToDirectoryAction.value,
  copyAction.value,
  editAccessAction.value,
  deleteAction.value,
]);
</script>
