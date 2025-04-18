<template>
  <SidebarSection class="sidebar-patterns" :padding="false">
    <SidebarSectionHeading class="sidebar-patterns__heading">
      Патерни
    </SidebarSectionHeading>

    <ul class="sidebar-patterns__list">
      <SidebarPattern
        v-for="(pattern, index) of patternsStore.patterns"
        :key="pattern.id"
        :index
        :pattern
      />
    </ul>
  </SidebarSection>
</template>

<script setup lang="ts">
import { usePatternsStore } from '@/modules/schemas/editor/stores';
import { SidebarSection, SidebarSectionHeading } from '../base';
import SidebarPattern from './SidebarPattern.vue';

const patternsStore = usePatternsStore();
</script>

<style scoped>
@layer page {
  .sidebar-patterns {
    padding: 8px 0;
  }

  .sidebar-patterns__heading {
    padding: 4px 12px 8px;
  }

  .sidebar-patterns__list {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }
}
</style>
