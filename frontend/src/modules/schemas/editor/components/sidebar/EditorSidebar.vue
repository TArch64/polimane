<template>
  <aside class="editor-sidebar">
    <SidebarGeneralSection />
  </aside>
</template>

<script setup lang="ts">
import { SidebarGeneralSection } from './sections';
</script>

<style scoped>
@layer page {
  .editor-sidebar {
    background-color: var(--color-background-1);
    border-right: var(--divider);
  }
}
</style>
