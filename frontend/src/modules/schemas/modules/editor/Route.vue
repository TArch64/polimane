<template>
  <EditorHeader />
  <EditorToolbar v-if="isEditable" />
  <EditorCanvas class="editor__fill" />
  <EditorBottomBar v-if="!isMobile" />
</template>

<script lang="ts" setup>
import './colorLib';

import { computed } from 'vue';
import { useEventListener } from '@vueuse/core';
import { definePreload } from '@/router/define';
import { destroyStore, lazyDestroyStore } from '@/helpers';
import { useMobileScreen } from '@/composables';
import {
  useBeadsStore,
  useCanvasStore,
  useEditorStore,
  useHistoryStore,
  useSelectionStore,
  useToolsStore,
} from './stores';
import { EditorBottomBar, EditorCanvas, EditorHeader, EditorToolbar } from './components';
import { provideHotKeysHandler, useEditorBackgroundRenderer } from './composables';

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
    lazyDestroyStore(useCanvasStore);
    lazyDestroyStore(useToolsStore);
    lazyDestroyStore(useSelectionStore);
    lazyDestroyStore(useBeadsStore);
    next();
  },
});

const editorStore = useEditorStore();
const isMobile = useMobileScreen();
useEditorBackgroundRenderer();
provideHotKeysHandler();

const isEditable = computed(() => editorStore.canEdit && !isMobile.value);

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
    --editor-ui-padding: 12px;
  }

  .editor__fill {
    flex-grow: 1;
    flex-basis: 0;
    min-height: 0;
  }
}
</style>
