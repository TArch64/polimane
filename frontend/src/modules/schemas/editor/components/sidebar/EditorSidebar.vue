<template>
  <aside class="editor-sidebar" ref="asideRef">
    <SidebarHeader class="editor-sidebar__header" />

    <div class="editor-sidebar__content">
      <SidebarGeneralSection />
      <SidebarStructureSection />
    </div>
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
    display: flex;
    flex-direction: column;
  }

  .editor-sidebar__header {
    flex-shrink: 0;
  }

  .editor-sidebar__content {
    overflow-y: auto;
    flex-basis: 0;
    flex-grow: 1;
    min-height: 0;
  }
}
</style>
