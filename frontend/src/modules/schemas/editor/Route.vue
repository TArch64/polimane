<template>
  <EditorHeader />
  <EditorCanvas class="editor__fill" />
  <EditorPalette />
</template>

<script lang="ts" setup>
import './colorLib';

import { useEventListener } from '@vueuse/core';
import { definePreload } from '@/router/define';
import { destroyStore, lazyDestroyStore } from '@/helpers';
import { useBeadsStore, useEditorStore, usePaletteStore } from './stores';
import { EditorCanvas, EditorHeader, EditorPalette } from './components';
import { useEditorBackgroundRenderer } from './composables';

defineProps<{
  schemaId: string;
}>();

defineOptions({
  beforeRouteEnter: definePreload<'schema-editor'>(async (route) => {
    const store = useEditorStore();
    await store.loadSchema(route.params.schemaId);
  }),

  beforeRouteLeave: async (_, __, next) => {
    lazyDestroyStore(useEditorStore);
    lazyDestroyStore(usePaletteStore);
    lazyDestroyStore(useBeadsStore);
    next();
  },
});

const editorStore = useEditorStore();
useEditorBackgroundRenderer();

useEventListener(window, 'beforeunload', (event) => {
  if (editorStore.hasUnsavedChanges) {
    event.preventDefault();
    destroyStore(editorStore);
  }
});
</script>

<style scoped>
@layer page {
  :global(.app--schema-editor) {
    background-color: var(--editor-background-color);
    overflow: hidden;
  }

  .editor__fill {
    flex-grow: 1;
    flex-basis: 0;
    min-height: 0;
  }
}
</style>
