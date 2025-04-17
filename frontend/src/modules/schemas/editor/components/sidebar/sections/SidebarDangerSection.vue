<template>
  <SidebarSection>
    <Button variant="danger" size="md" class="sidebar-danger-section__button">
      Видалити Схему
    </Button>
  </SidebarSection>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
import SidebarSection from './SidebarSection.vue';
</script>

<style scoped>
@layer page {
  .sidebar-danger-section__button {
    width: 100%;
  }
}
</style>
