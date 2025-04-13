<template>
  <EditorTopBar />

  <div class="editor__fill editor__row" v-if="patternsStore.hasPatterns">
    <EditorSidebar class="editor__sidebar" />
    <EditorCanvas class="editor__fill" />
  </div>

  <EditorEmpty class="editor__fill" v-else />
</template>

<script lang="ts" setup>
import { definePreload } from '@/router/define';
import { useEditorStore, usePatternsStore } from '../stores';
import { EditorCanvas, EditorEmpty, EditorSidebar, EditorTopBar } from './components';

defineProps<{
  schemaId: string;
}>();

defineOptions({
  beforeRouteEnter: definePreload<'schema-editor'>(async (route) => {
    const store = useEditorStore();
    await store.loadSchema(route.params.schemaId);
  }),
});

const patternsStore = usePatternsStore();
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
  }
}
</style>
