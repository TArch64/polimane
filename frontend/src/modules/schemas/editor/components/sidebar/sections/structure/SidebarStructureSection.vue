<template>
  <SidebarSection class="sidebar-structure" :padding="false">
    <SidebarSectionHeading class="sidebar-structure__heading">
      Структура

      <Button icon @click="addPatternModal.open()">
        <PlusIcon />
      </Button>
    </SidebarSectionHeading>

    <SidebarPatternList />
  </SidebarSection>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import { PlusIcon } from '@/components/icon';
import { useModal } from '@/components/modal';
import { PatternAddModal } from '../../../modals';
import { SidebarSection, SidebarSectionHeading } from '../base';
import { SidebarPatternList } from './pattern';

const addPatternModal = useModal(PatternAddModal);
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
}
</style>
