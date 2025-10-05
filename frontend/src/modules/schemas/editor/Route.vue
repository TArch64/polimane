<template>
  <EditorHeader />
  <EditorToolbar />
  <EditorCanvas class="editor__fill" />
</template>

<script lang="ts" setup>
import './colorLib';

import { useEventListener } from '@vueuse/core';
import { definePreload } from '@/router/define';
import { destroyStore, lazyDestroyStore } from '@/helpers';
import { useBeadsStore, useEditorStore, useHistoryStore, useToolsStore } from './stores';
import { EditorCanvas, EditorHeader, EditorToolbar } from './components';
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
    lazyDestroyStore(useHistoryStore);
    lazyDestroyStore(useToolsStore);
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
@property --editor-background-color {
  syntax: '<color>';
  inherits: true;
  initial-value: #F8F8F8;
}

@layer page {
  :global(.app--schema-editor) {
    overflow: hidden;
    background-color: var(--editor-background-color);
    transition: background-color 0.15s ease-out;
    will-change: background-color;
  }

  .editor__fill {
    flex-grow: 1;
    flex-basis: 0;
    min-height: 0;
  }
}
</style>
