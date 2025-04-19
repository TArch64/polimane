<template>
  <SidebarSection class="sidebar-structure" :padding="false">
    <SidebarSectionHeading class="sidebar-structure__heading">
      Структура

      <Button icon @click="addPatternModal.open()">
        <PlusIcon />
      </Button>
    </SidebarSectionHeading>

    <ul class="sidebar-structure__list">
      <SidebarStructurePattern
        v-for="pattern of patternsStore.patterns.values"
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
import { SidebarStructurePattern } from './pattern';

const patternsStore = usePatternsStore();
const addPatternModal = useModal(EditorAddPatternModal);
</script>

<style scoped>
@layer page {
  .sidebar-structure {
    padding-bottom: 8px;
  }

  .sidebar-structure__heading {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 6px 8px 6px 12px;
  }

  .sidebar-structure__list {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }
}
</style>
