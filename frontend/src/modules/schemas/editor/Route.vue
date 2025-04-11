<template>
  <EditorTopBar />
  <EditorCanvas class="editor__canvas" />
</template>

<script lang="ts" setup>
import { definePreload } from '@/router/define';
import { useEditorStore } from '../stores';
import { EditorCanvas, EditorTopBar } from './components';

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

  .editor__canvas {
    flex-grow: 1;
  }
}
</style>
