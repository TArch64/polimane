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
  </header>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { ArrowBackIcon } from '@/components/icon';
</script>

<style scoped>
@layer page {
  .editor-top-bar {
    display: flex;
    align-items: center;
    padding: 8px;
    background-color: var(--color-white);
    border-bottom: 1px solid var(--color-divider);
  }

  .editor-top-bar__back {
    gap: 8px;
  }
}
</style>
