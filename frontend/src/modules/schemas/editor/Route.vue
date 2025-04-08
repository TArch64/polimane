<template>
  <EditorTopBar />
</template>

<script lang="ts" setup>
import { definePreload } from '@/router/define';
import { useEditorStore } from '../stores';
import { EditorTopBar } from './components';

defineProps<{
  schemaId: string;
}>();

defineOptions({
  beforeRouteEnter: definePreload<'schema-editor'>(async (route) => {
    const store = useEditorStore();
    await store.loadSchema(route.params.schemaId);
  }),
});
</script>

<style>
@layer page {
  .app--schema-editor {
    background-color: var(--color-background-2);
    overflow: hidden;
  }
}
</style>
