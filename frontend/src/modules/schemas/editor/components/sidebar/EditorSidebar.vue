<template>
  <aside class="editor-sidebar" ref="asideRef">
    <SidebarHeader />
    <SidebarGeneralSection />
    <SidebarStructureSection />
  </aside>
</template>

<script setup lang="ts">
import { provide, ref } from 'vue';
import { TOKEN_SCROLLER } from '@/InjectionToken';
import { SidebarGeneralSection, SidebarStructureSection } from './sections';
import SidebarHeader from './SidebarHeader.vue';

const asideRef = ref<HTMLElement>(null!);
provide(TOKEN_SCROLLER, asideRef);
</script>

<style scoped>
@layer page {
  .editor-sidebar {
    background-color: var(--color-background-1);
    border-right: var(--divider);
    height: 100%;
    overflow-y: auto;
  }
}
</style>
