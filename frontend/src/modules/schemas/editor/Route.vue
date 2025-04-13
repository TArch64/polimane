<template>
  <EditorTopBar />
  <div class="editor__row">
    <EditorSidebar class="editor__sidebar" />
    <EditorCanvas class="editor__canvas" />
  </div>
</template>

<script lang="ts" setup>
import { definePreload } from '@/router/define';
import { useEditorStore } from '../stores';
import { EditorCanvas, EditorSidebar, EditorTopBar } from './components';

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

<style scoped>
@layer page {
  :global(.app--schema-editor) {
    background-color: var(--color-background-2);
    overflow: hidden;
  }

  .editor__row {
    flex-grow: 1;
    flex-basis: 0;
    display: flex;
  }

  .editor__sidebar {
    width: 270px;
    flex-shrink: 0;
  }

  .editor__canvas {
    flex-grow: 1;
    flex-basis: 0;
  }
}
</style>
