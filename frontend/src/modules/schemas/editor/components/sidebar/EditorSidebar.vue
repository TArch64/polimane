<template>
  <aside class="editor-sidebar">
    <SidebarGeneralSection />
    <SidebarDangerSection />
  </aside>
</template>

<script setup lang="ts">
import { SidebarDangerSection, SidebarGeneralSection } from './sections';
</script>

<style scoped>
@layer page {
  .editor-sidebar {
    background-color: var(--color-background-1);
    border-right: var(--divider);
  }
}
</style>
