<template>
  <EditorHeader />

  <template v-if="patternsStore.hasPatterns">
    <EditorCanvas class="editor__fill" />
    <EditorPalette />
  </template>

  <EditorEmpty class="editor__fill" v-else />
</template>

<script lang="ts" setup>
import { useEventListener } from '@vueuse/core';
import { definePreload } from '@/router/define';
import { destroyStore, lazyDestroyStore } from '@/helpers';
import {
  disposeBeadsStores,
  disposeRowsStores,
  useCursorStore,
  useDraggingStore,
  useEditorStore,
  usePaletteStore,
  usePatternsStore,
} from './stores';
import { EditorCanvas, EditorEmpty, EditorHeader, EditorPalette } from './components';

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
    lazyDestroyStore(usePatternsStore);
    disposeBeadsStores();
    disposeRowsStores();
    lazyDestroyStore(usePaletteStore);
    lazyDestroyStore(useDraggingStore);
    lazyDestroyStore(useCursorStore);
    next();
  },
});

const editorStore = useEditorStore();
const patternsStore = usePatternsStore();

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
    background-color: var(--color-background-2);
    overflow: hidden;
  }

  .editor__fill {
    flex-grow: 1;
    flex-basis: 0;
    min-height: 0;
  }
}
</style>
