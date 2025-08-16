<template>
  <EditorHeader />
  <EditorCanvas class="editor__fill" @rendered="loaderStore.hide" />
  <EditorPalette />
</template>

<script lang="ts" setup>
import { useEventListener } from '@vueuse/core';
import { definePreload } from '@/router/define';
import { destroyStore, lazyDestroyStore } from '@/helpers';
import { useLoaderStore } from '@/stores';
import { useEditorStore, usePaletteStore } from './stores';
import { EditorCanvas, EditorHeader, EditorPalette } from './components';

defineProps<{
  schemaId: string;
}>();

defineOptions({
  beforeRouteEnter: definePreload<'schema-editor'>(async (route) => {
    const loaderStore = useLoaderStore();
    const editorStore = useEditorStore();

    loaderStore.show();
    await editorStore.loadSchema(route.params.schemaId);
  }, { appLoader: true }),

  beforeRouteLeave: async (_, __, next) => {
    lazyDestroyStore(useEditorStore);
    lazyDestroyStore(usePaletteStore);
    next();
  },
});

const loaderStore = useLoaderStore();
const editorStore = useEditorStore();

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
    --app-background-color: var(--color-background-2);
    overflow: hidden;
  }

  .editor__fill {
    flex-grow: 1;
    flex-basis: 0;
    min-height: 0;
  }
}
</style>
