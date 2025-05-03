<template>
  <div class="editor__fill editor__row" v-if="patternsStore.hasPatterns">
    <EditorSidebar class="editor__sidebar" />
    <EditorCanvas class="editor__fill" />
  </div>

  <EditorEmpty class="editor__fill" v-else />
</template>

<script lang="ts" setup>
import { useEventListener } from '@vueuse/core';
import { definePreload } from '@/router/define';
import { destroyStore } from '@/helpers';
import { disposeBeadsStores, disposeRowsStores, useEditorStore, usePatternsStore } from './stores';
import { EditorCanvas, EditorEmpty, EditorSidebar } from './components';

defineProps<{
  schemaId: string;
}>();

defineOptions({
  beforeRouteEnter: definePreload<'schema-editor'>(async (route) => {
    const store = useEditorStore();
    await store.loadSchema(route.params.schemaId);
  }),

  beforeRouteLeave: async (_, __, next) => {
    destroyStore(useEditorStore());
    disposeBeadsStores();
    disposeRowsStores();
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

  .editor__row {
    display: flex;
  }

  .editor__sidebar {
    width: 270px;
    flex-shrink: 0;
  }

  .editor__fill {
    flex-grow: 1;
    flex-basis: 0;
    min-height: 0;
  }
}
</style>
