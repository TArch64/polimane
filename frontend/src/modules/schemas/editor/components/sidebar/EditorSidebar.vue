<template>
  <aside class="editor-sidebar">
    <SidebarHeader />
    <SidebarGeneralSection />
    <SidebarPatternsSection />
  </aside>
</template>

<script setup lang="ts">
import { SidebarGeneralSection, SidebarPatternsSection } from './sections';
import SidebarHeader from './SidebarHeader.vue';
</script>

<style scoped>
@layer page {
  .editor-sidebar {
    background-color: var(--color-background-1);
    border-right: var(--divider);
  }
}
</style>
