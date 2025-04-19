<template>
  <aside class="editor-sidebar">
    <SidebarHeader />
    <SidebarGeneralSection />
    <SidebarStructureSection />
  </aside>
</template>

<script setup lang="ts">
import { SidebarGeneralSection, SidebarStructureSection } from './sections';
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
