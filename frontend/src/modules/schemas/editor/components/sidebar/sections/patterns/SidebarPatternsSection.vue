<template>
  <SidebarSection class="sidebar-patterns" :padding="false">
    <SidebarSectionHeading class="sidebar-patterns__heading">
      Патерни

      <Button icon @click="addPatternModal.open()">
        <PlusIcon />
      </Button>
    </SidebarSectionHeading>

    <ul class="sidebar-patterns__list">
      <SidebarPattern
        v-for="pattern of patternsStore.patterns"
        :key="pattern.id"
        :pattern
      />
    </ul>
  </SidebarSection>
</template>

<script setup lang="ts">
import { usePatternsStore } from '@/modules/schemas/editor/stores';
import { Button } from '@/components/button';
import { PlusIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import EditorAddPatternModal from '../../../EditorAddPatternModal.vue';
import { SidebarSection, SidebarSectionHeading } from '../base';
import SidebarPattern from './SidebarPattern.vue';

const patternsStore = usePatternsStore();
const addPatternModal = useModal(EditorAddPatternModal);
</script>

<style scoped>
@layer page {
  .sidebar-patterns {
    padding-bottom: 8px;
  }

  .sidebar-patterns__heading {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 6px 8px 6px 12px;
  }

  .sidebar-patterns__list {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }
}
</style>
