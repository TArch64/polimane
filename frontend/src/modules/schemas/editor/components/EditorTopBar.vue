<template>
  <header class="editor-top-bar">
    <Button
      class="editor-top-bar__back"
      variant="secondary"
      size="md"
      :to="{ name: 'home' }"
    >
      <ArrowBackIcon />
      Редактор
    </Button>

    <p class="editor-top-bar__saving-state" v-if="editorStore.hasUnsavedChanges">
      Є незбережені зміни
    </p>
  </header>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { ArrowBackIcon } from '@/components/icon';
import { useEditorStore } from '../stores';

const editorStore = useEditorStore();
</script>

<style scoped>
@layer page {
  .editor-top-bar {
    display: flex;
    align-items: center;
    padding: 4px 8px;
    background-color: var(--color-white);
    border-bottom: 1px solid var(--color-divider);
  }

  .editor-top-bar__back {
    gap: 8px;
  }

  .editor-top-bar__saving-state {
    margin-left: 8px;
    font-size: var(--font-sm);
    color: color-mix(in srgb, var(--color-black), transparent 40%);
  }
}
</style>
