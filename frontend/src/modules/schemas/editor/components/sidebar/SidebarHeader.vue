<template>
  <header class="editor-sidebar__header">
    <Button
      class="editor-sidebar__header-back"
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
  .editor-sidebar__header {
    display: flex;
    align-items: center;
    padding: 4px 8px;
    background-color: var(--color-white);
    border-bottom: 1px solid var(--color-divider);
  }

  .editor-sidebar__header-back {
    gap: 8px;
  }
}
</style>
